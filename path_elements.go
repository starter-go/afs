package afs

import (
	"fmt"
	"os/user"
	"strings"
)

// PathElements 提供一组工具函数，用来处理路径中的各个项
type PathElements struct {
}

// Split 把路径字符串拆分成各个元素
func (inst *PathElements) Split(path string) []string {
	const (
		sep1 = "\\"
		sep2 = "/"
	)
	p2 := strings.ReplaceAll(path, sep1, sep2)
	return strings.Split(p2, sep2)
}

// Stringify 把一组路径元素拼合成字符串
func (inst *PathElements) Stringify(elist []string, prefix string, sep rune) string {
	b := &strings.Builder{}
	for i, el := range elist {
		if i == 0 {
			b.WriteString(prefix)
		} else {
			b.WriteRune(sep)
		}
		b.WriteString(el)
	}
	return b.String()
}

// Resolve 把路径中的特殊元素 (.|..|~|""|等) 处理，输出纯粹的路径
func (inst *PathElements) Resolve(elist []string) ([]string, error) {
	dst := make([]string, 0)
	for _, el := range elist {
		if el == "" {
			continue
		} else if el == "." {
			continue
		} else if el == ".." {
			count := len(dst)
			if count > 0 {
				dst = dst[0 : count-1]
			} else {
				return nil, fmt.Errorf("path contains too many '..' element(s)")
			}
		} else if el == "~" {
			home, err := inst.loadUserHome()
			if err != nil {
				return nil, err
			}
			dst = home
		} else {
			dst = append(dst, el)
		}
	}
	return dst, nil
}

func (inst *PathElements) loadUserHome() ([]string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	dir := u.HomeDir
	if strings.ContainsRune(dir, '~') {
		return nil, fmt.Errorf("bad user home dir path: %s", dir)
	}
	elist := inst.Split(dir)
	return inst.Resolve(elist)
}
