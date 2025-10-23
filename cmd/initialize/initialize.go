package initialize

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/reakaleek/gh-action-readme/internal/helpers"
	"github.com/urfave/cli/v2"
)

//go:embed templates/*.md
var f embed.FS

func NewCommand() *cli.Command {
	var template string
	var readmePath string
	var recursive bool
	return &cli.Command{
		Name:  "init",
		Usage: "Initialize README.md",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "readme",
				Value:       "README.md",
				Destination: &readmePath,
			},
			&cli.StringFlag{
				Name:        "template",
				Value:       "default",
				Destination: &template,
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Value:       false,
				Destination: &recursive,
				Usage:       "Search recursively for all action.yml/action.yaml files and create README.md next to each",
			},
		},
		Action: func(ctx *cli.Context) error {
			return initRun(template, readmePath, recursive)
		},
	}
}

func initRun(template string, readmeFilename string, recursive bool) error {
	if recursive {
		return initRunRecursive(template, readmeFilename)
	}
	return initRunSingle(template, readmeFilename)
}

func initRunRecursive(template string, readmeFilename string) error {
	actionFiles, err := helpers.FindAllActionFiles(".")
	if err != nil {
		return err
	}

	if len(actionFiles) == 0 {
		return fmt.Errorf("no action.yml or action.yaml files found")
	}

	helpers.PrintHeader("Found %d action file(s)\n\n", len(actionFiles))

	created := 0
	skipped := 0
	
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	for _, actionPath := range actionFiles {
		dir := filepath.Dir(actionPath)
		readmePath := filepath.Join(dir, readmeFilename)

		// Check if README already exists
		if _, err := os.Stat(readmePath); err == nil {
			fmt.Printf("%s Skipped: %s (already exists)\n", yellow("○"), readmePath)
			skipped++
			continue
		}

		if err := createReadmeFromTemplate(template, readmePath); err != nil {
			return fmt.Errorf("error creating %s: %w", readmePath, err)
		}

		fmt.Printf("%s Created: %s\n", green("✓"), readmePath)
		created++
	}

	helpers.PrintSummary(created, "created", color.FgGreen, skipped, "skipped", color.FgYellow)
	return nil
}

func initRunSingle(template string, readmePath string) error {
	// Check if file already exists
	if _, err := os.Stat(readmePath); err == nil {
		return fmt.Errorf("%s already exists", readmePath)
	}

	if err := createReadmeFromTemplate(template, readmePath); err != nil {
		return err
	}

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s Created: %s\n", green("✓"), readmePath)
	return nil
}

func createReadmeFromTemplate(template string, readmePath string) error {
	switch template {
	case "default":
		file, err := f.ReadFile("templates/default.md")
		if err != nil {
			return err
		}
		return writeToFile(readmePath, string(file))
	default:
		return fmt.Errorf("unknown template: %s", template)
	}
}

func writeToFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0755)
}
