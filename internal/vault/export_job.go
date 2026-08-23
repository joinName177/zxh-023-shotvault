package vault

import (
	"context"
	"fmt"
)

type ExportSink func([]byte) error

// ExportJob stages a snapshot before handing it to the desktop exporter.
type ExportJob struct {
	staging *StagingManager
}

func NewExportJob(staging *StagingManager) *ExportJob { return &ExportJob{staging: staging} }

func (j *ExportJob) Run(ctx context.Context, jobID string, payload []byte, sink ExportSink) (err error) {
	if err = j.staging.Stage(ctx, jobID, payload); err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if cleanupErr := j.staging.Cleanup(ctx, jobID); cleanupErr != nil {
				err = fmt.Errorf("cleanup export %s: %w", jobID, cleanupErr)
			}
		}
	}()
	if err = ctx.Err(); err != nil {
		return err
	}
	if sink == nil {
		return fmt.Errorf("export %s: sink required", jobID)
	}
	if err = sink(payload); err != nil {
		return fmt.Errorf("export %s: %w", jobID, err)
	}
	return nil
}
