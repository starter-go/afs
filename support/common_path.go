package support

import (
	"bytes"
	"io"
	"os"
	"strings"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/base/util"
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

func (inst *myCommonPath) Exists() bool {
	return inst.GetInfo().Exists()
}

func (inst *myCommonPath) IsDirectory() bool {
	return inst.GetInfo().IsDirectory()
}

func (inst *myCommonPath) IsFile() bool {
	return inst.GetInfo().IsFile()
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

func (inst *myCommonPath) Mkdir(op *afs.Options) error {
	if op == nil {
		op = &afs.Options{}
	}
	path := inst.path
	return os.Mkdir(path, op.Permission)
}

func (inst *myCommonPath) Mkdirs(op *afs.Options) error {
	if op == nil {
		op = &afs.Options{}
	}
	path := inst.path
	return os.MkdirAll(path, op.Permission)
}

func (inst *myCommonPath) Delete() error {
	return os.Remove(inst.path)
}

func (inst *myCommonPath) Create(op *afs.Options) error {
	return inst.CreateWithSource(nil, op)
}

func (inst *myCommonPath) CreateWithData(data []byte, op *afs.Options) error {
	mem := &bytes.Buffer{}
	if data != nil {
		mem.Write(data)
	}
	return inst.CreateWithSource(mem, op)
}

func (inst *myCommonPath) CreateWithSource(src io.Reader, op *afs.Options) error {
	if src == nil {
		data := []byte{}
		src = bytes.NewReader(data)
	}
	dst, err := inst.GetIO().OpenWriter(op)
	if err != nil {
		return err
	}
	defer util.Close(dst)
	_, err = util.PumpStream(src, dst, nil)
	return err
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
