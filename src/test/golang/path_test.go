package golang

import (
	"testing"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
)

func logFileInfo(info afs.FileInfo, t *testing.T) {

	path := info.Path().String()

	size := info.Length()
	time0 := info.CreatedAt()
	time1 := info.UpdatedAt()
	exists := info.Exists()
	isdir := info.IsDirectory()
	isfile := info.IsFile()
	mode := info.Mode()

	table := make(map[string]interface{}, 0)
	table["length"] = size
	table["created_at"] = time0
	table["updated_at"] = time1
	table["exists"] = exists
	table["is_dir"] = isdir
	table["is_file"] = isfile
	table["mode"] = mode

	t.Log("path = ", path)
	for k, v := range table {
		t.Log("  info.", k, " = ", v)
	}
}

func TestDirPathInfo(t *testing.T) {

	tmp := t.TempDir()
	fs1 := files.FS()
	dir1 := fs1.NewPath(tmp + "/dir1")

	info := dir1.GetInfo()
	logFileInfo(info, t)

	if !info.Exists() {
		err := dir1.Mkdirs(&afs.Options{})
		if err != nil {
			t.Error(err)
		}
	}

	info = dir1.GetInfo()
	logFileInfo(info, t)
}

func TestFilePathInfo(t *testing.T) {

	tmp := t.TempDir()
	fs1 := files.FS()
	file1 := fs1.NewPath(tmp + "/file1")

	info := file1.GetInfo()
	logFileInfo(info, t)

	if !info.Exists() {
		data := []byte("hello, afs file")
		err := file1.CreateWithData(data, &afs.Options{Create: true})
		if err != nil {
			t.Error(err)
		}
	}

	info = file1.GetInfo()
	logFileInfo(info, t)

}

func TestGetParent(t *testing.T) {
	tmp := t.TempDir()
	fs1 := files.FS()
	dir1 := fs1.NewPath(tmp + "/dir1")
	p := dir1
	timeout := 99
	for ; timeout > 0; timeout-- {
		if p == nil {
			break
		}
		t.Log("current path = ", p.GetPath())
		p = p.GetParent()
	}
	if p == nil {
		return
	}
	t.Error("timeout")
}
