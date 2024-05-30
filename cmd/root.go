package cmd

import (
	"github.com/reakaleek/gh-action-readme/cmd/diff"
	"github.com/reakaleek/gh-action-readme/cmd/initialize"
	"github.com/reakaleek/gh-action-readme/cmd/update"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Execute() {
	app := &cli.App{
		Name:  "gh action-readme",
		Usage: "Generate or update GitHub Actions documentation.",
		Commands: []*cli.Command{
			diff.NewCommand(),
			update.NewCommand(),
			initialize.NewCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
