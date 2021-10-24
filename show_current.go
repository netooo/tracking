package main

import (
	"fmt"
	"strconv"
	"time"

	tracking "github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli/v2"
)

func Current(c *cli.Context) error {
	histories, err := tracking.TrackRead(tracking.TodayPath)
	if err != nil {
		return err
	}

	if len(histories) > 0 {
		lastHistory := histories[len(histories)-1]
		if lastHistory.Duration == 0 {
			records := []string{
				"ID       " + strconv.Itoa(lastHistory.Task.ID),
				"Name     " + lastHistory.Task.Name,
				"Duration " + formatDuration(calcDuration(lastHistory.StartedAt)),
			}

			for _, record := range records {
				fmt.Println(record)
			}
			return nil
		}
	}

	fmt.Println("Not tracking")
	return nil
}

func calcDuration(startedAt time.Time) time.Duration {
	now := time.Now()
	return now.Sub(startedAt)
}

func formatDuration(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := duration / time.Minute % 60
	seconds := duration / time.Second % 60
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}
