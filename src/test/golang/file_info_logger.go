package golang

import (
	"strconv"
	"strings"
	"testing"

	"github.com/starter-go/afs"
)

type fileInfoLogger struct {
	LogFileContent bool
	t              *testing.T
}

func (inst *fileInfoLogger) log(p afs.Path) {

	if p == nil {
		return
	}

	builder := strings.Builder{}

	// path
	builder.WriteString(" path:[")
	builder.WriteString(p.GetPath())
	builder.WriteString("]")

	// mode
	builder.WriteString(" mode:")
	builder.WriteString(p.GetInfo().Mode().String())

	// size
	size := p.GetInfo().Length()
	builder.WriteString(" size:")
	builder.WriteString(strconv.FormatInt(size, 10))

	inst.t.Log(builder.String())
}
