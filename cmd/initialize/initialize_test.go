package initialize_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/reakaleek/gh-action-readme/cmd/initialize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

// TestInitCommand tests the init command
func TestInitCommand(t *testing.T) {
	tmpDir := t.TempDir()
	readmePath := filepath.Join(tmpDir, "README.md")
	
	app := &cli.App{
		Commands: []*cli.Command{initialize.NewCommand()},
	}
	
	err := app.Run([]string{"app", "init", "--readme", readmePath})
	assert.NoError(t, err)
	
	// Verify README was created
	_, err = os.Stat(readmePath)
	assert.NoError(t, err)
	
	content, err := os.ReadFile(readmePath)
	require.NoError(t, err)
	
	// Verify it has the expected tags
	contentStr := string(content)
	assert.Contains(t, contentStr, "<!--name-->")
	assert.Contains(t, contentStr, "<!--description-->")
	assert.Contains(t, contentStr, "<!--inputs-->")
	assert.Contains(t, contentStr, "<!--outputs-->")
}
