package tracking

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

var (
	trackFile = "track.json"
	trackPath = filepath.Join(cacheDir, trackFile)
	tracks    []*Track
)

type Track struct {
	TaskID     int           `json:"task_id"`
	TaskName   string        `json:"task_name"`
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at"`
	Duration   time.Duration `json:"duration"`
}

func (t *Track) Start() error {
	histories, err := TrackRead()
	if err != nil {
		return err
	}
	newHistories := append(histories, t)

	buf, err := json.Marshal(newHistories)
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
