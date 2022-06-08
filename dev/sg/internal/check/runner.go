package check

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/sourcegraph/sourcegraph/dev/sg/internal/std"
	"github.com/sourcegraph/sourcegraph/dev/sg/internal/usershell"
	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/sourcegraph/sourcegraph/lib/output"
)

type Runner[Args any] struct {
	in         io.Reader
	out        *std.Output
	categories []Category[Args]

	renderDescription func(*std.Output)
}

func NewRunner[Args any](in io.Reader, out *std.Output, categories []Category[Args]) *Runner[Args] {
	return &Runner[Args]{
		in:         in,
		out:        out,
		categories: categories,
	}
}

func (r *Runner[Args]) SetDescription(render func(out *std.Output)) {
	r.renderDescription = render
}

// Check executes all checks exactly once and exits.
func (r *Runner[Args]) Check(
	ctx context.Context,
	args Args,
) error {
	ctx, err := usershell.Context(ctx)
	if err != nil {
		return err
	}

	results := r.runAllCategoryChecks(ctx, args)
	if len(results.failed) > 0 {
		return errors.Newf("%d checks failed (%d skipped)", len(results.failed), len(results.skipped))
	}

	return nil
}

// Fix attempts to execute available fixes on checks that are not fulfilled.
func (r *Runner[Args]) Fix(
	ctx context.Context,
	args Args,
) error {
	ctx, err := usershell.Context(ctx)
	if err != nil {
		return err
	}

	// Get state
	results := r.runAllCategoryChecks(ctx, args)
	if len(results.failed) == 0 {
		// Nothing failed, we're good to go!
		return nil
	}

	r.out.WriteNoticef("Attempting to fix %d failed categories", len(results.failed))
	for _, i := range results.failed {
		category := r.categories[i]

		ok := r.fixCategoryAutomatically(ctx, i+1, &category, args, results)
		results.categories[category.Name] = ok
	}

	// Report what is still bust
	failedCategories := []string{}
	for _, category := range r.categories {
		if !category.IsSatisfied() {
			failedCategories = append(failedCategories, fmt.Sprintf("%q", category.Name))
		}
	}
	if len(failedCategories) > 0 {
		return errors.Newf("Some categories are still unsatisfied: %s", strings.Join(failedCategories, ", "))
	}

	return nil
}

// Interactive runs both checks and fixes in an interactive manner, prompting the user for
// input.
func (r *Runner[Args]) Interactive(
	ctx context.Context,
	args Args,
) error {
	ctx, err := usershell.Context(ctx)
	if err != nil {
		return err
	}

	// Keep interactive runner up until all issues are fixed.
	results := &runAllCategoryChecksResult{
		failed: []int{1}, // initialize, this gets reset immediately
	}
	for len(results.failed) != 0 {
		r.out.Output.ClearScreen()

		// Get args when we run
		results = r.runAllCategoryChecks(ctx, args)
		if len(results.failed) == 0 {
			break
		}

		r.out.WriteWarningf("Some checks failed. Which one do you want to fix?")

		idx, err := r.getNumberOutOf(results.all)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		selectedCategory := r.categories[idx]

		r.out.ClearScreen()

		err = r.presentFailedCategoryWithOptions(ctx, idx, &selectedCategory, args, results)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}

// runAllCategoryChecksResult provides a summary of categories checks results.
type runAllCategoryChecksResult struct {
	all     []int
	failed  []int
	skipped []int

	// Indicates whether each category succeeded or not
	categories map[string]bool
}

// runAllCategoryChecks is the main entrypoint for running the checks in this runner.
func (r *Runner[Args]) runAllCategoryChecks(ctx context.Context, args Args) *runAllCategoryChecksResult {
	if r.renderDescription != nil {
		r.renderDescription(r.out)
	}

	var results runAllCategoryChecksResult
	for i := range r.categories {
		results.failed = append(results.failed, i)
		results.all = append(results.all, i)
	}

	results.categories = make(map[string]bool)
	bars := []output.ProgressBar{}
	for _, category := range r.categories {
		results.categories[category.Name] = false
		bars = append(bars, output.ProgressBar{
			Label: category.Name,
			Max:   float64(len(category.Checks)),
		})
	}
	progress := r.out.Progress(bars, nil)

	var resultsLock sync.Mutex
	var wg sync.WaitGroup
	for i, cat := range r.categories {
		idx := i + 1
		category := cat

		progress.VerboseLine(output.Styledf(output.StylePending, "%d. %s - Determining status...", idx, category.Name))
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			if err := category.CheckEnabled(ctx, args); err != nil {
				progress.WriteLine(output.Linef(output.EmojiQuestionMark, output.StyleGrey, "%d. %s %s[SKIPPED: %s]%s",
					idx, category.Name, output.StyleBold, err.Error(), output.StyleReset))
				results.skipped = append(results.skipped, idx)
				results.failed = removeEntry(results.failed, i)
				return
			}

			// Validate checks
			var failed bool
			for ci, c := range category.Checks {
				var out strings.Builder
				cio := IO{
					Input:  r.in,
					Writer: std.NewFixedOutput(&out, true),
				}

				if !c.IsEnabled(ctx, cio, args) {
					continue
				}

				if err := c.Update(ctx, cio, args); err != nil {
					failed = true
					outStr := out.String()
					if len(outStr) > 0 {
						progress.Verbose(outStr)
					}
					progress.WriteLine(output.Styled(output.StyleWarning, err.Error()))
				}
				progress.SetValue(i, float64(ci+1))
			}

			// Report results by updating the results value
			resultsLock.Lock()
			defer resultsLock.Unlock()
			results.categories[category.Name] = !failed
			if !failed {
				progress.WriteLine(output.Linef(output.EmojiSuccess, output.StyleSuccess, "%d. %s", idx, category.Name))
				results.failed = removeEntry(results.failed, i)
			} else {
				progress.WriteLine(output.Linef(output.EmojiFailure, output.StyleFailure, "%d. %s", idx, category.Name))
			}
		}(i)
	}

	wg.Wait()
	progress.Destroy()

	if len(results.failed) == 0 {
		if len(results.skipped) == 0 {
			r.out.Write("")
			r.out.WriteLine(output.Linef(output.EmojiOk, output.StyleBold, "Everything looks good! Happy hacking!"))
		} else {
			r.out.Write("")
			r.out.WriteWarningf("Some checks were skipped.")
		}
	}

	return &results
}

func removeEntry(s []int, val int) (result []int) {
	for _, e := range s {
		if e != val {
			result = append(result, e)
		}
	}
	return result
}

func (r *Runner[Args]) getNumberOutOf(numbers []int) (int, error) {
	var strs []string
	var idx = make(map[int]struct{})
	for _, num := range numbers {
		strs = append(strs, fmt.Sprintf("%d", num+1))
		idx[num+1] = struct{}{}
	}

	for {
		r.out.Writef("[%s]: ", strings.Join(strs, ","))
		var num int
		_, err := fmt.Fscan(r.in, &num)
		if err != nil {
			return 0, err
		}

		if _, ok := idx[num]; ok {
			return num - 1, nil
		}
		r.out.Writef("%d is an invalid choice :( Let's try again?\n", num)
	}
}

func (r *Runner[Args]) presentFailedCategoryWithOptions(ctx context.Context, categoryIdx int, category *Category[Args], args Args, results *runAllCategoryChecksResult) error {
	r.printCategoryHeaderAndDependencies(categoryIdx+1, category)
	fixableCategory := category.HasFixable()

	choices := map[int]string{1: "I want to fix these manually"}
	if fixableCategory {
		choices[2] = "I'm feeling lucky. You try fixing all of it for me."
		choices[3] = "Go back"
	} else {
		choices[2] = "Go back"
	}

	choice, err := r.getChoice(choices)
	if err != nil {
		return err
	}

	switch choice {
	case 1:
		err = r.fixCategoryManually(ctx, categoryIdx, category, args)
	case 2:
		if fixableCategory {
			r.out.ClearScreen()
			if !r.fixCategoryAutomatically(ctx, categoryIdx, category, args, results) {
				err = errors.Newf("%s: failed to fix category automatically", category.Name)
			}
		}
	case 3:
		return nil
	}
	return err
}

func (r *Runner[Args]) printCategoryHeaderAndDependencies(categoryIdx int, category *Category[Args]) {
	r.out.WriteLine(output.Linef(output.EmojiLightbulb, output.CombineStyles(output.StyleSearchQuery, output.StyleBold), "%d. %s", categoryIdx, category.Name))
	r.out.Write("")
	r.out.Write("Checks:")

	for i, dep := range category.Checks {
		idx := i + 1
		if dep.IsSatisfied() {
			r.out.WriteSuccessf("%d. %s", idx, dep.Name)
		} else {
			if dep.checkErr != nil {
				r.out.WriteFailuref("%d. %s: %s", idx, dep.Name, dep.checkErr)
			} else {
				r.out.WriteFailuref("%d. %s: %s", idx, dep.Name, "check failed")
			}
		}
	}
}

func (r *Runner[Args]) getChoice(choices map[int]string) (int, error) {
	for {
		r.out.Write("")
		r.out.WriteNoticef("What do you want to do?")

		for i := 0; i < len(choices); i++ {
			num := i + 1
			desc, ok := choices[num]
			if !ok {
				return 0, errors.Newf("internal error: %d not found in provided choices", i)
			}
			r.out.Writef("%s[%d]%s: %s", output.StyleBold, num, output.StyleReset, desc)
		}

		fmt.Printf("Enter choice: ")

		var s int
		_, err := fmt.Scan(&s)
		if err != nil {
			return 0, err
		}

		if _, ok := choices[s]; ok {
			return s, nil
		}
		r.out.WriteFailuref("Invalid choice")
	}
}

func (r *Runner[Args]) fixCategoryManually(ctx context.Context, categoryIdx int, category *Category[Args], args Args) error {
	for {
		toFix := []int{}

		for i, dep := range category.Checks {
			if dep.IsSatisfied() {
				continue
			}

			toFix = append(toFix, i)
		}

		if len(toFix) == 0 {
			break
		}

		var idx int

		if len(toFix) == 1 {
			idx = toFix[0]
		} else {
			r.out.WriteNoticef("Which one do you want to fix?")
			var err error
			idx, err = r.getNumberOutOf(toFix)
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}

		dep := category.Checks[idx]

		r.out.WriteLine(output.Linef(output.EmojiFailure, output.CombineStyles(output.StyleWarning, output.StyleBold), "%s", dep.Name))
		r.out.Write("")

		if dep.checkErr != nil {
			r.out.WriteLine(output.Styledf(output.StyleBold, "Encountered the following error:\n\n%s%s\n", output.StyleReset, dep.checkErr))
		}

		r.out.WriteLine(output.Styled(output.StyleBold, "How to fix:"))

		r.out.WriteMarkdown(dep.Description)

		pending := r.out.Pending(output.Styled(output.StylePending, "Determining status..."))
		for _, dep := range category.Checks {
			// update check state
			_ = dep.Update(ctx, IO{
				Input:  r.in,
				Writer: pending,
			}, args)
		}
		pending.Destroy()

		r.printCategoryHeaderAndDependencies(categoryIdx, category)
	}

	return nil
}

func (r *Runner[Args]) fixCategoryAutomatically(ctx context.Context, categoryIdx int, category *Category[Args], args Args, results *runAllCategoryChecksResult) (ok bool) {
	pending := r.out.Pending(output.Styledf(output.StylePending, "Trying my hardest to fix %q automatically...", category.Name))
	cio := IO{
		Input:  r.in,
		Writer: pending,
	}

	// Make sure to call this with a final message before returning!
	complete := func(emoji string, style output.Style, fmtStr string, args ...any) {
		pending.Complete(output.Linef(emoji, style, "%d. %s - "+fmtStr, append([]any{categoryIdx, category.Name}, args...)...))
	}

	if err := category.CheckEnabled(ctx, args); err != nil {
		complete(output.EmojiQuestionMark, output.StyleGrey, "Skipped: %s", err.Error())
		return true
	}

	// If nothing in this category is fixable, we are done
	if !category.HasFixable() {
		complete(output.EmojiFailure, output.StyleFailure, "Cannot be fixed automatically.")
		return false
	}

	// Only run if dependents are fixed
	var unmetDependencies []string
	for _, d := range category.DependsOn {
		if met, exists := results.categories[d]; !exists {
			complete(output.EmojiFailure, output.StyleFailure, "Required check category %q not found", d)
			return false
		} else if !met {
			unmetDependencies = append(unmetDependencies, fmt.Sprintf("%q", d))
		}
	}
	if len(unmetDependencies) > 0 {
		complete(output.EmojiFailure, output.StyleFailure, "Required dependencies %s not met.", strings.Join(unmetDependencies, ", "))
		return false
	}

	// now go through the real dependencies
	var fixFailed bool
	for _, c := range category.Checks {
		// If category is fixed, we are good to go
		if c.IsSatisfied() {
			continue
		}

		// Skip
		if !c.IsEnabled(ctx, cio, args) {
			continue
		}

		// Otherwise, check if this is fixable at all
		if c.Fix == nil {
			pending.WriteLine(output.Styledf(output.StyleGrey, "%q cannot be fixed automatically.", c.Name))
			continue
		}

		// Attempt fix. Get new args because things might have changed due to another
		// fix being run.
		pending.VerboseLine(output.Styledf(output.StylePending, "Trying to fix %q...", c.Name))
		if err := c.Fix(ctx, cio, args); err != nil {
			pending.WriteLine(output.Styledf(output.StyleWarning, "Failed to fix %q: %s", c.Name, err.Error()))
			fixFailed = true
			continue
		}

		// Check is the fix worked
		if err := c.Update(ctx, cio, args); err != nil {
			pending.WriteLine(output.Styledf(output.StyleWarning, "Check %q still failing: %s",
				c.Name, err.Error()))
			fixFailed = true
		}
	}

	if fixFailed {
		complete(output.EmojiFailure, output.StyleFailure, "Fixes failed")
	} else {
		complete(output.EmojiSuccess, output.StyleSuccess, "Done!")
	}

	return !fixFailed
}

func waitForReturn() { fmt.Scanln() }
