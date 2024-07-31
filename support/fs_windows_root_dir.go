package support

import (
	"fmt"
	"io"
	"io/fs"
	"sort"

	"github.com/starter-go/afs"
)

type myWindowsVRootDirLoader struct {
	context *myFSContext
}

func (inst *myWindowsVRootDirLoader) load() afs.Path {
	roots := inst.context.platform.ListRoots()
	sort.Strings(roots)
	drvlist := make([]afs.Path, 0)
	for _, path := range roots {
		driver := inst.context.shell.NewPath(path)
		drvlist = append(drvlist, driver)
	}
	vroot := &myWindowsVRootDir{
		context: inst.context,
		drivers: drvlist,
	}
	return vroot
}

////////////////////////////////////////////////////////////////////////////////

type myWindowsVRootDir struct {
	context *myFSContext
	drivers []afs.Path
}

func (inst *myWindowsVRootDir) _impl() afs.Path {
	return inst
}

func (inst *myWindowsVRootDir) Exists() bool {
	return true
}

func (inst *myWindowsVRootDir) IsFile() bool {
	return false
}

func (inst *myWindowsVRootDir) IsDirectory() bool {
	return true
}

func (inst *myWindowsVRootDir) GetName() string {
	return "[ROOT]"
}

func (inst *myWindowsVRootDir) GetPath() string {
	return "/"
}

func (inst *myWindowsVRootDir) GetURI() afs.URI {
	return "file:///"
}

func (inst *myWindowsVRootDir) GetInfo() afs.FileInfo {
	info := &myCommonFileInfo{
		path:   inst,
		exists: true,
	}
	return info
}

func (inst *myWindowsVRootDir) String() string {
	return inst.GetPath()
}

func (inst *myWindowsVRootDir) GetFS() afs.FS {
	return inst.context.shell
}

func (inst *myWindowsVRootDir) GetParent() afs.Path {
	return nil
}

func (inst *myWindowsVRootDir) GetChild(name string) afs.Path {
	me := inst.GetPath()
	path := me + "/" + name
	sh := inst.context.shell
	return sh.NewPath(path)
}

func (inst *myWindowsVRootDir) ListNames() []string {
	src := inst.drivers
	dst := make([]string, len(src))
	for i, item := range src {
		dst[i] = item.GetName()
	}
	return dst
}

func (inst *myWindowsVRootDir) ListPaths() []string {
	src := inst.drivers
	dst := make([]string, len(src))
	for i, item := range src {
		dst[i] = item.GetPath()
	}
	return dst
}

func (inst *myWindowsVRootDir) ListChildren() []afs.Path {
	src := inst.drivers
	dst := make([]afs.Path, len(src))
	for i, item := range src {
		dst[i] = item
	}
	return dst
}

func (inst *myWindowsVRootDir) Mkdir(opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) Mkdirs(opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) MakeParents(opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) Chmod(m fs.FileMode) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) Chown(uid, gid int) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) Delete() error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) Create(opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) CreateWithData(data []byte, opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) CreateWithSource(src io.Reader, opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) MoveTo(dst afs.Path, opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) CopyTo(dst afs.Path, opt *afs.Options) error {
	return fmt.Errorf("unsupported")
}

func (inst *myWindowsVRootDir) GetIO() afs.FileIO {
	return nil
}
