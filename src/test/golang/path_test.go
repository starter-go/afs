package golang

import (
	"io/fs"
	"os"
	"testing"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/vlog"
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

	opt := &afs.Options{
		Flag:       os.O_CREATE | os.O_WRONLY,
		Permission: fs.ModePerm,
	}

	if !info.Exists() {
		data := []byte("hello, afs file")
		err := file1.CreateWithData(data, opt)
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
		// t.Log("current path = ", p.GetPath())
		vlog.Info("current path = %s", p.GetPath())
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

	opt := afs.Todo().Options()
	opt = nil

	err := file1.GetIO().WriteText(text1, opt)
	if err != nil {
		t.Error(err)
		return
	}

	file2.MakeParents(nil)
	err = file1.CopyTo(file2, opt)
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
	info := path.GetInfo()
	info.AccessedAt()
	info.CreatedAt()
	info.UpdatedAt()
	info.Mode()

	// roots
	roots := files.FS().ListRoots()
	for _, root := range roots {
		t.Log("  find root: ", root.GetPath())
	}

}

func TestNormalizePath(t *testing.T) {

	fs1 := files.FS()
	paths := make([]string, 0)

	paths = append(paths, "file:///a/b/cd")
	paths = append(paths, "~/x/y/z1")
	paths = append(paths, "c:\\i\\j\\kkk")
	paths = append(paths, "~/a/b/c/./d/e/f/g")
	paths = append(paths, "~/a/b/c/../d/e/f")
	paths = append(paths, "~/a/b/c/./////////./d/e/f")
	paths = append(paths, "c:\\i\\j\\k////file:////c:////kk")
	paths = append(paths, "c:\\i\\j\\k///////~////kk")

	for _, path := range paths {
		p1 := path
		p2 := fs1.NewPath(p1).GetPath()
		vlog.Info("test normalizePath(p1) p2")
		vlog.Info("    p1 = %s", p1)
		vlog.Info("    p2 = %s", p2)
	}
}
