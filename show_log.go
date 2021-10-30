package main

import (
	"path/filepath"
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
		//fmt.Println(log)
	}

	return nil
}
