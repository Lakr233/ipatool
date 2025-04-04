package keychain

import (
	"fmt"
	"os"
)

func (s *simplestore) Get(key string) ([]byte, error) {
	filePath := s.getFilePath(key)

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("key not found: %s", key)
		}
		return nil, fmt.Errorf("failed to read key: %w", err)
	}

	return data, nil
}
