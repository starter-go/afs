package golang

import (
	"testing"

	"bitwormhole.com/starter/afs/files"
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
