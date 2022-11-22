package afs

import (
	"io"
	"io/fs"
)

// Options ...
type Options struct {

	// file with fs.ModeXXX
	Permission fs.FileMode

	// fill with os.O_xxxx
	Flag int

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
