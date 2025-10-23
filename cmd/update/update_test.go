package update_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reakaleek/gh-action-readme/cmd/update"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

// setupTestDir creates a temporary test directory with action.yml
func setupTestDir(t *testing.T) string {
	tmpDir := t.TempDir()
	
	actionYML := `name: Test Action
description: A test action for integration tests
inputs:
  test-input:
    description: 'Test input'
    required: true
    default: 'test'
outputs:
  test-output:
    description: 'Test output'
`
	
	err := os.WriteFile(filepath.Join(tmpDir, "action.yml"), []byte(actionYML), 0644)
	require.NoError(t, err)
	
	return tmpDir
}

// TestUpdateCommand tests the update command
func TestUpdateCommand(t *testing.T) {
	tmpDir := setupTestDir(t)
	readmePath := filepath.Join(tmpDir, "README.md")
	
	// Create initial README with tags
	initialReadme := `# Test
<!--name--><!--/name-->
<!--description-->
<!--inputs-->
<!--outputs-->
`
	err := os.WriteFile(readmePath, []byte(initialReadme), 0644)
	require.NoError(t, err)
	
	// Change to tmpDir so action.yml is found
	originalWd, _ := os.Getwd()
	defer func() { _ = os.Chdir(originalWd) }()
	err = os.Chdir(tmpDir)
	require.NoError(t, err)
	
	// Run update command
	app := &cli.App{
		Commands: []*cli.Command{update.NewCommand()},
	}
	
	err = app.Run([]string{"app", "update", "--readme", readmePath})
	assert.NoError(t, err)
	
	// Verify README was updated
	content, err := os.ReadFile(readmePath)
	require.NoError(t, err)
	
	contentStr := string(content)
	assert.Contains(t, contentStr, "Test Action")
	assert.Contains(t, contentStr, "A test action")
	assert.Contains(t, contentStr, "test-input")
	assert.Contains(t, contentStr, "test-output")
}

// TestUpdateCommandBackwardsCompatible tests that --action flag is accepted
func TestUpdateCommandBackwardsCompatible(t *testing.T) {
	tmpDir := setupTestDir(t)
	readmePath := filepath.Join(tmpDir, "README.md")
	
	initialReadme := `# Test
<!--name--><!--/name-->
<!--description-->
<!--inputs-->
<!--outputs-->
`
	err := os.WriteFile(readmePath, []byte(initialReadme), 0644)
	require.NoError(t, err)
	
	// Change to tmpDir
	originalWd, _ := os.Getwd()
	defer func() { _ = os.Chdir(originalWd) }()
	err = os.Chdir(tmpDir)
	require.NoError(t, err)
	
	app := &cli.App{
		Commands: []*cli.Command{update.NewCommand()},
	}
	
	// Old way with --action flag (should be ignored but accepted)
	err = app.Run([]string{"app", "update", "--action", "action.yml", "--readme", readmePath})
	assert.NoError(t, err)
	
	// Verify it still works
	content, err := os.ReadFile(readmePath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Test Action")
}

// TestRecursiveUpdate tests recursive mode
func TestRecursiveUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	
	// Create multiple action directories
	dirs := []string{"action1", "action2"}
	for _, dir := range dirs {
		actionDir := filepath.Join(tmpDir, dir)
		err := os.MkdirAll(actionDir, 0755)
		require.NoError(t, err)
		
		actionYML := `name: ` + dir + `
description: Test action ` + dir + `
inputs:
  input1:
    description: 'Input 1'
`
		err = os.WriteFile(filepath.Join(actionDir, "action.yml"), []byte(actionYML), 0644)
		require.NoError(t, err)
		
		readme := `<!--name--><!--/name-->
<!--description-->
<!--inputs-->
`
		err = os.WriteFile(filepath.Join(actionDir, "README.md"), []byte(readme), 0644)
		require.NoError(t, err)
	}
	
	// Change to tmpDir to run recursive update
	originalWd, _ := os.Getwd()
	defer func() { _ = os.Chdir(originalWd) }()
	_ = os.Chdir(tmpDir)
	
	app := &cli.App{
		Commands: []*cli.Command{update.NewCommand()},
	}
	
	err := app.Run([]string{"app", "update", "--recursive"})
	assert.NoError(t, err)
	
	// Verify both READMEs were updated
	for _, dir := range dirs {
		readmePath := filepath.Join(tmpDir, dir, "README.md")
		content, err := os.ReadFile(readmePath)
		require.NoError(t, err)
		assert.Contains(t, string(content), dir)
	}
}

// TestActionYamlDetection tests that both .yml and .yaml are detected
func TestActionYamlDetection(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{"yml extension", "action.yml"},
		{"yaml extension", "action.yaml"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			
			actionContent := `name: Test
description: Test action
`
			err := os.WriteFile(filepath.Join(tmpDir, tt.filename), []byte(actionContent), 0644)
			require.NoError(t, err)
			
			readme := `<!--name--><!--/name-->
<!--description-->
`
			readmePath := filepath.Join(tmpDir, "README.md")
			err = os.WriteFile(readmePath, []byte(readme), 0644)
			require.NoError(t, err)
			
		// Change to tmpDir
		originalWd, _ := os.Getwd()
		defer func() { _ = os.Chdir(originalWd) }()
		_ = os.Chdir(tmpDir)
			
			app := &cli.App{
				Commands: []*cli.Command{update.NewCommand()},
			}
			
			err = app.Run([]string{"app", "update", "--readme", readmePath})
			assert.NoError(t, err)
			
			content, err := os.ReadFile(readmePath)
			require.NoError(t, err)
			assert.Contains(t, string(content), "Test action")
		})
	}
}
