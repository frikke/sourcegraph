package codeintel

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"

	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/cmd/worker/job"
	"github.com/sourcegraph/sourcegraph/cmd/worker/shared/init/codeintel"
	workerdb "github.com/sourcegraph/sourcegraph/cmd/worker/shared/init/db"

	"github.com/sourcegraph/sourcegraph/internal/codeintel/policies"
	policiesEnterprise "github.com/sourcegraph/sourcegraph/internal/codeintel/policies/enterprise"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/uploads"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/background/expiration"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/trace"
)

type uploadExpirerJob struct{}

func NewUploadExpirerJob() job.Job {
	return &uploadExpirerJob{}
}

func (j *uploadExpirerJob) Description() string {
	return ""
}

func (j *uploadExpirerJob) Config() []env.Config {
	return []env.Config{
		expiration.ConfigInst,
	}
}

func (j *uploadExpirerJob) Routines(startupCtx context.Context, logger log.Logger) ([]goroutine.BackgroundRoutine, error) {
	observationContext := &observation.Context{
		Logger:     logger.Scoped("routines", "codeintel job routines"),
		Tracer:     &trace.Tracer{TracerProvider: otel.GetTracerProvider()},
		Registerer: prometheus.DefaultRegisterer,
	}
	metrics := expiration.NewMetrics(observationContext)

	rawDB, err := workerdb.Init()
	if err != nil {
		return nil, err
	}
	db := database.NewDB(logger, rawDB)

	rawCodeIntelDB, err := codeintel.InitCodeIntelDatabase()
	if err != nil {
		return nil, err
	}
	codeIntelDB := database.NewDB(logger, rawCodeIntelDB)

	gitserverClient, err := codeintel.InitGitserverClient()
	if err != nil {
		return nil, err
	}

	policyMatcher := policiesEnterprise.NewMatcher(gitserverClient, policiesEnterprise.RetentionExtractor, true, false)

	uploadSvc := uploads.GetService(db, codeIntelDB, gitserverClient)
	policySvc := policies.GetService(db, uploadSvc, gitserverClient)

	return []goroutine.BackgroundRoutine{
		expiration.NewExpirer(uploadSvc, policySvc, policyMatcher, metrics),
		expiration.NewReferenceCountUpdater(uploadSvc),
	}, nil
}
