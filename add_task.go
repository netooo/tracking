package main

import (
	"context"
	"errors"
	"github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
)

func AddTask(c *cli.Context) error {
	task := tracking.Task{}
	if !c.Args().Present() {
		return errors.New("command failed")
	}

	task.Name = c.Args().First()
	task.ContentID = c.Int("content-id")
	if err := task.Add(context.Background()); err != nil {
		return err
	}

	return nil
}
