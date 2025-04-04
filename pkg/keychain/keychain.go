package keychain

import (
	"os"
	"path/filepath"
)

//go:generate go run go.uber.org/mock/mockgen -source=keychain.go -destination=keychain_mock.go -package keychain
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
	BundleID string // 添加BundleID字段
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
	// 使用传入的bundleID
	if s.bundleID != "" {
		return filepath.Join(s.baseDir, s.bundleID, key)
	}

	// 兼容旧的解析方式，以防bundleID未设置
	parts := filepath.SplitList(key)
	if len(parts) < 2 {
		return filepath.Join(s.baseDir, "default", key)
	}
	return filepath.Join(s.baseDir, parts[0], parts[1])
}
