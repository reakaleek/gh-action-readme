package action

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAction_GetInputsMarkdown(t *testing.T) {
	// arrange
	a := Action{
		Name:        "Test Action",
		Description: "Test Action Description",
		Inputs: Inputs{
			"input1": {
				Description: "input1 description.",
				Required:    true,
				Default:     "",
			},
			"input2": {
				Description: "input2 description.",
				Required:    false,
				Default:     "default input2",
			},
		},
		InputsOrder: []string{"input1", "input2"},
	}
	// act
	matrix := a.GetInputsMatrix()
	// assert

	assert.Equal(t,
		[][]string{
			{"Name", "Description", "Required", "Default"},
			{"`input1`", "input1 description.", "`true`", "` `"},
			{"`input2`", "input2 description.", "`false`", "`default input2`"},
		},
		matrix,
	)

}
