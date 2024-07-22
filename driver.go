package afs

// Driver ... 根据传入的 URI， 提供相应的 Path
type Driver interface {
	Support(uri URI) bool
	Fetch(uri URI) (Path, error)
}
