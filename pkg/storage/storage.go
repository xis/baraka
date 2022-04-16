package storage

type Storage interface {
	Store(path string, data []byte) error
}
