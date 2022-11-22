package golang

import (
	"os"
	"strings"
	"testing"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
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
	name := "demo-for-test-"
	file := files.FS().NewPath(tmp).GetChild(name)
	ops := &afs.Options{Create: true}
	size := len(versions)

	ops.Flag = os.O_TRUNC

	for i := size - 1; i > 0; i-- {
		text1 := versions[i]
		file.GetIO().WriteText(text1, ops)
		text2, err := file.GetIO().ReadText(nil)
		if err != nil {
			t.Error(err)
		}
		if text1 != text2 {
			t.Errorf("text1 != text2, want:[%v], have:[%v]", text1, text2)
		}
	}
}
