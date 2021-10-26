package main

import (
	"context"
	"errors"
	"time"

	tracking "github.com/netooo/tracking/lib"
	"github.com/urfave/cli/v2"
)

func Finish(c *cli.Context) error {
	histories, err := tracking.TrackRead(tracking.TodayPath)
	if err != nil {
		return err
	}

	if len(histories) == 0 {
		return errors.New("task is not running")
	}

	lastHistory := histories[len(histories)-1]
	if lastHistory.Duration > 0 {
		return errors.New("task is not running")
	}

	now := time.Now()
	lastHistory.FinishedAt = now
	lastHistory.Duration = now.Sub(lastHistory.StartedAt)

	if err := lastHistory.Finish(context.Background()); err != nil {
		return err
	}

	return nil
}
