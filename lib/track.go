package tracking

import (
	"context"
	"fmt"
	"math"
	"strconv"
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
	histories, err := TrackRead(TodayPath)
	if err != nil {
		return err
	}
	newHistories := append(histories, t)

	if err := Write(TodayPath, newHistories); err != nil {
		return err
	}

	return nil
}

func (t Track) Finish(ctx context.Context) error {
	histories, err := TrackRead(TodayPath)
	if err != nil {
		return err
	}

	newHistories := append(histories[:len(histories)-1], &t)

	if err := Write(TodayPath, newHistories); err != nil {
		return err
	}

	client, err := NewSheetClient(ctx)
	if err != nil {
		return err
	}

	hours := CalcHours(t, newHistories)

	cell, err := GetCell(t)
	if err != nil {
		return err
	}

	if err := client.Update(cell, [][]interface{}{
		{
			hours,
		},
	}); err != nil {
		return err
	}

	return nil
}

func CalcHours(t Track, histories []*Track) string {
	var minutes float64 = 0
	for _, h := range histories {
		if h.Task.ContentLine == t.Task.ContentLine {
			minutes += h.Duration.Minutes()
		}
	}

	return fmt.Sprintf("%.1f", math.Ceil(minutes)/60)
}

func GetCell(t Track) (string, error) {
	// time.Truncate がUTC規格で行われるので先に9時間ずらしておく
	now := time.Now().UTC().Add(9 * time.Hour).Truncate(time.Hour * 24)

	originDate_, err := GetConfigString("origin_date")
	if err != nil {
		return "", err
	}

	originDate, err := time.Parse("2006-01-02", originDate_)
	if err != nil {
		return "", err
	}

	diff := now.Sub(originDate)
	var days int = int(diff.Hours()/24) + int(diff.Hours()/24)/7

	originRow_, err := GetConfigString("origin_row")
	if err != nil {
		return "", err
	}

	var originRow []rune = []rune(originRow_)
	var quotient int = days

	for i := len(originRow) - 1; i >= 0; i-- {
		tmp := int(originRow[i]) - 65 + quotient
		quotient = tmp / 26
		remainder := tmp % 26

		originRow[i] = rune(remainder + 65)
	}

	var row string = string(originRow)
	if quotient > 0 {
		prefix := string(rune(64 + quotient))
		row = prefix + row
	}

	line := strconv.Itoa(t.Task.ContentLine)

	return row + line, nil
}
