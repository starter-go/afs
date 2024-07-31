package platforms

import (
	"io/fs"
	"syscall"
	"time"

	"github.com/starter-go/base/lang"
)

// CreatedAt ...
func CreatedAt(info fs.FileInfo) time.Time {
	st, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		lt := lang.Time(0)
		return lt.Time()
	}
	ctim := st.Ctim
	sec, nsec := ctim.Unix()
	return time.Unix(sec, nsec)
}

// Roots ...
func Roots() []string {
	return []string{"/"}
}
