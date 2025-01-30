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
	f := DefaultOptionsBuilderFactory()
	b := f.NewBuilder()
	b.Create()
	return b.Options()
}

// ToReadFile ...
func ToReadFile() *Options {
	f := DefaultOptionsBuilderFactory()
	b := f.NewBuilder()
	b.ReadOnly()
	return b.Options()
}

// ToWriteFile ...
func ToWriteFile() *Options {
	f := DefaultOptionsBuilderFactory()
	b := f.NewBuilder()
	b.WriteOnly()
	return b.Options()
}

// ToCreateFile ...
func ToCreateFile() *Options {
	f := DefaultOptionsBuilderFactory()
	b := f.NewBuilder()
	b.Create()
	return b.Options()
}

////////////////////////////////////////////////////////////////////////////////

// OptionsBuilder 用于创建复合的 Options
// [已废弃]：用 OptionsBuilderV2 代替
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

// OptionsBuilderV2 是第二版的 Options-Builder 接口
type OptionsBuilderV2 interface {
	SetMode(mode fs.FileMode) OptionsBuilderV2
	SetFlag(flag int) OptionsBuilderV2

	//	CreateDir() OptionsBuilderV2
	//	ReadDir() OptionsBuilderV2

	Create() OptionsBuilderV2
	ReadOnly() OptionsBuilderV2
	WriteOnly() OptionsBuilderV2
	ReadWrite() OptionsBuilderV2
	Append() OptionsBuilderV2
	Truncate() OptionsBuilderV2
	Excl() OptionsBuilderV2
	Synchronous() OptionsBuilderV2

	Options() *Options
}

////////////////////////////////////////////////////////////////////////////////

// OptionsBuilderFactory 用于创建 OptionsBuilderV2
type OptionsBuilderFactory interface {
	NewBuilder() OptionsBuilderV2
}

var theDefaultOptionsBuilderFactory OptionsBuilderFactory

// DefaultOptionsBuilderFactory 用于获取默认的 OptionsBuilderFactory
func DefaultOptionsBuilderFactory() OptionsBuilderFactory {
	f := theDefaultOptionsBuilderFactory
	if f == nil {
		f = new(myDefaultOptionsBuilderFactory)
		theDefaultOptionsBuilderFactory = f
	}
	return f
}

// SetDefaultOptionsBuilderFactory 用于设置默认的 OptionsBuilderFactory
func SetDefaultOptionsBuilderFactory(f OptionsBuilderFactory) {
	if f == nil {
		return
	}
	theDefaultOptionsBuilderFactory = f
}

////////////////////////////////////////////////////////////////////////////////

// Todo 新建一个 OptionsBuilder 对象
func Todo() *OptionsBuilder {
	return new(OptionsBuilder)
}

////////////////////////////////////////////////////////////////////////////////

type myDefaultOptionsBuilderFactory struct {
}

func (inst *myDefaultOptionsBuilderFactory) _impl() OptionsBuilderFactory {
	return inst
}

func (inst *myDefaultOptionsBuilderFactory) NewBuilder() OptionsBuilderV2 {
	b := &myDefaultOptionsBuilder{}
	b.opt.Permission = fs.ModePerm
	return b
}

////////////////////////////////////////////////////////////////////////////////

type myDefaultOptionsBuilder struct {
	opt Options
}

func (inst *myDefaultOptionsBuilder) _impl() OptionsBuilderV2 {
	return inst
}

func (inst *myDefaultOptionsBuilder) SetMode(mode fs.FileMode) OptionsBuilderV2 {
	inst.opt.Permission = mode
	return inst
}

func (inst *myDefaultOptionsBuilder) SetFlag(flag int) OptionsBuilderV2 {
	inst.opt.Flag = flag
	return inst
}

func (inst *myDefaultOptionsBuilder) Create() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_CREATE
	return inst
}

func (inst *myDefaultOptionsBuilder) ReadOnly() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_RDONLY
	return inst
}

func (inst *myDefaultOptionsBuilder) WriteOnly() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_WRONLY
	return inst
}

func (inst *myDefaultOptionsBuilder) ReadWrite() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_RDWR
	return inst
}

func (inst *myDefaultOptionsBuilder) Append() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_APPEND
	return inst
}

func (inst *myDefaultOptionsBuilder) Truncate() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_TRUNC
	return inst
}

func (inst *myDefaultOptionsBuilder) Excl() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_EXCL
	return inst
}

func (inst *myDefaultOptionsBuilder) Synchronous() OptionsBuilderV2 {
	inst.opt.Flag |= os.O_SYNC
	return inst
}

func (inst *myDefaultOptionsBuilder) Options() *Options {
	dst := new(Options)
	*dst = inst.opt
	return dst
}

////////////////////////////////////////////////////////////////////////////////
