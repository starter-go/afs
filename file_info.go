package afs

import (
	"io/fs"
	"time"
)

// FileInfo ...
type FileInfo interface {
	Path() Path

	Length() int64

	CreatedAt() time.Time

	UpdatedAt() time.Time

	Mode() fs.FileMode

	Exists() bool

	IsFile() bool

	IsDirectory() bool
}
