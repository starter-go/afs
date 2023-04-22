package support

import (
	"os"

	"bitwormhole.com/starter/afs"
)

type myShellFS struct {
	context *myFSContext
}

func (inst *myShellFS) _Impl() afs.FS {
	return inst
}

func (inst *myShellFS) NewPath(path string) afs.Path {
	p2, _ := inst.context.platform.NormalizePath(path)
	return &myCommonPath{context: inst.context, path: p2}
}

func (inst *myShellFS) ListRoots() []afs.Path {
	src := inst.context.platform.ListRoots()
	dst := make([]afs.Path, 0)
	for _, item := range src {
		path := inst.NewPath(item)
		dst = append(dst, path)
	}
	return dst
}

func (inst *myShellFS) CreateTempFile(prefix, suffix string, folder afs.Path) (afs.Path, error) {
	dir := ""
	pattern := prefix + "*" + suffix
	if folder != nil {
		dir = folder.String()
	}
	file, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	path := file.Name()
	return inst.NewPath(path), nil
}

// PathSeparator return ';'(windows) | ':'(unix)
func (inst *myShellFS) PathSeparator() string {
	return inst.context.platform.PathSeparator()
}

// Separator return '/'(unix) | '\'(windows)
func (inst *myShellFS) Separator() string {
	return inst.context.platform.Separator()
}

// func (inst *myShellFS) OpenReaderPool() afs.ReaderPool {
// 	return nil
// }

func (inst *myShellFS) SetDefaultOptionsHandler(h afs.OptionsHandlerFunc) error {
	return inst.context.common.SetDefaultOptionsHandler(h)
}
