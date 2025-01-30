package support

import (
	"io/fs"
	"os"
	"time"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/support/platforms"
)

type myCommonFileInfo struct {
	context *myFSContext
	path    afs.Path
	err     error
	exists  bool
	info    fs.FileInfo
}

func (inst *myCommonFileInfo) _Impl() afs.FileInfo {
	return inst
}

func (inst *myCommonFileInfo) load(path afs.Path) error {
	info, err := os.Stat(path.String())
	if err == nil && info != nil {
		inst.exists = true
	} else {
		inst.exists = os.IsExist(err)
		info = inst.makeEmptyInfo()
	}
	inst.err = err
	inst.info = info
	inst.path = path
	return nil
}

func (inst *myCommonFileInfo) makeEmptyInfo() fs.FileInfo {
	return &myEmptyFileInfo{}
}

func (inst *myCommonFileInfo) Path() afs.Path {
	return inst.path
}

func (inst *myCommonFileInfo) Length() int64 {
	if !inst.exists {
		return 0
	}
	return inst.info.Size()
}

func (inst *myCommonFileInfo) CreatedAt() time.Time {
	if !inst.exists {
		return time.Unix(0, 0)
	}
	info := inst.info
	return platforms.CreatedAt(info)
}

func (inst *myCommonFileInfo) AccessedAt() time.Time {
	return inst.UpdatedAt()
}

func (inst *myCommonFileInfo) UpdatedAt() time.Time {
	if !inst.exists {
		return time.Unix(0, 0)
	}
	return inst.info.ModTime()
}

func (inst *myCommonFileInfo) Mode() fs.FileMode {
	return inst.info.Mode()
}

func (inst *myCommonFileInfo) Exists() bool {
	return inst.exists
}

func (inst *myCommonFileInfo) IsFile() bool {
	if !inst.exists {
		return false
	}
	return !inst.info.IsDir()
}

func (inst *myCommonFileInfo) IsDirectory() bool {
	if !inst.exists {
		return false
	}
	return inst.info.IsDir()
}

func (inst *myCommonFileInfo) UID() int {
	return 0
}

func (inst *myCommonFileInfo) GID() int {
	return 0
}

////////////////////////////////////////////////////////////////////////////////

type myEmptyFileInfo struct {
}

func (inst *myEmptyFileInfo) _Impl() fs.FileInfo {
	return inst
}

func (inst *myEmptyFileInfo) Name() string {
	return ""
}

func (inst *myEmptyFileInfo) Size() int64 {
	return 0
}

func (inst *myEmptyFileInfo) Mode() fs.FileMode {
	return 0
}

func (inst *myEmptyFileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

func (inst *myEmptyFileInfo) IsDir() bool {
	return false
}

func (inst *myEmptyFileInfo) Sys() any {
	return inst
}
