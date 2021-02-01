package baraka

type Filter interface {
	Filter(part *Part) bool
}

type ExtensionFilter struct {
	extensions []string
}

func NewExtensionFilter(extensions ...string) Filter {
	return &ExtensionFilter{
		extensions: extensions,
	}
}

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
