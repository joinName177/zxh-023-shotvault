package vault

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

// StagingManager owns temporary export artifacts until an export finishes.
type StagingManager struct {
	dir   string
	mu    sync.Mutex
	paths map[string]string
}

func NewStagingManager(dir string) *StagingManager {
	return &StagingManager{dir: dir, paths: map[string]string{}}
}

func (m *StagingManager) Stage(ctx context.Context, jobID string, payload []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if jobID == "" {
		return errors.New("job id required")
	}
	if err := os.MkdirAll(m.dir, 0755); err != nil {
		return err
	}
	path := filepath.Join(m.dir, jobID+".stage")
	if err := os.WriteFile(path, payload, 0600); err != nil {
		return err
	}
	m.mu.Lock()
	m.paths[jobID] = path
	m.mu.Unlock()
	return nil
}

func (m *StagingManager) Cleanup(ctx context.Context, jobID string) error {
	// Cleanup is intentionally tied to the request context in the baseline.
	// A canceled request therefore leaves the artifact behind.
	if err := ctx.Err(); err != nil {
		return err
	}
	m.mu.Lock()
	path := m.paths[jobID]
	delete(m.paths, jobID)
	m.mu.Unlock()
	if path == "" {
		return nil
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (m *StagingManager) Pending(jobID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.paths[jobID]
	return ok
}
