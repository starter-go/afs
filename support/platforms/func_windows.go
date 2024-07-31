package platforms

import (
	"io/fs"
	"syscall"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
	"golang.org/x/sys/windows"
)

// CreatedAt ...
func CreatedAt(info fs.FileInfo) time.Time {
	lt0 := lang.Time(0)
	attr, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return lt0.Time()
	}
	ctim := attr.CreationTime
	ns := ctim.Nanoseconds()
	dur := time.Nanosecond * time.Duration(ns)
	return lt0.Add(dur).Time()
}

// Roots ...
func Roots() []string {
	const (
		driveA rune = 'A'
		driveZ rune = 'Z'
	)
	list := make([]string, 0)
	bits, err := windows.GetLogicalDrives()
	if err != nil {
		vlog.Warn(err.Error())
		return list
	}
	for drive := driveA; drive <= driveZ; drive++ {
		idx := int(drive - driveA)
		has := (bits >> idx) & 0x0001
		path := string(drive) + ":\\"
		if has != 0 {
			list = append(list, path)
		}
	}
	return list
}
