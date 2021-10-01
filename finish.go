package main

import (
	"errors"
	"github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
	"time"
)

func Finish(c *cli.Context) error {
	histories, err := tracking.TrackRead()
	if err != nil {
		return err
	}

	if len(histories) == 0 {
		return errors.New("task is not running")
	}

	lastHistory := histories[len(histories)-1]
	if !lastHistory.FinishedAt.IsZero() {
		return errors.New("task is not running")
	}

	now := time.Now()
	lastHistory.FinishedAt = now
	lastHistory.Duration = now.Sub(lastHistory.StartedAt)

	if err := lastHistory.Finish(); err != nil {
		return err
	}

	return nil
}
