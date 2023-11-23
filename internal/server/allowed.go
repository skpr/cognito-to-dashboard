package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

// IsAllowed loads the allowed list file and checks if the dashboard exists.
func IsAllowed(filePath, dashboardName string) (bool, error) {
	var allowed []string

	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to open allowed list file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return false, fmt.Errorf("failed to read allowed list file: %w", err)
	}

	err = json.Unmarshal(data, &allowed)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal allowed list file: %w", err)
	}

	if slices.Contains(allowed, dashboardName) {
		return true, nil
	}

	return false, nil
}
