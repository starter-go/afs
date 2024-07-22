package afs

import (
	"io"
	"io/fs"
)

// Path 是表示绝对路径的接口
type Path interface {

	// 属性

	Exists() bool

	IsFile() bool

	IsDirectory() bool

	GetName() string

	GetPath() string

	GetURI() URI

	GetInfo() FileInfo

	String() string

	// 导航

	GetFS() FS

	GetParent() Path

	GetChild(name string) Path

	// 查询

	ListNames() []string

	ListPaths() []string

	ListChildren() []Path

	// 操作

	Mkdir(opt *Options) error

	Mkdirs(opt *Options) error

	MakeParents(opt *Options) error

	Chmod(m fs.FileMode) error

	Chown(uid, gid int) error

	Delete() error

	Create(opt *Options) error

	CreateWithData(data []byte, opt *Options) error

	CreateWithSource(src io.Reader, opt *Options) error

	MoveTo(dst Path, opt *Options) error

	CopyTo(dst Path, opt *Options) error

	// 读写

	GetIO() FileIO
}
