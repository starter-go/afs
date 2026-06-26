package afs

import (
	"io/fs"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

type OptionsMaker struct {
	op Options
}

func (inst *OptionsMaker) Reset() *OptionsMaker {
	inst.op.Flag = 0
	inst.op.Permission = 0
	return inst
}

func (inst *OptionsMaker) Options() Options {
	return inst.op
}

func (inst *OptionsMaker) SetPermissions() PermissionMaker {
	inst.op.Permission = 0
	return inst
}

func (inst *OptionsMaker) SetFlags() FlagMaker {
	inst.op.Flag = 0
	return inst
}

func (inst *OptionsMaker) Mode() fs.FileMode {
	return inst.op.Permission
}

func (inst *OptionsMaker) Flags() int {
	return inst.op.Flag
}

func (inst *OptionsMaker) SetMode(owner, group, other int) PermissionMaker {

	// mask: 1 - 11 111 111
	// const mask777 uint32 = 0x1ff

	const mask7 uint32 = 0x07
	const all uint32 = 0xffffffff

	tmp := uint32(all)
	tmp <<= 9
	older := uint32(inst.op.Permission) & tmp

	tmp = (uint32(owner) & mask7)
	tmp = (uint32(group) & mask7) | (tmp << 3)
	tmp = (uint32(other) & mask7) | (tmp << 3)

	tmp = tmp | older

	inst.op.Permission = fs.FileMode(tmp)
	return inst
}

func (inst *OptionsMaker) SetDir(isDir bool) PermissionMaker {
	inst.op.Permission = 0
	//todo ...
	return inst
}

func (inst *OptionsMaker) Create() FlagMaker {
	inst.op.Flag |= os.O_CREATE
	return inst
}

func (inst *OptionsMaker) Append() FlagMaker {
	inst.op.Flag |= os.O_APPEND
	return inst
}

func (inst *OptionsMaker) Truncate() FlagMaker {
	inst.op.Flag |= os.O_TRUNC
	return inst
}

func (inst *OptionsMaker) Synchronous() FlagMaker {
	inst.op.Flag |= os.O_SYNC
	return inst
}

func (inst *OptionsMaker) Excl() FlagMaker {
	inst.op.Flag |= os.O_EXCL
	return inst
}

func (inst *OptionsMaker) ReadWrite() FlagMaker {
	inst.op.Flag |= os.O_RDWR
	return inst
}

func (inst *OptionsMaker) ReadOnly() FlagMaker {
	inst.op.Flag |= os.O_RDONLY
	return inst
}

func (inst *OptionsMaker) WriteOnly() FlagMaker {
	inst.op.Flag |= os.O_WRONLY
	return inst
}

func (inst *OptionsMaker) _impl() innerOptionsMaker {

	// n := os.O_CREATE
	// n |= os.O_APPEND
	// n |= os.O_TRUNC
	// n |= os.O_SYNC
	// n |= os.O_EXCL

	// n |= os.O_RDONLY
	// n |= os.O_RDWR
	// n |= os.O_WRONLY

	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerOptionsMaker interface {
	SetPermissions() PermissionMaker

	SetFlags() FlagMaker

	Options() Options
}

////////////////////////////////////////////////////////////////////////////////
// EOF
