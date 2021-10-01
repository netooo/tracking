package tracking

import (
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

	if err := Write(newHistories); err != nil {
		return err
	}

	return nil
}

func (t Track) Finish() error {
	histories, err := TrackRead()
	if err != nil {
		return err
	}

	newHistories := append(histories[:len(histories)-1], &t)

	if err := Write(newHistories); err != nil {
		return err
	}

	return nil
}
