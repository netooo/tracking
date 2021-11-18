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
	dateFlag := &cli.StringFlag{
		Name:  "date, d",
		Usage: "-date {yyyy-mm-dd}",
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
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete task",
			Action:  Delete,
			Flags: []cli.Flag{
				taskIDFlag,
			},
			ArgsUsage: "",
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start task",
			Action:  Start,
			Flags: []cli.Flag{
				taskIDFlag,
			},
			ArgsUsage: "",
		},
		{
			Name:      "finish",
			Aliases:   []string{"f"},
			Usage:     "Finish task",
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
		{
			Name:    "log",
			Aliases: []string{},
			Usage:   "Show tracking logs",
			Action:  Log,
			Flags: []cli.Flag{
				dateFlag,
			},
			ArgsUsage: "",
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
