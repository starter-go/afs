package afs

import "io/fs"

// Options ...
type Options struct {

	// fill with fs.ModeXXX
	Permission fs.FileMode

	// fill with os.O_xxxx
	Flag int

	Mkdirs    bool
	Create    bool
	Read      bool
	Write     bool
	File      bool
	Directory bool
}
