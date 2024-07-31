package support

import (
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/support/platforms"
)

// GetWindowsFS ...
func GetWindowsFS() afs.FS {
	ctx := &myFSContext{}
	ctx.shell = &myShellFS{context: ctx}
	ctx.common = &CommonFileSystemCore{context: ctx}
	ctx.platform = &myWindowsFS{context: ctx}
	return ctx.shell
}

////////////////////////////////////////////////////////////////////////////////

type myWindowsFS struct {
	context *myFSContext
}

func (inst *myWindowsFS) _Impl() PlatformFileSystem {
	return inst
}

func (inst *myWindowsFS) GetFS() afs.FS {
	return inst.context.shell
}

func (inst *myWindowsFS) New(path string) afs.Path {
	const pathLenLimit = 8
	l := len(path)
	if l < pathLenLimit {
		p2, err := inst.NormalizePath(path)
		if err == nil {
			if p2 == "" || p2 == "/" || p2 == "//" || p2 == "///" {
				return inst.loadWindowsVRootDir()
			}
		}
	}
	return inst.context.common.New(path)
}

func (inst *myWindowsFS) loadWindowsVRootDir() afs.Path {
	loader := &myWindowsVRootDirLoader{context: inst.context}
	return loader.load()
}

func (inst *myWindowsFS) NormalizePath(path string) (string, error) {
	comm := inst.context.common
	sep := inst.Separator()
	elements := comm.PathToElements(path)
	elements, err := comm.NormalizePathElements(elements)
	if err != nil {
		return "", err
	}
	path = comm.ElementsToPath(elements, "", sep)
	if strings.HasSuffix(path, ":") {
		if !strings.Contains(path, sep) {
			return path + sep, nil
		}
	}
	return path, nil
}

func (inst *myWindowsFS) PathSeparator() string {
	return ";"
}

func (inst *myWindowsFS) Separator() string {
	return "\\"
}

// func (inst *myWindowsFS) isRootExists(path string) bool {
// 	_, err := os.Stat(path)
// 	return err == nil || os.IsExist(err)
// }

// func (inst *myWindowsFS) innerListRoots() ([]string, error) {
// }

func (inst *myWindowsFS) ListRoots() []string {
	return platforms.Roots()
}

func (inst *myWindowsFS) GetCommonFileSystem() CommonFileSystem {
	return inst.context.common
}
