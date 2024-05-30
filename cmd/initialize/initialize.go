package initialize

import (
	"embed"
	"github.com/urfave/cli/v2"
	"os"
)

//go:embed templates/*.md
var f embed.FS

func NewCommand() *cli.Command {
	var template string
	var readmePath string
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
		},
		Action: func(ctx *cli.Context) error {
			return initRun(template, readmePath)
		},
	}
}

func initRun(template string, readmePath string) error {
	switch template {
	case "default":
		file, err := f.ReadFile("templates/default.md")
		if err != nil {
			return err
		}
		err = writeToFile(readmePath, string(file))
	}
	return nil
}

func writeToFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0755)
}
