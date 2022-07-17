package afs

import "errors"

// FS 是表示文件系统的接口
type FS interface {
	NewPath(path string) Path

	ListRoots() []Path

	CreateTempFile(prefix, suffix string, dir Path) (Path, error)

	// PathSeparator return ';'(windows) | ':'(unix)
	PathSeparator() string

	// Separator return '/'(unix) | '\'(windows)
	Separator() string
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
