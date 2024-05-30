package action

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
)

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions

type Action struct {
	Name         string
	Author       string
	Description  string
	Inputs       Inputs
	InputsOrder  []string
	Outputs      Outputs
	OutputsOrder []string
}

func New(
	name string,
	author string,
	description string,
	inputs Inputs,
	inputsOrder []string,
	outputs Outputs,
	outputsOrder []string,
) *Action {
	return &Action{
		Name:         name,
		Author:       author,
		Description:  description,
		Inputs:       inputs,
		InputsOrder:  inputsOrder,
		Outputs:      outputs,
		OutputsOrder: outputsOrder,
	}
}

func codeBlock(code string) string {
	if strings.TrimSpace(code) == "" {
		return "` `"
	}
	return fmt.Sprintf("`%s`", code)
}

func (a *Action) GetInputsMatrix() [][]string {
	var matrix [][]string
	matrix = append(matrix, []string{"Name", "Description", "Required", "Default"})
	for _, key := range a.InputsOrder {
		input := a.Inputs[key]
		matrix = append(
			matrix,
			[]string{
				codeBlock(key),
				input.Description,
				codeBlock(strconv.FormatBool(input.Required)),
				codeBlock(input.Default),
			},
		)
	}
	return matrix
}

func (a *Action) GetOutputsMatrix() [][]string {
	var matrix [][]string
	matrix = append(matrix, []string{"Name", "Description"})
	for _, key := range a.OutputsOrder {
		output := a.Outputs[key]
		matrix = append(
			matrix,
			[]string{
				codeBlock(key),
				output.Description,
			},
		)
	}
	return matrix
}

type Input struct {
	Description string
	Required    bool
	Default     string
}

type ActionNodes struct {
	Inputs  yaml.Node `yaml:"inputs,omitempty"`
	Outputs yaml.Node `yaml:"outputs,omitempty"`
}

type Inputs = map[string]Input

type Output struct {
	Description string
}

type Outputs = map[string]Output
