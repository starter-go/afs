package afs

import (
	"io"
	"io/fs"
)

// Options ...
type Options struct {
	Permission fs.FileMode
}

// FileIO ...
type FileIO interface {
	Path() Path

	OpenReader(op Options) (io.ReadCloser, error)

	OpenWriter(op Options) (io.WriteCloser, error)
}
