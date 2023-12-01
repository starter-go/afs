package afs

import (
	"io/fs"
	"os"
)

// Flags ...
type Flags struct {
	Mkdirs      bool
	Create      bool
	Append      bool
	Read        bool
	Write       bool
	File        bool
	Dir         bool
	FromBegin   bool
	ResetLength bool // set length = 0
}

////////////////////////////////////////////////////////////////////////////////

// Options ...
type Options struct {
	flags Flags

	// fill with fs.ModeXXX
	Permission fs.FileMode

	// fill with os.O_xxxx
	Flag int
}

// Create ...
func (inst *Options) Create(value bool) *Options {
	inst.flags.Create = value
	return inst
}

// Read ...
func (inst *Options) Read(value bool) *Options {
	inst.flags.Read = value
	return inst
}

// Write ...
func (inst *Options) Write(value bool) *Options {
	inst.flags.Write = value
	return inst
}

// File ...
func (inst *Options) File(value bool) *Options {
	inst.flags.File = value
	return inst
}

// Dir ...
func (inst *Options) Dir(value bool) *Options {
	inst.flags.Dir = value
	return inst
}

// Mkdirs ...
func (inst *Options) Mkdirs(value bool) *Options {
	inst.flags.Mkdirs = value
	return inst
}

// Append ...
func (inst *Options) Append(value bool) *Options {
	inst.flags.Append = value
	return inst
}

// FromBegin ...
func (inst *Options) FromBegin(value bool) *Options {
	inst.flags.FromBegin = value
	return inst
}

// ResetLength ...
func (inst *Options) ResetLength(value bool) *Options {
	inst.flags.ResetLength = value
	return inst
}

// Prepare ...
func (inst *Options) Prepare() *Options {

	flags := inst.flags
	f := 0

	if flags.Create {
		f = f | os.O_CREATE
	}

	if flags.Read {
		if flags.Write {
			f = f | os.O_RDWR
		} else {
			f = f | os.O_RDONLY
		}
	} else {
		if flags.Write {
			f = f | os.O_WRONLY
		} else {
			f = f | 0
		}
	}

	if flags.ResetLength {
		f = f | os.O_TRUNC
	}

	if flags.Append {
		f = f | os.O_APPEND
	}

	inst.Flag = f
	inst.Permission = fs.ModePerm
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// Todo ...
func Todo() *Options {
	return &Options{}
}

// ToMakeDir ...
func ToMakeDir() *Options {
	inst := Todo()
	inst.Permission = fs.ModePerm
	inst.Flag = 0
	return inst
}

// ToReadFile ...
func ToReadFile() *Options {
	inst := Todo()
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_RDONLY
	return inst
}

// ToWriteFile ...
func ToWriteFile() *Options {
	inst := Todo()
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_WRONLY | os.O_TRUNC
	return inst
}

// ToCreateFile ...
func ToCreateFile() *Options {
	inst := Todo()
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	return inst
}
