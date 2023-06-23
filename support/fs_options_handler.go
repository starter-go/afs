package support

import (
	"io/fs"
	"os"

	"github.com/starter-go/afs"
)

type defaultOptionsHandler struct{}

func (inst *defaultOptionsHandler) handle(path string, have *afs.Options, want afs.WantOption) *afs.Options {

	if have != nil {
		return have
	}

	// make new opt
	have = &afs.Options{}

	have.Mkdirs = (want & afs.WantToMakeDir) != 0
	have.Create = (want & afs.WantToCreateFile) != 0
	have.Write = (want & afs.WantToWriteFile) != 0
	have.Read = (want & afs.WantToReadFile) != 0
	have.Permission = fs.ModePerm

	if have.Create {
		have.Flag |= os.O_CREATE
	}

	if have.Write {
		have.Flag |= os.O_WRONLY
		have.Flag |= os.O_TRUNC
	}

	if have.Read {
		have.Flag |= os.O_RDONLY
	}

	return have
}
