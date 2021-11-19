package main

import (
	"errors"

	tracking "github.com/netooo/tracking/lib"
	"github.com/urfave/cli/v2"
)

func Delete(c *cli.Context) error {

	taskID := c.Int("task")
	if taskID == 0 {
		return errors.New("command failed")
	}

	taskList, err := tracking.TaskRead()
	if err != nil {
		return err
	}

	ok := false
	task := tracking.Task{}
	for _, t := range taskList {
		if taskID == t.ID {
			ok = true
			task = *t
			break
		}
	}
	if !ok {
		return errors.New("task not found")
	}

	if err := task.Delete(); err != nil {
		return err
	}

	return nil
}
