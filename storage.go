package baraka

type Storage interface {
	Store(path string, data []byte) error
}
