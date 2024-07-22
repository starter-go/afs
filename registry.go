package afs

// Registration ...
type Registration struct {
	Name     string
	Enabled  bool
	Priority int
	Driver   Driver
}

// Registry ...
type Registry interface {
	Registration() *Registration
}
