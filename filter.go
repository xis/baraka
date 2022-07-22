package baraka

// Filter is a interface which wraps the Filter function
// you can create your own filters with this
type Filter interface {
	Filter(data []byte) (bool, error)
}
