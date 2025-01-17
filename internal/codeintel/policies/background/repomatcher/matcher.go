package repomatcher

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
)

type matcher struct {
	policySvc PolicyService
	metrics   *metrics
}

var (
	_ goroutine.Handler      = &matcher{}
	_ goroutine.ErrorHandler = &matcher{}
)

func (m *matcher) Handle(ctx context.Context) error {
	policies, err := m.policySvc.SelectPoliciesForRepositoryMembershipUpdate(ctx, ConfigInst.ConfigurationPolicyMembershipBatchSize)
	if err != nil {
		return err
	}

	for _, policy := range policies {
		var patterns []string
		if policy.RepositoryPatterns != nil {
			patterns = *policy.RepositoryPatterns
		}

		var repositoryMatchLimit *int
		if val := conf.CodeIntelAutoIndexingPolicyRepositoryMatchLimit(); val != -1 {
			repositoryMatchLimit = &val
		}

		// Always call this even if patterns are not supplied. Otherwise we run into the
		// situation where we have deleted all of the patterns associated with a policy
		// but it still has entries in the lookup table.
		if err := m.policySvc.UpdateReposMatchingPatterns(ctx, patterns, policy.ID, repositoryMatchLimit); err != nil {
			return err
		}

		m.metrics.numPoliciesUpdated.Inc()
	}

	return nil
}

func (m *matcher) HandleError(err error) {}
