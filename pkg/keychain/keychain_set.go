package keychain

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *simplestore) Set(key string, data []byte) error {
	filePath := s.getFilePath(key)

	// Ensure directory exists
	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, 0700); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write data to file
	if err := os.WriteFile(filePath, data, 0600); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}
