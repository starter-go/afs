package support

import "bitwormhole.com/starter/afs"

type myCommonFileIO struct {
	path afs.Path
}

func (inst *myCommonFileIO) _Impl() afs.FileIO {
	return inst
}

func (inst *myCommonFileIO) Path() afs.Path {
	return inst.path
}
