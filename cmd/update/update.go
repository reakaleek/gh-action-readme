package update

import (
	"github.com/reakaleek/gh-action-readme/internal/action"
	"github.com/reakaleek/gh-action-readme/internal/markdown"
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	var actionPath string
	var readmePath string
	return &cli.Command{
		Name:  "update",
		Usage: "Update README.md",
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
			return updateRun(actionPath, readmePath)
		},
	}
}

func updateRun(actionPath string, readmePath string) error {
	actionParser := action.NewParser()
	a, err := actionParser.Parse(actionPath)
	if err != nil {
		return err
	}
	doc, err := markdown.NewDoc(readmePath)
	if err != nil {
		return err
	}
	oldDoc := doc.Copy()
	err = doc.Update(&a)
	if err != nil {
		return err
	}

	if doc.Equals(oldDoc) {
		return nil
	}

	err = doc.WriteToFile()
	if err != nil {
		return err
	}
	println(doc.GetName())
	return nil
}
