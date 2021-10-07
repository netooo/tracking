package main

import (
	"errors"
	"github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
	"time"
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
		if IsEnd(lastHistory.FinishedAt) {
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

func IsEnd(t time.Time) bool {
	return t.Hour() == 23 && t.Minute() == 59 && t.Second() == 59 && t.Nanosecond() == 999999999
}
