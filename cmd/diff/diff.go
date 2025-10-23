package diff

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
	return &cli.Command{
		Name:  "diff",
		Usage: "Diff README.md",
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
			return diffRun(readmePath, recursive)
		},
	}
}

func diffRun(readmePath string, recursive bool) error {
	if recursive {
		return diffRunRecursive(readmePath)
	}
	return diffRunSingle(readmePath)
}

func diffRunRecursive(readmeFilename string) error {
	actionFiles, err := helpers.FindAllActionFiles(".")
	if err != nil {
		return err
	}
	
	if len(actionFiles) == 0 {
		return fmt.Errorf("no action.yml or action.yaml files found")
	}
	
	helpers.PrintHeader("Found %d action file(s)\n\n", len(actionFiles))
	
	hasAnyDiff := false
	upToDate := 0
	outOfDate := 0
	
	for _, actionPath := range actionFiles {
		dir := filepath.Dir(actionPath)
		readmePath := filepath.Join(dir, readmeFilename)
		
		hasDiff, err := diffSingleActionWithOutput(actionPath, readmePath)
		if err != nil {
			return fmt.Errorf("error diffing %s: %w", readmePath, err)
		}
		if hasDiff {
			hasAnyDiff = true
			outOfDate++
		} else {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("%s %s\n", green("✓"), readmePath)
			upToDate++
		}
	}
	
	helpers.PrintSummary(upToDate, "up-to-date", color.FgGreen, outOfDate, "out-of-date", color.FgRed)
	
	if hasAnyDiff {
		return cli.Exit("", 1)
	}
	
	return nil
}

func diffSingleActionWithOutput(actionPath, readmePath string) (bool, error) {
	actionParser := action.NewParser()
	a, err := actionParser.Parse(actionPath)
	if err != nil {
		return false, err
	}
	doc, err := markdown.NewDoc(readmePath)
	if err != nil {
		return false, err
	}
	newDoc := doc.Copy()
	err = doc.Update(&a)
	if err != nil {
		return false, err
	}
	diff := newDoc.Diff(doc)
	if diff.HasDiff {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("%s %s\n\n", red("✗"), readmePath)
		fmt.Println(diff.PrettyDiff)
		fmt.Println()
		return true, nil
	}
	return false, nil
}

func diffSingleAction(actionPath, readmePath string) (bool, error) {
	actionParser := action.NewParser()
	a, err := actionParser.Parse(actionPath)
	if err != nil {
		return false, err
	}
	doc, err := markdown.NewDoc(readmePath)
	if err != nil {
		return false, err
	}
	newDoc := doc.Copy()
	err = doc.Update(&a)
	if err != nil {
		return false, err
	}
	diff := newDoc.Diff(doc)
	if diff.HasDiff {
		fmt.Printf("\n%s\n\n", readmePath)
		fmt.Println(diff.PrettyDiff)
		return true, nil
	}
	return false, nil
}

func diffRunSingle(readmePath string) error {
	actionPath, err := helpers.FindActionFile()
	if err != nil {
		return err
	}

	hasDiff, err := diffSingleAction(actionPath, readmePath)
	if err != nil {
		return err
	}
	
	if hasDiff {
		return cli.Exit(fmt.Sprintf("%s is not up-to-date", readmePath), 1)
	}
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s %s is up-to-date\n", green("✓"), readmePath)
	return nil
}
