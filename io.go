package afs

import "io"

// // ReadSeekCloser ...
// type ReadSeekCloser interface {
// 	io.Closer
// 	io.ReadSeeker
// }

// WriteSeekCloser ...
type WriteSeekCloser interface {
	io.Closer
	io.WriteSeeker
}

// ReadWriteSeekCloser ...
type ReadWriteSeekCloser interface {
	io.Closer
	io.ReadWriteSeeker
}
