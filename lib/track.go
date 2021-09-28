package tracking

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	trackFile = "track.json"
	trackPath = filepath.Join(cacheDir, trackFile)
	tracks    []Track
)

type Track struct {
	TaskID   int    `json:"task_id"`
	TaskName string `json:"task_name"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Duration int64  `json:"duration"`
}

func (t *Track) Start() error {
	logs, err := TrackRead()
	if err != nil {
		return err
	}
	newLogs := append(*logs, *t)

	buf, err := json.Marshal(newLogs)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(trackPath, os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
