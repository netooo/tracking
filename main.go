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

	taskNameFlag := &cli.StringFlag{
		Name:  "name, n",
		Usage: "-name {task name}",
	}

	contentIDFlag := &cli.IntFlag{
		Name:  "content-id, C",
		Usage: "content id",
	}

	app.Commands = []*cli.Command{
		{
			Name:    "add-task",
			Aliases: []string{"at"},
			Usage:   "Add task",
			Action:  AddTask,
			Flags: []cli.Flag{
				taskNameFlag,
				contentIDFlag,
			},
			ArgsUsage: "<Task name>",
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
