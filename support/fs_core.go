package support

import (
	"errors"
	"fmt"
	"strings"

	"github.com/starter-go/afs"
)

// CommonFileSystem 所有平台共享的核心结构
type CommonFileSystem interface {
	PathToElements(path string) []string
	ElementsToPath(elements []string, prefix string, sep string) string
	NormalizePathElements(elements []string) ([]string, error)
	PrepareOptions(path afs.Path, have *afs.Options, want afs.WantOption) *afs.Options
	SetDefaultOptionsHandler(fn afs.OptionsHandlerFunc) error
}

// PlatformFileSystem 各个平台独有的核心结构
type PlatformFileSystem interface {
	NormalizePath(path string) (string, error)
	PathSeparator() string
	Separator() string
	ListRoots() []string
	GetCommonFileSystem() CommonFileSystem
	GetFS() afs.FS
}

////////////////////////////////////////////////////////////////////////////////

// CommonFileSystemCore 所有平台共享的核心结构
type CommonFileSystemCore struct {
	context *myFSContext
}

func (inst *CommonFileSystemCore) _Impl() CommonFileSystem {
	return inst
}

// PathToElements 把路径拆分成文件名元素
func (inst *CommonFileSystemCore) PathToElements(path string) []string {
	path = strings.ReplaceAll(path, "\\", "/")
	return strings.Split(path, "/")
}

// ElementsToPath 把文件名元素拼接成路径
func (inst *CommonFileSystemCore) ElementsToPath(elements []string, prefix string, gap string) string {
	builder := &strings.Builder{}
	sep := prefix
	for _, el := range elements {
		builder.WriteString(sep)
		builder.WriteString(el)
		sep = gap
	}
	return builder.String()
}

// NormalizePathElements 标准化
func (inst *CommonFileSystemCore) NormalizePathElements(src []string) ([]string, error) {
	dst := make([]string, 0)
	for _, el := range src {
		el = strings.TrimSpace(el)
		if el == "" {
			// NOP
		} else if el == "." {
			// NOP
		} else if el == ".." {
			length := len(dst)
			if length > 0 {
				dst = dst[0 : length-1]
			} else {
				return nil, errors.New("too many '..'")
			}
		} else {
			dst = append(dst, el)
		}
	}
	return dst, nil
}

// PrepareOptions ...
func (inst *CommonFileSystemCore) PrepareOptions(p afs.Path, have *afs.Options, want afs.WantOption) *afs.Options {
	fn := inst.getDefaultOptionsHandler()
	path := p.GetPath()
	return fn(path, have, want)
}

// SetDefaultOptionsHandler ...
func (inst *CommonFileSystemCore) SetDefaultOptionsHandler(h afs.OptionsHandlerFunc) error {
	old := inst.context.optionsHandler
	if old != nil {
		return fmt.Errorf("an older handler has been configured")
	}
	if h == nil {
		return fmt.Errorf("param:handler is nil")
	}
	inst.context.optionsHandler = h
	return nil
}

func (inst *CommonFileSystemCore) getDefaultOptionsHandler() afs.OptionsHandlerFunc {
	h := inst.context.optionsHandler
	if h == nil {
		h = (&defaultOptionsHandler{}).handle
		inst.context.optionsHandler = h
	}
	return h
}
