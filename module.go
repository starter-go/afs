package afs

import (
	"strconv"
	"strings"
)

const (
	theModuleName     = "github.com/starter-go/afs"
	theModuleVersion  = "v0.9.13"
	theModuleRevision = 21
)

func NewModule() string {

	b := strings.Builder{}
	b.WriteString("[module")

	b.WriteString(" name:")
	b.WriteString(theModuleName)

	b.WriteString(" version:")
	b.WriteString(theModuleVersion)

	b.WriteString(" revision:")
	b.WriteString(strconv.Itoa(theModuleRevision))

	b.WriteString("]")
	return b.String()
}
