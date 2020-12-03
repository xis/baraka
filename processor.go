package baraka

// Processor is the interface that wraps the Saver and the Informer interfaces.
type Processor interface {
	Saver
	Informer
}
