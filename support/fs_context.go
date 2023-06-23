package support

import "github.com/starter-go/afs"

type myFSContext struct {
	common         CommonFileSystem
	platform       PlatformFileSystem
	shell          afs.FS
	optionsHandler afs.OptionsHandlerFunc
}
