package main

import (
	"fmt"
	"github.com/netooo/TimeTracking/cache"
	"github.com/urfave/cli"
	"os"
)

func main() {
	cache.New()
	cache.Init()

	app := cli.NewApp()
	app.Name = "Tracking"
	app.Usage = "Tracking CLI Client"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	taskNameFlag := &cli.StringFlag{
		Name:  "name, n",
		Usage: "-name {task name}",
	}

	contentLineFlag := &cli.IntFlag{
		Name:  "content, c",
		Usage: "-content {content line number}",
	}

	taskIDFlag := &cli.IntFlag{
		Name:  "task, t",
		Usage: "-task {task id}",
	}

	app.Commands = []*cli.Command{
		{
			Name:    "add-task",
			Aliases: []string{"at"},
			Usage:   "Add task",
			Action:  AddTask,
			Flags: []cli.Flag{
				taskNameFlag,
				contentLineFlag,
			},
			ArgsUsage: "<Task name>",
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start Task",
			Action:  StartTask,
			Flags: []cli.Flag{
				taskIDFlag,
			},
			ArgsUsage: "<Task name>",
		},
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
