package afs

import (
	"io/fs"
	"os"
)

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

func (inst *Options) ToMakeDir() *Options {
	inst.Permission = fs.ModePerm
	inst.Flag = 0

	inst.Mkdirs = true

	return inst
}

func (inst *Options) ToReadFile() *Options {
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_RDONLY

	inst.Read = true
	inst.File = true

	return inst
}

func (inst *Options) ToWriteFile() *Options {
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_WRONLY | os.O_TRUNC

	inst.Write = true
	inst.File = true

	return inst
}

func (inst *Options) ToCreateFile() *Options {
	inst.Permission = fs.ModePerm
	inst.Flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE

	inst.Create = true
	inst.Write = true
	inst.File = true

	return inst
}
