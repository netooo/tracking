package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tracking "github.com/netooo/tracking/lib"

	"github.com/urfave/cli/v2"
)

func Log(c *cli.Context) error {
	var logPath string

	if c.String("date") == "" {
		logPath = tracking.TodayPath
	} else {
		dateStr := c.String("date")
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return err
		}
		logPath = filepath.Join(tracking.CacheDir, dateStr+".json")
	}

	histories, err := tracking.TrackRead(logPath)
	if err != nil {
		return err
	}

	for _, log := range histories {
		started := log.StartedAt.Format("15:04")
		finished := log.FinishedAt.Format("15:04")
		duration := strconv.FormatFloat(log.Duration.Minutes(), 'f', 0, 64)
		if log.Duration == 0 {
			finished = " Now "
			duration_ := time.Now().Sub(log.StartedAt).Minutes()
			duration = strconv.FormatFloat(duration_, 'f', 0, 64)
		}
		space := strings.Repeat(" ", 4-len(duration))

		fmt.Println(started + "-" + finished + "(" + duration + "m)" + space + log.Task.Name)
	}

	return nil
}
