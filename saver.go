package baraka

// Saver is a interface for determine which data type to store
// Parse interface functions returns this interface
type Saver interface {
	Save(prefix string, path string, excludedContentTypes ...string) error
}
