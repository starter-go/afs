package afs

import "io"

// ReaderPool ...
type ReaderPool interface {
	io.Closer
	Clean()
	OpenReader(file Path, op *Options) (io.ReadSeekCloser, error)
}

////////////////////////////////////////////////////////////////////////////////

// NopReaderPool ...
type NopReaderPool struct {
}

func (inst *NopReaderPool) _Impl() ReaderPool {
	return inst
}

// Clean ...
func (inst *NopReaderPool) Clean() {
}

// OpenReader ...
func (inst *NopReaderPool) OpenReader(file Path, op *Options) (io.ReadSeekCloser, error) {
	return file.GetIO().OpenSeekerR(op)
}

// Close ...
func (inst *NopReaderPool) Close() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
