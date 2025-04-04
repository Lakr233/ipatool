package keychain

import (
	"os"
	"path/filepath"
)

type Keychain interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Remove(key string) error
}

type simplestore struct {
	baseDir  string
	bundleID string
}

type Args struct {
	BaseDir  string
	BundleID string
}

func New(args Args) Keychain {
	return &simplestore{
		baseDir:  args.BaseDir,
		bundleID: args.BundleID,
	}
}

func (s *simplestore) ensureDirectoryExists() error {
	dirPath := filepath.Join(s.baseDir, s.bundleID)
	return os.MkdirAll(dirPath, 0700)
}

func (s *simplestore) getFilePath(key string) string {
	if s.bundleID != "" {
		return filepath.Join(s.baseDir, s.bundleID, key)
	}

	parts := filepath.SplitList(key)
	if len(parts) < 2 {
		return filepath.Join(s.baseDir, "default", key)
	}
	return filepath.Join(s.baseDir, parts[0], parts[1])
}
