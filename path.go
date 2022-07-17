package afs

// Path 是表示绝对路径的接口
type Path interface {
	GetFS() FS

	GetParent() Path

	GetChild(name string) Path

	String() string

	GetName() string

	GetPath() string

	GetInfo() FileInfo

	GetIO() FileIO

	Mkdir(op Options) error

	Mkdirs(op Options) error

	Delete() error

	ListNames() []string

	ListPaths() []string

	ListChildren() []Path
}
