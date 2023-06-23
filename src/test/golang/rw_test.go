package golang

import (
	"io"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
)

func TestTrunc(t *testing.T) {
	rows := []string{
		"abc",
		"def",
		"ijk",
		"opq",
		"xyz",
	}
	builder := strings.Builder{}
	versions := []string{}
	for _, row := range rows {
		builder.WriteString(row)
		builder.WriteString("\n")
		ver := builder.String()
		versions = append(versions, ver)
	}

	tmp := t.TempDir()
	name := "demo-for-test-x"
	file := files.FS().NewPath(tmp).GetChild(name)
	ops := &afs.Options{
		Create:     true,
		Mkdirs:     true,
		Permission: fs.ModePerm,
	}
	size := len(versions)

	// ops.Flag = os.O_TRUNC

	for i := size - 1; i > 0; i-- {
		text1 := versions[i]
		err := file.GetIO().WriteText(text1, ops)
		if err != nil {
			t.Error(err)
			break
		}
		text2, err := file.GetIO().ReadText(nil)
		if err != nil {
			t.Error(err)
			break
		}
		if text1 != text2 {
			t.Errorf("text1 != text2, want:[%v], have:[%v]", text1, text2)
			break
		}
	}
}

func TestSeekerRW(t *testing.T) {
	rows := []string{
		"abc",
		"def",
		"ijk",
		"opq",
		"xyz",
	}
	builder := strings.Builder{}
	versions := []string{}
	for _, row := range rows {
		builder.WriteString(row)
		builder.WriteString("\n")
		ver := builder.String()
		versions = append(versions, ver)
	}

	tmp := t.TempDir()
	name := "demo-for-test-"
	file := files.FS().NewPath(tmp).GetChild(name)
	ops := &afs.Options{
		Create:     true,
		Mkdirs:     true,
		Flag:       os.O_RDWR,
		Permission: fs.ModePerm,
	}

	// ops.Flag = os.O_TRUNC

	rw, err := file.GetIO().OpenSeekerRW(ops)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() { rw.Close() }()

	size := len(versions)
	for i := size - 1; i > 0; i-- {
		text1 := versions[i]
		size := len(text1)

		// write
		n64, err := rw.Seek(0, io.SeekEnd)
		if err != nil {
			t.Error(err)
			return
		}

		n32, err := rw.Write([]byte(text1))
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("write: cb=%v, pos=%v", n32, n64)

		// read
		n64, err = rw.Seek(int64(-size), io.SeekCurrent)
		if err != nil {
			t.Error(err)
			return
		}

		readBuffer := make([]byte, size)
		n32, err = rw.Read(readBuffer)
		if err != nil {
			t.Error(err)
			return
		}
		text2 := string(readBuffer)

		t.Logf("read : cb=%v, pos=%v", n32, n64)

		if text1 != text2 {
			t.Errorf("text1 != text2, want:[%v], have:[%v]", text1, text2)
		}
	}
}
