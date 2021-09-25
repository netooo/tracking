package cache

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	trackFile = "track.json"
	tracks    []Track
)

type Track struct {
	TaskID   int    `json:"task_id"`
	TaskName string `json:"task_name"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Duration int64  `json:"duration"`
}

func TrackRead() (*[]Track, error) {
	logs, err := tr.Read()
	return logs, err
}
func (tr *TrackCache) Read() (*[]Track, error) {
	jsonBytes, err := ioutil.ReadFile(tr.Filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return &tracks, nil
	}

	err = json.Unmarshal(jsonBytes, &tracks)
	if err != nil {
		log.Fatal(err)
	}

	return &tracks, nil
}

func TrackWrite() { tr.Write() }
func (tr *TrackCache) Write() error {
	logs, err := TrackRead()
	newLogs := append(*logs, tr.Track)

	buf, err := json.Marshal(newLogs)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(tr.Filename, os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}

	return nil
}
