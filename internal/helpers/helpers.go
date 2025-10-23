package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindActionFile looks for action.yml or action.yaml in the current directory
func FindActionFile() (string, error) {
	// Check for action.yml first (more common)
	if _, err := os.Stat("action.yml"); err == nil {
		return "action.yml", nil
	}
	// Check for action.yaml
	if _, err := os.Stat("action.yaml"); err == nil {
		return "action.yaml", nil
	}
	return "", fmt.Errorf("neither action.yml nor action.yaml found in current directory")
}

// FindAllActionFiles recursively searches for all action.yml and action.yaml files
// starting from the given root directory
func FindAllActionFiles(root string) ([]string, error) {
	var actionFiles []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and common directories to ignore
		if d.IsDir() && d.Name() != "." && (d.Name()[0] == '.' || d.Name() == "node_modules" || d.Name() == "vendor") {
			return filepath.SkipDir
		}

		if !d.IsDir() && (d.Name() == "action.yml" || d.Name() == "action.yaml") {
			actionFiles = append(actionFiles, path)
		}

		return nil
	})

	return actionFiles, err
}
