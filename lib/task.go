package tracking

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	tasks []*Task
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentLine int    `json:"content_line"`
	IssueId     string `json:"issue_id"`
}

func (t *Task) Add() error {
	taskList, err := TaskRead()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		if t.ID == task.ID {
			return errors.New("the task already exists")
		}
	}

	newTasks := append(taskList, t)
	buf, err := json.Marshal(newTasks)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(TaskPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}

func (t *Task) Delete() error {
	taskList, err := TaskRead()
	if err != nil {
		return err
	}

	var newTasks []*Task
	for i, task := range taskList {
		if t.ID == task.ID {
			newTasks = taskList[:i+copy(taskList[i:], taskList[i+1:])]
			break
		}
	}

	buf, err := json.Marshal(newTasks)
	if err != nil {
		return err
	}

	f, err := os.Create(TaskPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
