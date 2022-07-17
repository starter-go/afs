package support

import (
	"os"
	"strings"

	"bitwormhole.com/starter/afs"
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

func (inst *myWindowsFS) isRootExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (inst *myWindowsFS) ListRoots() []string {
	const driveA rune = 'A'
	const driveZ rune = 'Z'
	list := make([]string, 0)
	for drive := driveA; drive <= driveZ; drive++ {
		path := string(drive) + ":\\"
		if inst.isRootExists(path) {
			list = append(list, path)
		}
	}
	return list
}

func (inst *myWindowsFS) GetCommonFileSystem() CommonFileSystem {
	return inst.context.common
}
