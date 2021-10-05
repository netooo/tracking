package tracking

import (
	"time"
)

var (
	tracks []*Track
)

type Track struct {
	Task       Task
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
