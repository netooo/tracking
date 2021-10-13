package main

import (
	"errors"
	"time"

	tracking "github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
)

func Start(c *cli.Context) error {

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

	histories, err := tracking.TrackRead()
	if err != nil {
		return err
	}

	if len(histories) > 0 {
		lastHistory := histories[len(histories)-1]
		if lastHistory.Duration == 0 {
			return errors.New("task is running")
		}
	}

	now := time.Now()
	track := tracking.Track{
		Task:      task,
		StartedAt: now,
		FinishedAt: time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			23, 59, 59, 999999999, time.Local,
		),
	}

	if err := track.Start(); err != nil {
		return err
	}

	return nil
}
