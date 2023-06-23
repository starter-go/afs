package afs

import "errors"

type WantOption int32

const (
	WantToMakeDir    = 0x0001
	WantToReadFile   = 0x0002
	WantToWriteFile  = 0x0004
	WantToCreateFile = 0x0008
)

// OptionsHandlerFunc 函数用于为I/O操作准备选项
type OptionsHandlerFunc func(path string, opt *Options, want WantOption) *Options

// FS 是表示文件系统的接口
type FS interface {
	NewPath(path string) Path

	ListRoots() []Path

	CreateTempFile(prefix, suffix string, dir Path) (Path, error)

	// PathSeparator return ';'(windows) | ':'(unix)
	PathSeparator() string

	// Separator return '/'(unix) | '\'(windows)
	Separator() string

	// 设置一个函数，用来处理默认的I/O选项
	SetDefaultOptionsHandler(fn OptionsHandlerFunc) error
}

// FileSystemFactory 是用来创建 FS 对象的工厂
type FileSystemFactory interface {
	Create() FS
}

////////////////////////////////////////////////////////////////////////////////

var theDefaultFS FS
var theDefaultFSFactory FileSystemFactory

// Default 获取默认的FS对象
func Default() FS {
	fs := theDefaultFS
	if fs != nil {
		return fs
	}
	factory := theDefaultFSFactory
	if factory == nil {
		panic("use SetDefaultFSFactory() to init Default()")
	}
	fs = factory.Create()
	theDefaultFS = fs
	return fs
}

// SetDefaultFSFactory 在调用 Default() 之前，必须设置默认的工厂
func SetDefaultFSFactory(factory FileSystemFactory) error {
	if theDefaultFS == nil {
		theDefaultFSFactory = factory
		return nil
	}
	return errors.New("SetDefaultFSFactory after init afs.Default")
}
