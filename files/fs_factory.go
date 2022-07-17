package files

import (
	"runtime"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/support"
)

// FS ...
func FS() afs.FS {
	const win = "windows"
	o := runtime.GOOS
	if o == win {
		return support.GetWindowsFS()
	}
	return support.GetPosixFS()
}
