package main

import (
	"fmt"
	"github.com/moznion/go-unicode-east-asian-width"
	tracking "github.com/netooo/tracking/lib"
	"github.com/urfave/cli/v2"
	"strings"
)

func List(c *cli.Context) error {
	taskList, err := tracking.TaskRead()
	if err != nil {
		return err
	}

	maxPadding := longestNameCnt(taskList) + 1
	for _, task := range taskList {
		fmt.Printf("%-9d", task.ID)
		fmt.Printf(task.Name)
		spacePad := maxPadding - countInName(task.Name)
		space := strings.Repeat(" ", spacePad)
		fmt.Printf(space)
		fmt.Printf("%3d ", task.ContentLine)
		fmt.Printf("%s\n", task.IssueId)
	}

	return nil
}

func longestNameCnt(tasks []*tracking.Task) int {
	max := 0
	for _, task := range tasks {
		cnt := countInName(task.Name)
		if cnt > max {
			max = cnt
		}
	}

	return max
}

func countInName(name string) int {
	cnt := 0
	for _, r := range []rune(name) {
		if eastasianwidth.IsFullwidth(r) {
			cnt += 2
		} else {
			cnt += 1
		}
	}

	return cnt
}
