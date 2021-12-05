package stubs

import (
	"io/fs"
)

type StubDirEntry struct {
	IsDirectory bool
}

func (s StubDirEntry) Name() string {
	return "dir entry name"
}

func (s StubDirEntry) IsDir() bool {
	return s.IsDirectory
}

func (s StubDirEntry) Type() fs.FileMode {
	return fs.ModeDir
}

func (s StubDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}
