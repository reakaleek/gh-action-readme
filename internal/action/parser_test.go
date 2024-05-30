package action_test

import (
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	// arrange
	actionReader := action.NewParser()

	// act
	a, err := actionReader.Parse(filepath.Join("..", "testdata", "1-action.yml"))

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "Test Action", a.Name)
	assert.Equal(t, "", a.Author)
	assert.Equal(t, "Test Action description.\n", a.Description)
	assert.Contains(t, a.Inputs, "input1")
	assert.Equal(t, true, a.Inputs["input1"].Required)
	assert.Equal(t, "input1 description.", a.Inputs["input1"].Description)
	assert.Contains(t, a.Inputs, "input2")
	assert.Equal(t, false, a.Inputs["input2"].Required)
	assert.Equal(t, "input2 description.", a.Inputs["input2"].Description)
	assert.Contains(t, a.Outputs, "output1")
	assert.Equal(t, "output1 description.", a.Outputs["output1"].Description)
}
