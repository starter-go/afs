package files

import (
	"runtime"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/support"
)

var theDefaultFS afs.FS

// FS ...
func FS() afs.FS {
	fs := theDefaultFS
	if fs != nil {
		return fs
	}

	const win = "windows"
	sys := runtime.GOOS

	if sys == win {
		fs = support.GetWindowsFS()
	} else {
		fs = support.GetPosixFS()
	}

	theDefaultFS = fs
	return fs
}
