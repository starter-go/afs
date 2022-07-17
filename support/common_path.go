package support

import (
	"os"
	"strings"

	"bitwormhole.com/starter/afs"
)

type myCommonPath struct {
	context *myFSContext
	path    string
}

func (inst *myCommonPath) _Impl() afs.Path {
	return inst
}

func (inst *myCommonPath) GetFS() afs.FS {
	return inst.context.shell
}

func (inst *myCommonPath) GetParent() afs.Path {
	path := inst.path
	return inst.GetFS().NewPath(path + "/..")
}

func (inst *myCommonPath) GetChild(name string) afs.Path {
	path := inst.path
	return inst.GetFS().NewPath(path + "/" + name)
}

func (inst *myCommonPath) String() string {
	return inst.path
}

func (inst *myCommonPath) GetName() string {

	sep := inst.context.platform.Separator()
	path := inst.path

	// fast method
	index := strings.LastIndex(path, sep)
	if index > 0 {
		name := path[index+1:]
		name = strings.TrimSpace(name)
		if len(name) > 0 {
			return name
		}
	}

	// full method
	elements := inst.context.common.PathToElements(path)
	for i := len(elements) - 1; i >= 0; i-- {
		el := elements[i]
		el = strings.TrimSpace(el)
		if len(el) > 0 {
			return el
		}
	}

	return ""
}

func (inst *myCommonPath) GetPath() string {
	return inst.path
}

func (inst *myCommonPath) GetInfo() afs.FileInfo {
	info := &myCommonFileInfo{}
	info.load(inst)
	return info
}

func (inst *myCommonPath) GetIO() afs.FileIO {
	return &myCommonFileIO{path: inst}
}

func (inst *myCommonPath) Mkdir(op afs.Options) error {
	path := inst.path
	return os.Mkdir(path, op.Permission)
}

func (inst *myCommonPath) Mkdirs(op afs.Options) error {
	path := inst.path
	return os.MkdirAll(path, op.Permission)
}

func (inst *myCommonPath) Delete() error {
	return os.Remove(inst.path)
}

func (inst *myCommonPath) ListNames() []string {
	file, err := os.Open(inst.path)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	names, err := file.Readdirnames(0)
	if err != nil {
		return []string{}
	}
	return names
}

func (inst *myCommonPath) ListPaths() []string {
	names := inst.ListNames()
	dst := make([]string, 0)
	for _, name := range names {
		child := inst.GetChild(name)
		dst = append(dst, child.String())
	}
	return dst
}

func (inst *myCommonPath) ListChildren() []afs.Path {
	names := inst.ListNames()
	dst := make([]afs.Path, 0)
	for _, name := range names {
		child := inst.GetChild(name)
		dst = append(dst, child)
	}
	return dst
}
