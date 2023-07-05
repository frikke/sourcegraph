package uploads

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/sourcegraph/sourcegraph/internal/honey"
	"github.com/sourcegraph/sourcegraph/internal/metrics"
	"github.com/sourcegraph/sourcegraph/internal/observation"
)

type operations struct {
	// Commits
	getOldestCommitDate       *observation.Operation
	getCommitsVisibleToUpload *observation.Operation
	getStaleSourcedCommits    *observation.Operation
	getCommitGraphMetadata    *observation.Operation
	updateSourcedCommits      *observation.Operation
	deleteSourcedCommits      *observation.Operation

	// Repositories
	getRepoName                             *observation.Operation
	getRepositoriesForIndexScan             *observation.Operation
	getRepositoriesMaxStaleAge              *observation.Operation
	getDirtyRepositories                    *observation.Operation
	getRecentUploadsSummary                 *observation.Operation
	getLastUploadRetentionScanForRepository *observation.Operation
	setRepositoryAsDirty                    *observation.Operation
	updateDirtyRepositories                 *observation.Operation
	setRepositoriesForRetentionScan         *observation.Operation

	// Uploads
	getUploads                        *observation.Operation
	getUploadByID                     *observation.Operation
	getUploadsByIDs                   *observation.Operation
	getVisibleUploadsMatchingMonikers *observation.Operation
	getUploadDocumentsForPath         *observation.Operation
	updateUploadsVisibleToCommits     *observation.Operation
	updateUploadRetention             *observation.Operation
	backfillReferenceCountBatch       *observation.Operation
	updateUploadsReferenceCounts      *observation.Operation
	softDeleteExpiredUploads          *observation.Operation
	deleteUploadsWithoutRepository    *observation.Operation
	deleteUploadsStuckUploading       *observation.Operation
	hardDeleteUploads                 *observation.Operation
	deleteUploadByID                  *observation.Operation
	inferClosestUploads               *observation.Operation
	backfillCommittedAtBatch          *observation.Operation

	// Dumps
	findClosestDumps                   *observation.Operation
	findClosestDumpsFromGraphFragment  *observation.Operation
	getDumpsWithDefinitionsForMonikers *observation.Operation
	getDumpsByIDs                      *observation.Operation

	// Packages
	updatePackages *observation.Operation

	// References
	updatePackageReferences *observation.Operation
	referencesForUpload     *observation.Operation

	// Audit Logs
	getAuditLogsForUpload *observation.Operation
	deleteOldAuditLogs    *observation.Operation

	// Tags
	getListTags *observation.Operation

	// Worker metrics
	uploadProcessor *observation.Operation
	uploadSizeGuage prometheus.Gauge
}

func newOperations(observationContext *observation.Context) *operations {
	m := metrics.NewREDMetrics(
		observationContext.Registerer,
		"codeintel_uploads",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of method invocations."),
	)

	op := func(name string) *observation.Operation {
		return observationContext.Operation(observation.Op{
			Name:              fmt.Sprintf("codeintel.uploads.%s", name),
			MetricLabelValues: []string{name},
			Metrics:           m,
		})
	}

	honeyObservationContext := *observationContext
	honeyObservationContext.HoneyDataset = &honey.Dataset{Name: "codeintel-worker"}
	uploadProcessor := honeyObservationContext.Operation(observation.Op{
		Name: "codeintel.uploadHandler",
		ErrorFilter: func(err error) observation.ErrorFilterBehaviour {
			return observation.EmitForTraces | observation.EmitForHoney
		},
	})

	uploadSizeGuage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "src_codeintel_upload_processor_upload_size",
		Help: "The combined size of uploads being processed at this instant by this worker.",
	})
	observationContext.Registerer.MustRegister(uploadSizeGuage)

	return &operations{
		// Commits
		getOldestCommitDate:       op("GetOldestCommitDate"),
		getCommitsVisibleToUpload: op("GetCommitsVisibleToUpload"),
		getStaleSourcedCommits:    op("GetStaleSourcedCommits"),
		getCommitGraphMetadata:    op("GetCommitGraphMetadata"),
		updateSourcedCommits:      op("UpdateSourcedCommits"),
		deleteSourcedCommits:      op("DeleteSourcedCommits"),

		// Repositories
		getRepoName:                             op("GetRepoName"),
		getRepositoriesForIndexScan:             op("GetRepositoriesForIndexScan"),
		getRepositoriesMaxStaleAge:              op("GetRepositoriesMaxStaleAge"),
		getDirtyRepositories:                    op("GetDirtyRepositories"),
		getRecentUploadsSummary:                 op("GetRecentUploadsSummary"),
		getLastUploadRetentionScanForRepository: op("GetLastUploadRetentionScanForRepository"),
		setRepositoryAsDirty:                    op("SetRepositoryAsDirty"),
		updateDirtyRepositories:                 op("UpdateDirtyRepositories"),
		setRepositoriesForRetentionScan:         op("SetRepositoriesForRetentionScan"),

		// Uploads
		getUploads:                        op("GetUploads"),
		getUploadByID:                     op("GetUploadByID"),
		getUploadsByIDs:                   op("GetUploadsByIDs"),
		getVisibleUploadsMatchingMonikers: op("GetVisibleUploadsMatchingMonikers"),
		getUploadDocumentsForPath:         op("GetUploadDocumentsForPath"),
		updateUploadsVisibleToCommits:     op("UpdateUploadsVisibleToCommits"),
		updateUploadRetention:             op("UpdateUploadRetention"),
		backfillReferenceCountBatch:       op("BackfillReferenceCountBatch"),
		updateUploadsReferenceCounts:      op("UpdateUploadsReferenceCounts"),
		deleteUploadsWithoutRepository:    op("DeleteUploadsWithoutRepository"),
		deleteUploadsStuckUploading:       op("DeleteUploadsStuckUploading"),
		softDeleteExpiredUploads:          op("SoftDeleteExpiredUploads"),
		hardDeleteUploads:                 op("HardDeleteUploads"),
		deleteUploadByID:                  op("DeleteUploadByID"),
		inferClosestUploads:               op("InferClosestUploads"),
		backfillCommittedAtBatch:          op("BackfillCommittedAtBatch"),

		// Dumps
		findClosestDumps:                   op("FindClosestDumps"),
		findClosestDumpsFromGraphFragment:  op("FindClosestDumpsFromGraphFragment"),
		getDumpsWithDefinitionsForMonikers: op("GetDumpsWithDefinitionsForMonikers"),
		getDumpsByIDs:                      op("GetDumpsByIDs"),

		// Packages
		updatePackages: op("UpdatePackages"),

		// References
		updatePackageReferences: op("UpdatePackageReferences"),
		referencesForUpload:     op("ReferencesForUpload"),

		// Audit Logs
		getAuditLogsForUpload: op("GetAuditLogsForUpload"),
		deleteOldAuditLogs:    op("DeleteOldAuditLogs"),

		// Tags
		getListTags: op("GetListTags"),

		// Worker metrics
		uploadProcessor: uploadProcessor,
		uploadSizeGuage: uploadSizeGuage,
	}
}