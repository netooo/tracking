package main

import (
	"context"
	"errors"
	"github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
)

func AddTask(c *cli.Context) error {
	task := tracking.Task{}

	if c.String("name") == "" {
		return errors.New("command failed")
	}

	task.Name = c.String("name")
	task.ContentID = c.Int("content")
	if err := task.Add(context.Background()); err != nil {
		return err
	}

	return nil
}
