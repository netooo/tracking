package main

import (
	"fmt"
	"os"

	tracking "github.com/netooo/tracking/lib"
	"github.com/urfave/cli/v2"
)

func main() {
	tracking.CacheInit()

	app := cli.NewApp()
	app.Name = "tracking"
	app.Usage = "tracking CLI Client"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	taskNameFlag := &cli.StringFlag{
		Name:  "name, n",
		Usage: "-name {task name}",
	}

	contentLineFlag := &cli.IntFlag{
		Name:  "line, l",
		Usage: "-line {content line number}",
	}

	taskIDFlag := &cli.IntFlag{
		Name:  "task, t",
		Usage: "-task {task id}",
	}

	app.Commands = []*cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "Show task list",
			Action:    List,
			Flags:     []cli.Flag{},
			ArgsUsage: "",
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add task",
			Action:  Add,
			Flags: []cli.Flag{
				taskNameFlag,
				contentLineFlag,
			},
			ArgsUsage: "",
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start Task",
			Action:  Start,
			Flags: []cli.Flag{
				taskIDFlag,
			},
			ArgsUsage: "",
		},
		{
			Name:      "finish",
			Aliases:   []string{"f"},
			Usage:     "Finish Task",
			Action:    Finish,
			Flags:     []cli.Flag{},
			ArgsUsage: "",
		},
		{
			Name:      "current",
			Aliases:   []string{"c"},
			Usage:     "Show current tracking",
			Action:    Current,
			Flags:     []cli.Flag{},
			ArgsUsage: "",
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
