package main

import (
	"errors"
	"github.com/netooo/TimeTracking/cache"
	"github.com/urfave/cli"
)

func StartTask(c *cli.Context) error {
	if c.Int("task") == 0 {
		return errors.New("command failed")
	}

	cache.Write()
	cache.Read()

	return nil
}
