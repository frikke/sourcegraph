package main

import (
	"testing"
	"time"

	"github.com/hexops/autogold/v2"
)

func TestBranchEventSummary(t *testing.T) {
	t.Run("unlocked", func(t *testing.T) {
		got := generateBranchEventSummary(false, "main", "#buildkite-main", []CommitInfo{})

		want := autogold.Expect(":white_check_mark: Pipeline healthy - `main` unlocked!")
		want.Equal(t, got)
	})

	t.Run("locked", func(t *testing.T) {
		got := generateBranchEventSummary(true, "main", "#buildkite-main", []CommitInfo{
			{Commit: "a", Author: "bob", AuthorSlackID: "123", BuildNumber: 3, BuildURL: "https://sourcegraph.com", BuildCreated: time.Now()},
			{Commit: "b", Author: "alice", AuthorSlackID: "124", BuildNumber: 2, BuildURL: "https://sourcegraph.com", BuildCreated: time.Now().Add(-1)},
			{Commit: "c", Author: "no_slack", AuthorSlackID: "", BuildNumber: 1, BuildURL: "https://sourcegraph.com", BuildCreated: time.Now().Add(-2)},
		})

		want := autogold.Expect(":alert: *Consecutive build failures detected - the `main` branch has been locked.* :alert:\nThe authors of the following failed commits who are Sourcegraph teammates have been granted merge access to investigate and resolve the issue:\n\n- <https://github.com/sourcegraph/sourcegraph/commit/c|c> (<https://sourcegraph.com|build 1>): no_slack\n- <https://github.com/sourcegraph/sourcegraph/commit/b|b> (<https://sourcegraph.com|build 2>): <@124>\n- <https://github.com/sourcegraph/sourcegraph/commit/a|a> (<https://sourcegraph.com|build 3>): <@123>\n\nThe branch will automatically be unlocked once a green build has run on `main`.\nPlease head over to #buildkite-main for relevant discussion about this branch lock.\n:bulb: First time being mentioned by this bot? :point_right: <https://handbook.sourcegraph.com/departments/product-engineering/engineering/process/incidents/playbooks/ci/#build-has-failed-on-the-main-branch|Follow this step by step guide!>.\n\nFor more, refer to the <https://handbook.sourcegraph.com/departments/product-engineering/engineering/process/incidents/playbooks/ci|CI incident playbook> for help.\n\nIf unable to resolve the issue, please start an incident with the '/incident' Slack command.")
		want.Equal(t, got)
	})
}
