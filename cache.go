package main

import (
	"fmt"

	tracking "github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
)

func Cache(c *cli.Context) error {
	taskList, err := tracking.TaskRead()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		fmt.Printf("%-9d%-4d%s\n", task.ID, task.ContentLine, task.Name)
	}

	return nil
}
