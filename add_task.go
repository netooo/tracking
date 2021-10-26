package main

import (
	"errors"
	"math/rand"
	"time"

	tracking "github.com/netooo/tracking/lib"
	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	task := tracking.Task{}

	if c.String("name") == "" {
		return errors.New("command failed")
	}
	if c.Int("line") == 0 {
		return errors.New("command failed")
	}

	rand.Seed(time.Now().UnixNano())
	task.ID = rand.Intn(99999999) + 1
	task.Name = c.String("name")
	task.ContentLine = c.Int("line")

	if err := task.Add(); err != nil {
		return err
	}

	return nil
}
