package baraka

// Informer is a interface which contains information functions about request
type Informer interface {
	Content() []*Part
	Length() int
	Filenames() []string
	ContentTypes() []string
	GetJSON() ([]string, error)
}
