package afs

import "io/fs"

// Options ...
type Options struct {
	Permission fs.FileMode
}

// FileIO ...
type FileIO interface {
	Path() Path
}
