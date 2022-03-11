package graphqlbackend

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/inconshreveable/log15"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/backend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/envvar"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
	"github.com/sourcegraph/sourcegraph/internal/httpcli"
	"github.com/sourcegraph/sourcegraph/internal/repos"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

var (
	cf = httpcli.ExternalClientFactory
)

// type codeHostResult struct {
// 	svcID  int64
// 	repos  []*codeHostRepositoryResolver
// 	svcErr error
// }

type affiliatedRepositoriesConnection struct {
	userID   int32
	orgID    int32
	codeHost int64
	query    string
	once     sync.Once
	nodes    []*codeHostRepositoryResolver
	err      error
	db       database.DB
}

func (a *affiliatedRepositoriesConnection) Nodes(ctx context.Context) ([]*codeHostRepositoryResolver, error) {
	a.once.Do(func() {
		var (
			svcs []*types.ExternalService
			err  error
		)
		// get all external services for the user, the organization, or for the specified external service
		if a.codeHost == 0 {
			svcs, err = a.db.ExternalServices().List(ctx, database.ExternalServicesListOptions{
				NamespaceUserID: a.userID,
				NamespaceOrgID:  a.orgID,
			})
			if err != nil {
				a.err = err
				return
			}
			fmt.Println("===1", len(svcs))

		} else {
			fmt.Println("===2")
			svc, err := a.db.ExternalServices().GetByID(ctx, a.codeHost)
			if err != nil {
				a.err = err
				return
			}
			// 🚨 SECURITY: check if user can access external service
			err = backend.CheckExternalServiceAccess(ctx, a.db, svc.NamespaceUserID, svc.NamespaceOrgID)
			if err != nil {
				a.err = err
				return
			}
			svcs = append(svcs, svc)
		}

		type affiliatedResult struct {
			svcID int64
			repos []types.CodeHostRepository
			err   error
		}

		fmt.Println("===3")

		// get Source for all external services
		var (
			results  = make(chan affiliatedResult, len(svcs))
			svcsByID = make(map[int64]*types.ExternalService)
			pending  int
		)
		for i, svc := range svcs {
			svcsByID[svc.ID] = svc
			fmt.Println("iterating over svcs", i)
			src, err := repos.NewSource(a.db.ExternalServices(), svc, cf)
			if err != nil {
				a.err = err
				return
			}
			af, ok := src.(repos.AffiliatedRepositorySource)
			if !ok {
				continue
			}
			pending++
			svcID := svc.ID
			goroutine.Go(func() {
				affiliated, err := af.AffiliatedRepositories(ctx)
				results <- affiliatedResult{
					svcID: svcID,
					repos: affiliated,
					err:   err,
				}
			})
		}

		// are we allowed to show the private repos?
		allowPrivate, err := allowPrivate(ctx, a.db, a.userID, a.orgID)
		if err != nil {
			a.err = err
			return
		}

		// collect all results
		var fetchErrors []error
		a.nodes = []*codeHostRepositoryResolver{}
		for i := 0; i < pending; i++ {
			select {
			case result := <-results:
				if result.err != nil {
					fmt.Println("result.err not nil", result.err)
					// An error from one code is not fatal
					log15.Error("getting affiliated repos", "externalServiceId", result.svcID, "err", result.err)
					fetchErrors = append(fetchErrors, errors.New("Error from "+svcsByID[result.svcID].DisplayName+" - "+result.err.Error()))
					// fetchErrors = append(fetchErrors, result.err)
					continue
				}

				fmt.Println("resul errrs len", result.err, result.repos, result.svcID)

				for _, repo := range result.repos {
					if a.query != "" && !strings.Contains(strings.ToLower(repo.Name), a.query) {
						continue
					}
					if !allowPrivate && repo.Private {
						continue
					}
					repo := repo

					a.nodes = append(a.nodes, &codeHostRepositoryResolver{
						db:       a.db,
						codeHost: svcsByID[repo.CodeHostID],
						repo:     &repo,
					})
				}
			case <-ctx.Done():
				a.err = ctx.Err()
				return
			}
		}

		fmt.Println("=== /////////// 5")

		sort.Slice(a.nodes, func(i, j int) bool {
			return a.nodes[i].repo.Name < a.nodes[j].repo.Name
		})

		if len(fetchErrors) == pending {
			// All hosts failed
			fmt.Println("heeeere")
			a.nodes = nil
			a.err = errors.New("failed to fetch from any code host")
		}

		if len(fetchErrors) > 0 {
			var allErrors []string
			for _, e := range fetchErrors {
				allErrors = append(allErrors, e.Error())
			}
			// a.nodes = nil
			a.err = errors.New(strings.Join(allErrors, ","))
		}
	})

	if envvar.SourcegraphDotComMode() && a.orgID != 0 {
		a.db.OrgStats().Upsert(ctx, a.orgID, int32(len(a.nodes)))
	}

	// If we return an error in the next line, the value (a.nodes) will be ignored . Do we want to show the partial result?
	return a.nodes, a.err
}

type codeHostRepositoryResolver struct {
	repo     *types.CodeHostRepository
	codeHost *types.ExternalService
	db       database.DB
}

func (r *codeHostRepositoryResolver) Name() string {
	return r.repo.Name
}

func (r *codeHostRepositoryResolver) Private() bool {
	return r.repo.Private
}

func (r *codeHostRepositoryResolver) CodeHost(ctx context.Context) *externalServiceResolver {
	return &externalServiceResolver{
		db:              r.db,
		externalService: r.codeHost,
	}
}

func allowPrivate(ctx context.Context, db database.DB, userID, orgID int32) (bool, error) {
	if userID > 0 {
		mode, err := db.Users().UserAllowedExternalServices(ctx, userID)
		if err != nil {
			return false, err
		}
		return mode == conf.ExternalServiceModeAll, nil
	}
	if orgID > 0 {
		return true, nil
	}

	return false, errors.New("either userID or orgID expected to be defined")
}
