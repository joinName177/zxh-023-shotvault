package vault

import (
	"context"
	"path/filepath"
	"testing"
)

func TestExportJobCompletesAndRemovesStage(t *testing.T) {
	staging := NewStagingManager(filepath.Join(t.TempDir(), "staging"))
	job := NewExportJob(staging)
	if err := job.Run(context.Background(), "opening-hero", []byte("frame-data"), func([]byte) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if staging.Pending("opening-hero") {
		t.Fatal("completed export retained a staging record")
	}
}
