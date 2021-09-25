package cache

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	taskFile = "task.json"
	tasks    []Task
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentLine int    `json:"content_line"`
}

func TaskRead() (*[]Task, error) {
	logs, err := ta.Read()
	return logs, err
}
func (ta *TaskCache) Read() (*[]Task, error) {
	jsonBytes, err := ioutil.ReadFile(ta.Filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return &tasks, nil
	}

	err = json.Unmarshal(jsonBytes, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return &tasks, nil
}

func TaskWrite() { ta.Write() }
func (ta *TaskCache) Write() error {
	logs, err := TaskRead()
	newLogs := append(*logs, ta.Task)

	buf, err := json.Marshal(newLogs)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(ta.Filename, os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}

	return nil
}
