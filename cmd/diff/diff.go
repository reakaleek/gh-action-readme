package diff

import (
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/reakaleek/gh-action-readme/internal/markdown"
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	var actionPath string
	var readmePath string
	return &cli.Command{
		Name:  "diff",
		Usage: "Diff README.md",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "action",
				Value:       "action.yml",
				Destination: &actionPath,
			},
			&cli.StringFlag{
				Name:        "readme",
				Value:       "README.md",
				Destination: &readmePath,
			},
		},
		Action: func(ctx *cli.Context) error {
			return diffRun(actionPath, readmePath)
		},
	}
}

func diffRun(actionPath string, readmePath string) error {
	actionParser := action.NewParser()
	a, err := actionParser.Parse(actionPath)
	if err != nil {
		return err
	}
	doc, err := markdown.NewDoc(readmePath)
	if err != nil {
		return err
	}
	newDoc := doc.Copy()
	err = doc.Update(&a)
	if err != nil {
		return err
	}
	diff := newDoc.Diff(doc)
	if diff.HasDiff {
		println(diff.PrettyDiff)
		return cli.Exit("README.md is not up-to-date", 1)
	} else {
		return cli.Exit("README.md is up-to-date", 0)
	}
}
