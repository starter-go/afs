package support

import (
	"io/fs"
	"os"

	"bitwormhole.com/starter/afs"
)

type defaultOptionsHandler struct{}

func (inst *defaultOptionsHandler) handle(path string, have, want *afs.Options) *afs.Options {

	opt := have
	if opt == nil {
		opt = want
	}
	if opt == nil {
		opt = &afs.Options{}
	}

	if opt.Write || opt.Create || opt.Mkdirs {
		if opt.Permission == 0 && opt.Flag == 0 {
			opt.Flag = os.O_RDWR | os.O_TRUNC
			opt.Permission = fs.ModePerm
		}
		if opt.Create {
			opt.Flag |= os.O_CREATE
		}
	}

	return opt
}
