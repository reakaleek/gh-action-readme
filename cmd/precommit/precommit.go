package precommit

import (
	"errors"
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/reakaleek/gh-action-readme/internal/markdown"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

func NewCommand() *cli.Command {
	var envs cli.StringSlice
	return &cli.Command{
		Name:  "pre-commit",
		Usage: "Pre-commit hook to update README.md",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "env",
				Usage:       "Pass multiple environment variables in the format --env=key=value",
				Destination: &envs,
			},
		},
		Hidden: true,
		Args:   true,
		Action: func(ctx *cli.Context) error {
			for _, env := range envs.Value() {
				tokens := strings.Split(env, "=")
				if len(tokens) != 2 {
					return errors.New("invalid env format. should be key=value")
				}
				err := os.Setenv(tokens[0], tokens[1])
				if err != nil {
					return err
				}
			}
			files := ctx.Args().Slice()
			for _, file := range files {
				dir := filepath.Dir(file)
				actionPath := filepath.Join(dir, "action.yml")
				readmePath := filepath.Join(dir, "README.md")
				if !fileExists(actionPath) {
					continue
				}
				doc, err := markdown.NewDoc(readmePath)
				oldDoc := doc.Copy()
				if err != nil {
					return err
				}
				parser := action.NewParser()
				a, err := parser.Parse(actionPath)
				if err != nil {
					return err
				}
				err = doc.Update(&a)
				if err != nil {
					return err
				}
				if !oldDoc.Equals(*doc) {
					err = doc.WriteToFile()
					if err != nil {
						return err
					}
					println(doc.GetName())
				}
			}
			return nil
		},
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
