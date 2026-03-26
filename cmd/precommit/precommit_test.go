package precommit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

// runPrecommit is a helper that invokes the pre-commit command with the given
// file arguments and returns the error (or nil on success).
func runPrecommit(t *testing.T, files []string) error {
	t.Helper()
	app := &cli.App{Commands: []*cli.Command{NewCommand()}}
	args := append([]string{"app", "pre-commit"}, files...)
	return app.Run(args)
}

// TestPrecommit_MissingReadme verifies that the pre-commit command does not
// panic when an action.yml exists but README.md does not. It should generate
// the README and write it to disk.
func TestPrecommit_MissingReadme(t *testing.T) {
	tmpDir := t.TempDir()
	actionPath := filepath.Join(tmpDir, "action.yml")
	readmePath := filepath.Join(tmpDir, "README.md")

	actionContent := `name: Test Action
description: A test action`
	require.NoError(t, os.WriteFile(actionPath, []byte(actionContent), 0644))

	// README.md must NOT exist before the call.
	_, err := os.Stat(readmePath)
	require.True(t, os.IsNotExist(err), "README.md should not exist before running pre-commit")

	// Run the pre-commit command — must not panic or return an error.
	err = runPrecommit(t, []string{actionPath})
	assert.NoError(t, err)

	// README.md should now have been created.
	content, err := os.ReadFile(readmePath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "Test Action", "generated README should contain the action name")
}

// TestPrecommit_ExistingUpToDateReadme verifies that an already-up-to-date
// README.md is not rewritten.
func TestPrecommit_ExistingUpToDateReadme(t *testing.T) {
	tmpDir := t.TempDir()
	actionPath := filepath.Join(tmpDir, "action.yml")
	readmePath := filepath.Join(tmpDir, "README.md")

	actionContent := `name: My Action
description: My Description`
	require.NoError(t, os.WriteFile(actionPath, []byte(actionContent), 0644))

	// First run: generate the README.
	require.NoError(t, runPrecommit(t, []string{actionPath}))
	first, err := os.ReadFile(readmePath)
	require.NoError(t, err)

	// Second run: the README should remain unchanged.
	require.NoError(t, runPrecommit(t, []string{actionPath}))
	second, err := os.ReadFile(readmePath)
	require.NoError(t, err)

	assert.Equal(t, string(first), string(second), "README.md should not change when already up-to-date")
}
