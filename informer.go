package baraka

// Informer is a interface which contains information functions about request
type Informer interface {
	Length() int
	Filenames() []string
	ContentTypes() []string
}

// Length returns total count of files
func (parts *Parts) Length() int {
	return parts.len
}

// Filenames returns names of files
func (parts *Parts) Filenames() []string {
	filenames := make([]string, parts.len)
	for k, v := range parts.files {
		filenames[k] = v.Filename
	}
	return filenames
}

// ContentTypes returns content types of files
func (parts *Parts) ContentTypes() []string {
	contentTypes := make([]string, parts.len)
	for k, v := range parts.files {
		contentTypes[k] = v.Header.Get("Content-Type")
	}
	return contentTypes
}
