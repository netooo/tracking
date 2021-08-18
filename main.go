package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Tracking"
	app.Usage = "Tracking CLI Client"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	projectIDFlag := &cli.IntFlag{
		Name:  "project-id, P",
		Usage: "project id",
	}

	app.Commands = []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add task",
			Action:  Add,
			Flags: []cli.Flag{
				projectIDFlag,
			},
			ArgsUsage: "<Task name>",
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
