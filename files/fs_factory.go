package files

import (
	"runtime"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/support"
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
