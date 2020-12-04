package baraka

// Informer is a interface that wraps information functions about parsed multipart request
type Informer interface {
	Content() []*Part
	Length() int
	Filenames() []string
	ContentTypes() []string
	GetJSON() ([]string, error)
}
