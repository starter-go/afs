package afs

// WantOption 。。。
type WantOption int32

// 定义 WantOption
const (
	WantToMakeDir    = 0x0001
	WantToReadFile   = 0x0002
	WantToWriteFile  = 0x0004
	WantToCreateFile = 0x0008
)

// OptionsHandlerFunc 函数用于为I/O操作准备选项
type OptionsHandlerFunc func(path string, opt *Options, want WantOption) *Options
