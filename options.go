package afs

import (
	"io/fs"
)

////////////////////////////////////////////////////////////////////////////////

// Options ...
type Options struct {

	// fill with fs.ModeXXX
	Permission fs.FileMode

	// fill with os.O_xxxx
	Flag int
}

////////////////////////////////////////////////////////////////////////////////

type PermissionMaker interface {
	Mode() fs.FileMode

	SetMode(owner, group, other int) PermissionMaker

	SetDir(isDir bool) PermissionMaker
}

type FlagMaker interface {
	Flags() int

	Create() FlagMaker
	Append() FlagMaker
	Truncate() FlagMaker
	Synchronous() FlagMaker
	Excl() FlagMaker

	ReadWrite() FlagMaker
	ReadOnly() FlagMaker
	WriteOnly() FlagMaker
}

////////////////////////////////////////////////////////////////////////////////
// EOF
