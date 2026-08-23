package vault

// Package vault owns storyboard archive behavior, storage boundaries, and desktop presentation queries.
// It keeps edits, annotations, previews, and durable snapshots behind narrow service APIs.
func PackageName() string { return "shot-vault" }
func PackageFeatures() []string {
	return []string{"durable-shot-archives", "asset-import", "preview-frames", "safe-snapshots", "event-log"}
}
