package support

import (
	"io"
	"os"

	"bitwormhole.com/starter/afs"
)

type myCommonFileIO struct {
	path afs.Path
}

func (inst *myCommonFileIO) _Impl() afs.FileIO {
	return inst
}

func (inst *myCommonFileIO) Path() afs.Path {
	return inst.path
}

func (inst *myCommonFileIO) OpenReader(op afs.Options) (io.ReadCloser, error) {
	path := inst.path.GetPath()
	flag := os.O_RDONLY
	return os.OpenFile(path, flag, op.Permission)
}

func (inst *myCommonFileIO) OpenWriter(op afs.Options) (io.WriteCloser, error) {
	path := inst.path.GetPath()
	flag := os.O_CREATE | os.O_WRONLY
	return os.OpenFile(path, flag, op.Permission)
}
