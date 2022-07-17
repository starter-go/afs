package support

// type myCommonFS struct {
// 	context *myFSContext
// }

// func (inst *myCommonFS) _Impl() CommonFileSystem {
// 	return inst
// }

// func (inst *myCommonFS) NewPath(path string) afs.Path {
// 	p2, _ := inst.core.NormalizePath(path)
// 	return &myCommonPath{fs: inst, path: p2}
// }

// func (inst *myCommonFS) ListRoots() []afs.Path {
// 	src := inst.core.ListRoots()
// 	dst := make([]afs.Path, 0)
// 	for _, item := range src {
// 		path := inst.NewPath(item)
// 		dst = append(dst, path)
// 	}
// 	return dst
// }

// func (inst *myCommonFS) CreateTempFile(prefix, suffix string, folder afs.Path) (afs.Path, error) {
// 	dir := ""
// 	pattern := prefix + "*" + suffix
// 	if folder != nil {
// 		dir = folder.String()
// 	}
// 	file, err := os.CreateTemp(dir, pattern)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	path := file.Name()
// 	return inst.NewPath(path), nil
// }

// // PathSeparator return ';'(windows) | ':'(unix)
// func (inst *myCommonFS) PathSeparator() string {
// 	return inst.core.PathSeparator()
// }

// // Separator return '/'(unix) | '\'(windows)
// func (inst *myCommonFS) Separator() string {
// 	return inst.core.Separator()
// }
