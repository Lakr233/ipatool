package keychain

import (
	"fmt"
	"os"
)

func (s *simplestore) Remove(key string) error {
	filePath := s.getFilePath(key)

	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Already removed
		}
		return fmt.Errorf("failed to remove key: %w", err)
	}

	return nil
}
