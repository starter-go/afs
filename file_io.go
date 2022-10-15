package afs

import (
	"io"
	"io/fs"
)

// Options ...
type Options struct {
	Permission fs.FileMode

	Mkdirs bool

	Create bool
}

// FileIO ...
type FileIO interface {
	Path() Path

	OpenReader(opt *Options) (io.ReadCloser, error)

	OpenWriter(opt *Options) (io.WriteCloser, error)

	WriteText(text string, opt *Options) error

	WriteBinary(b []byte, opt *Options) error

	ReadText(opt *Options) (string, error)

	ReadBinary(opt *Options) ([]byte, error)
}
