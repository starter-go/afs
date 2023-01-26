package support

import (
	"bitwormhole.com/starter/afs"
)

// GetPosixFS ...
func GetPosixFS() afs.FS {
	ctx := &myFSContext{}
	ctx.shell = &myShellFS{context: ctx}
	ctx.common = &CommonFileSystemCore{context: ctx}
	ctx.platform = &myPosixFS{context: ctx}
	return ctx.shell
}

// core
type myPosixFS struct {
	context *myFSContext
}

func (inst *myPosixFS) _Impl() PlatformFileSystem {
	return inst
}

func (inst *myPosixFS) GetFS() afs.FS {
	return inst.context.shell
}

func (inst *myPosixFS) GetCommonFileSystem() CommonFileSystem {
	return inst.context.common
}

func (inst *myPosixFS) NormalizePath(path string) (string, error) {
	comm := inst.context.common
	sep := inst.Separator()
	elements := comm.PathToElements(path)
	elements, err := comm.NormalizePathElements(elements)
	if err != nil {
		return "", err
	}
	path = comm.ElementsToPath(elements, sep, sep)
	return path, nil
}

func (inst *myPosixFS) PathSeparator() string {
	return ":"
}

func (inst *myPosixFS) Separator() string {
	return "/"
}

func (inst *myPosixFS) ListRoots() []string {
	return []string{"/"}
}
