package update

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/reakaleek/gh-action-readme/internal/helpers"
	"github.com/reakaleek/gh-action-readme/internal/markdown"
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	var readmePath string
	var recursive bool
	var _ string // unused actionPath for backwards compatibility
	return &cli.Command{
		Name:  "update",
		Usage: "Update README.md",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:   "action",
				Hidden: true,
				Usage:  "Deprecated: action.yml/action.yaml are now auto-detected",
			},
			&cli.StringFlag{
				Name:        "readme",
				Value:       "README.md",
				Destination: &readmePath,
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Value:       false,
				Destination: &recursive,
				Usage:       "Search recursively for all action.yml/action.yaml files",
			},
		},
		Action: func(ctx *cli.Context) error {
			return updateRun(readmePath, recursive)
		},
	}
}

func updateRun(readmePath string, recursive bool) error {
	if recursive {
		return updateRunRecursive(readmePath)
	}
	return updateRunSingle(readmePath)
}

func updateRunRecursive(readmeFilename string) error {
	actionFiles, err := helpers.FindAllActionFiles(".")
	if err != nil {
		return err
	}
	
	if len(actionFiles) == 0 {
		return fmt.Errorf("no action.yml or action.yaml files found")
	}
	
	helpers.PrintHeader("Found %d action file(s)\n\n", len(actionFiles))
	
	updated := 0
	unchanged := 0
	
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	
	for _, actionPath := range actionFiles {
		dir := filepath.Dir(actionPath)
		readmePath := filepath.Join(dir, readmeFilename)
		
		wasUpdated, err := updateSingleActionWithResult(actionPath, readmePath)
		if err != nil {
			return fmt.Errorf("error updating %s: %w", readmePath, err)
		}
		
		if wasUpdated {
			fmt.Printf("%s Updated: %s\n", green("✓"), readmePath)
			updated++
		} else {
			fmt.Printf("%s Unchanged: %s\n", yellow("○"), readmePath)
			unchanged++
		}
	}
	
	helpers.PrintSummary(updated, "updated", color.FgGreen, unchanged, "unchanged", color.FgYellow)
	return nil
}

func updateSingleActionWithResult(actionPath, readmePath string) (bool, error) {
	actionParser := action.NewParser()
	a, err := actionParser.Parse(actionPath)
	if err != nil {
		return false, err
	}
	doc, err := markdown.NewDocOrCreate(readmePath)
	if err != nil {
		return false, err
	}
	oldDoc := doc.Copy()
	err = doc.Update(&a)
	if err != nil {
		return false, err
	}

	if doc.Equals(oldDoc) {
		return false, nil
	}

	err = doc.WriteToFile()
	if err != nil {
		return false, err
	}
	return true, nil
}

func updateSingleAction(actionPath, readmePath string) error {
	wasUpdated, err := updateSingleActionWithResult(actionPath, readmePath)
	if err != nil {
		return err
	}
	if wasUpdated {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s Updated: %s\n", green("✓"), readmePath)
	}
	return nil
}

func updateRunSingle(readmePath string) error {
	actionPath, err := helpers.FindActionFile()
	if err != nil {
		return err
	}

	return updateSingleAction(actionPath, readmePath)
}
