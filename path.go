package afs

import "io"

// Path 是表示绝对路径的接口
type Path interface {

	// 属性

	Exists() bool

	IsFile() bool

	IsDirectory() bool

	GetName() string

	GetPath() string

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

	Mkdir(op Options) error

	Mkdirs(op Options) error

	Delete() error

	Create(op Options) error

	CreateWithData(data []byte, op Options) error

	CreateWithSource(src io.Reader, op Options) error

	// 读写

	GetIO() FileIO
}
