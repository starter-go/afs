package afs

// Service ... 根据传入的 URI，查找相应的 Path
type Service interface {
	Fetch(uri URI) (Path, error)
}
