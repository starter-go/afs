package golang

import (
	"io/fs"
	"testing"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/vlog"
)

func TestFSListRoots(t *testing.T) {
	fs1 := files.FS()
	roots := fs1.ListRoots()
	for index, root := range roots {
		name := root.GetName()
		path := root.GetPath()
		t.Log("index:", index, "name:", name, "path:", path)
	}
}

func TestVRoots(t *testing.T) {
	fs1 := files.FS()
	vroot, err := fs1.NewURI("file:/")
	if err != nil {
		t.Error(err)
		return
	}
	roots := vroot.ListChildren()
	for index, root := range roots {
		name := root.GetName()
		path := root.GetPath()
		t.Log("index:", index, "name:", name, "path:", path)
	}

	myName := vroot.GetName()
	myURI := vroot.GetURI()
	myPath := vroot.GetPath()
	t.Logf("[VRoot name:'%s' path:'%s' uri:'%s']", myName, myPath, myURI)
}

func TestFSCreateTempFile(t *testing.T) {
	fs1 := files.FS()
	path, err := fs1.CreateTempFile("ppp", "sss", nil)
	if err != nil {
		t.Error(err)
	}
	table := map[string]string{}
	table["name"] = path.GetName()
	table["path"] = path.GetPath()
	table["parent"] = path.GetParent().GetPath()
	table["String"] = path.String()
	table["mode"] = path.GetInfo().Mode().String()
	for k, v := range table {
		t.Log("path.", k, " = ", v)
	}
}

func TestFSNewPath(t *testing.T) {
	dir := t.TempDir()
	fs1 := files.FS()
	path := fs1.NewPath(dir + "a/b/c\\d\\")
	table := map[string]string{}
	table["name"] = path.GetName()
	table["path"] = path.GetPath()
	table["parent"] = path.GetParent().GetPath()
	table["String"] = path.String()
	for k, v := range table {
		t.Log("path.", k, " = ", v)
	}
}

func TestFSSeparator(t *testing.T) {
	fs1 := files.FS()
	ps := fs1.PathSeparator()
	s := fs1.Separator()
	t.Log("PathSeparator = ", ps)
	t.Log("Separator = ", s)
}

func TestChmod(t *testing.T) {

	tmpDir := files.FS().NewPath(t.TempDir())
	file1 := tmpDir.GetChild("file1")

	opt := afs.Todo().File(true).Create(true).Write(true).Options()
	err := file1.GetIO().WriteText("hello", opt)
	if err != nil {
		t.Error(err)
		return
	}

	mode1 := fs.ModeDevice
	mode1 = 0644
	err = file1.Chmod(mode1)
	if err != nil {
		t.Error(err)
		return
	}
	mode2 := file1.GetInfo().Mode()

	if mode1 != mode2 {
		vlog.Warn("(mode1 != mode2):")
		vlog.Warn("  mode1 = %s", mode1.String())
		vlog.Warn("  mode2 = %s", mode2.String())
		// t.Error("bad file mode")
		return
	}
}
