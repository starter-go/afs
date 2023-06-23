package afs

import (
	"io"
)

// FileIO ...
type FileIO interface {
	Path() Path

	OpenReader(opt *Options) (io.ReadCloser, error)

	OpenWriter(opt *Options) (io.WriteCloser, error)

	OpenSeekerR(opt *Options) (io.ReadSeekCloser, error)

	OpenSeekerW(opt *Options) (WriteSeekCloser, error)

	OpenSeekerRW(opt *Options) (ReadWriteSeekCloser, error)

	// text

	ReadText(opt *Options) (string, error)

	WriteText(text string, opt *Options) error

	// binary

	WriteBinary(b []byte, opt *Options) error

	ReadBinary(opt *Options) ([]byte, error)
}
