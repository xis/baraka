package baraka

// Filter is a interface which wraps the Filter function
// you can create your own filters with this
type Filter interface {
	Filter(part *Part) bool
}

// ExtensionFilter is a filter which filters the unwanted extensions
// passes the part if the part's extension is in the extensions field
type ExtensionFilter struct {
	extensions []string
}

// NewExtensionFilter creates a new extension filter
func NewExtensionFilter(extensions ...string) Filter {
	return &ExtensionFilter{
		extensions: extensions,
	}
}

// Filter function filters the part with it's extension
func (f *ExtensionFilter) Filter(part *Part) bool {
	if part.Extension == "" {
		return false
	}

	for _, validExtension := range f.extensions {
		if part.Extension == validExtension {
			return true
		}
	}

	return false
}
