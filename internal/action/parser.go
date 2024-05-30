package action

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Parser struct {
}

func (a *Parser) Parse(path string) (Action, error) {
	var action Action
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return action, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, &action)
	if err != nil {
		return action, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}
	err = attachInputsAndOutputsOrder(&action, yamlFile)
	if err != nil {
		return action, fmt.Errorf("failed to attach inputs and outputs order: %w", err)
	}
	return action, err
}

func attachInputsAndOutputsOrder(action *Action, yamlFile []byte) error {
	var actionNodes ActionNodes
	err := yaml.Unmarshal(yamlFile, &actionNodes)
	if err != nil {
		return err
	}
	err = attachInputsOrder(action, &actionNodes)
	if err != nil {
		return err
	}
	err = attachOutputsOrder(action, &actionNodes)
	if err != nil {
		return err
	}
	return nil
}

func attachInputsOrder(action *Action, actionNodes *ActionNodes) error {
	var err error
	if actionNodes.Inputs.Content != nil {
		action.InputsOrder, err = getOrderedKeys(&actionNodes.Inputs)
		if err != nil {
			return fmt.Errorf("failed to parse inputs order: %w", err)
		}
	}
	return nil
}

func attachOutputsOrder(action *Action, actionNodes *ActionNodes) error {
	var err error
	if actionNodes.Outputs.Content != nil {
		action.OutputsOrder, err = getOrderedKeys(&actionNodes.Outputs)
		if err != nil {
			return fmt.Errorf("failed to parse outputs order: %w", err)
		}
	}
	return nil
}

func getOrderedKeys(node *yaml.Node) ([]string, error) {
	if node.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("expected mapping node, got %v", node.Kind)
	}
	var keys []string
	for i := 0; i < len(node.Content); i++ {
		value := node.Content[i].Value
		if value != "" {
			keys = append(keys, value)
		}
	}
	return keys, nil
}

func NewParser() *Parser {
	return &Parser{}
}
