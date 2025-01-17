package codeintel

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"

	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/internal/codeintel/autoindexing"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/codenav"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/policies"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/stores"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/stores/gitserver"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/stores/lsifuploadstore"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/stores/repoupdater"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/uploads"
	uploadshttp "github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/transport/http"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/conf/conftypes"
	"github.com/sourcegraph/sourcegraph/internal/database"
	connections "github.com/sourcegraph/sourcegraph/internal/database/connections/live"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/trace"
)

type Services struct {
	// shared with executor queue
	InternalUploadHandler http.Handler
	ExternalUploadHandler http.Handler

	gitserverClient *gitserver.Client

	// used by resolvers
	AutoIndexingSvc *autoindexing.Service
	UploadsSvc      *uploads.Service
	CodeNavSvc      *codenav.Service
	PoliciesSvc     *policies.Service
	UploadSvc       *uploads.Service
}

func NewServices(ctx context.Context, config *Config, siteConfig conftypes.WatchableSiteConfig, db database.DB) (*Services, error) {
	// Initialize tracing/metrics
	logger := log.Scoped("codeintel", "codeintel services")
	observationContext := &observation.Context{
		Logger:     logger,
		Tracer:     &trace.Tracer{TracerProvider: otel.GetTracerProvider()},
		Registerer: prometheus.DefaultRegisterer,
	}

	// Connect to the separate LSIF database
	codeIntelDBConnection := mustInitializeCodeIntelDB(logger)

	// Initialize lsif stores (TODO: these should be integrated, they are basically pointing to the same thing)
	codeIntelLsifStore := database.NewDBWith(observationContext.Logger, codeIntelDBConnection)
	uploadStore, err := lsifuploadstore.New(context.Background(), config.LSIFUploadStoreConfig, observationContext)
	if err != nil {
		return nil, err
	}

	// Initialize gitserver client & repoupdater
	gitserverClient := gitserver.New(db, observationContext)
	repoUpdaterClient := repoupdater.New(observationContext)

	// Initialize services
	uploadSvc := uploads.GetService(db, codeIntelLsifStore, gitserverClient)
	codenavSvc := codenav.GetService(db, codeIntelLsifStore, uploadSvc, gitserverClient)
	policySvc := policies.GetService(db, uploadSvc, gitserverClient)
	autoindexingSvc := autoindexing.GetService(db, uploadSvc, gitserverClient, repoUpdaterClient)

	// Initialize http endpoints
	newUploadHandler := func(withCodeHostAuth bool) http.Handler {
		return uploadshttp.GetHandler(uploadSvc, db, uploadStore, withCodeHostAuth)
	}
	internalUploadHandler := newUploadHandler(false)
	externalUploadHandler := newUploadHandler(true)

	return &Services{
		InternalUploadHandler: internalUploadHandler,
		ExternalUploadHandler: externalUploadHandler,

		gitserverClient: gitserverClient,

		AutoIndexingSvc: autoindexingSvc,
		UploadsSvc:      uploadSvc,
		CodeNavSvc:      codenavSvc,
		PoliciesSvc:     policySvc,
		UploadSvc:       uploadSvc,
	}, nil
}

func mustInitializeCodeIntelDB(logger log.Logger) stores.CodeIntelDB {
	dsn := conf.GetServiceConnectionValueAndRestartOnChange(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.CodeIntelPostgresDSN
	})
	db, err := connections.EnsureNewCodeIntelDB(dsn, "frontend", &observation.TestContext)
	if err != nil {
		logger.Fatal("Failed to connect to codeintel database", log.Error(err))
	}
	return stores.NewCodeIntelDB(db)
}
