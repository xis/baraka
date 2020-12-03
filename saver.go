package baraka

// Saver is a interface that wraps the Save function
type Saver interface {
	Save(prefix string, path string, excludedContentTypes ...string) error
}
