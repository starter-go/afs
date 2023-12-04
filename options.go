package afs

import (
	"io/fs"
	"os"
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

// ToMakeDir ...
func ToMakeDir() *Options {
	opt := new(Options)
	opt.Permission = fs.ModePerm
	opt.Flag = 0
	return opt
}

// ToReadFile ...
func ToReadFile() *Options {
	opt := new(Options)
	opt.Permission = fs.ModePerm
	opt.Flag = os.O_RDONLY
	return opt
}

// ToWriteFile ...
func ToWriteFile() *Options {
	opt := new(Options)
	opt.Permission = fs.ModePerm
	opt.Flag = os.O_WRONLY | os.O_TRUNC
	return opt
}

// ToCreateFile ...
func ToCreateFile() *Options {
	opt := new(Options)
	opt.Permission = fs.ModePerm
	opt.Flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	return opt
}

////////////////////////////////////////////////////////////////////////////////

// OptionsBuilder 用于创建复合的 Options
type OptionsBuilder struct {
	mkdirs      bool
	create      bool
	append      bool
	read        bool
	write       bool
	file        bool
	dir         bool
	fromBegin   bool
	resetLength bool // set length = 0
}

// Create ...
func (inst *OptionsBuilder) Create(value bool) *OptionsBuilder {
	inst.create = value
	return inst
}

// Read ...
func (inst *OptionsBuilder) Read(value bool) *OptionsBuilder {
	inst.read = value
	return inst
}

// Write ...
func (inst *OptionsBuilder) Write(value bool) *OptionsBuilder {
	inst.write = value
	return inst
}

// File ...
func (inst *OptionsBuilder) File(value bool) *OptionsBuilder {
	inst.file = value
	return inst
}

// Dir ...
func (inst *OptionsBuilder) Dir(value bool) *OptionsBuilder {
	inst.dir = value
	return inst
}

// Mkdirs ...
func (inst *OptionsBuilder) Mkdirs(value bool) *OptionsBuilder {
	inst.mkdirs = value
	return inst
}

// Append ...
func (inst *OptionsBuilder) Append(value bool) *OptionsBuilder {
	inst.append = value
	return inst
}

// FromBegin ...
func (inst *OptionsBuilder) FromBegin(value bool) *OptionsBuilder {
	inst.fromBegin = value
	return inst
}

// ResetLength ...
func (inst *OptionsBuilder) ResetLength(value bool) *OptionsBuilder {
	inst.resetLength = value
	return inst
}

// Options 创建 Options
func (inst *OptionsBuilder) Options() *Options {

	f := 0

	if inst.create {
		f = f | os.O_CREATE
	}

	if inst.read {
		if inst.write {
			f = f | os.O_RDWR
		} else {
			f = f | os.O_RDONLY
		}
	} else {
		if inst.write {
			f = f | os.O_WRONLY
		} else {
			f = f | 0
		}
	}

	if inst.resetLength {
		f = f | os.O_TRUNC
	}

	if inst.append {
		f = f | os.O_APPEND
	}

	opt := new(Options)
	opt.Flag = f
	opt.Permission = fs.ModePerm
	return opt
}

////////////////////////////////////////////////////////////////////////////////

// Todo 新建一个 OptionsBuilder 对象
func Todo() *OptionsBuilder {
	return new(OptionsBuilder)
}
