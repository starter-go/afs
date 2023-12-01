package support

import (
	"github.com/starter-go/afs"
)

type defaultOptionsHandler struct{}

func (inst *defaultOptionsHandler) handle(path string, have *afs.Options, want afs.WantOption) *afs.Options {
	if have != nil {
		return have
	}
	switch want {
	case afs.WantToCreateFile:
		return afs.ToCreateFile()

	case afs.WantToMakeDir:
		return afs.ToMakeDir()

	case afs.WantToReadFile:
		return afs.ToReadFile()

	case afs.WantToWriteFile:
		return afs.ToWriteFile()
	}
	return afs.Todo()
}
