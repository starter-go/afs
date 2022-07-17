package support

import "bitwormhole.com/starter/afs"

type myFSContext struct {
	common   CommonFileSystem
	platform PlatformFileSystem
	shell    afs.FS
}
