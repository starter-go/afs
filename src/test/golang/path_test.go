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

func TestCopyTo(t *testing.T) {

	const text1 = "lsptu9a7ur0cyw4tiarnq03rnc0q2rq"

	tmp := files.FS().NewPath(t.TempDir())
	file1 := tmp.GetChild("f1.txt")
	file2 := tmp.GetChild("a/b/c/f2.txt")

	err := file1.GetIO().WriteText(text1, &afs.Options{Create: true})
	if err != nil {
		t.Error(err)
		return
	}

	err = file1.CopyTo(file2, &afs.Options{Mkdirs: true, Create: true})
	if err != nil {
		t.Error(err)
		return
	}

	text2, err := file2.GetIO().ReadText(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if text1 != text2 {
		t.Errorf("bad input & output text, want:%v have:%v", text1, text2)
	}
}

func TestMoveTo(t *testing.T) {

	const text1 = "arnq03rnc0q2rqlsptu9a7ur0cyw4ti"

	tmp := files.FS().NewPath(t.TempDir())
	file1 := tmp.GetChild("f1.txt")
	file2 := tmp.GetChild("a/b/c/f2.txt")

	err := file1.GetIO().WriteText(text1, nil)
	if err != nil {
		t.Error(err)
		return
	}

	file2.GetParent().Mkdirs(nil)

	err = file1.MoveTo(file2, nil)
	if err != nil {
		t.Error(err)
		return
	}

	text2, err := file2.GetIO().ReadText(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if text1 != text2 {
		t.Errorf("bad input & output text, want:%v have:%v", text1, text2)
	}
}

func TestRoot(t *testing.T) {

	// root dir
	path := files.FS().NewPath("////")
	if path.IsDirectory() {
		list := path.ListNames()
		for _, name := range list {
			t.Log("  find item: ", name)
		}
	}

	// roots
	roots := files.FS().ListRoots()
	for _, root := range roots {
		t.Log("  find root: ", root.GetPath())
	}

}
