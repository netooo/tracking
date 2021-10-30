package tracking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CacheInit() error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDir, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(TodayPath); os.IsNotExist(err) {
		fp, err := os.Create(TodayPath)
		if err != nil {
			return err
		}
		defer fp.Close()

		// 一週間前まで遡って日跨ぎのトラッキングを修正
		// Durationを更新
		// Sheetには書き込まない
		for i := 1; i < 8; i++ {
			day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
			dayPath := filepath.Join(cacheDir, day+".json")

			if _, err := os.Stat(dayPath); !os.IsNotExist(err) {
				histories, err := TrackRead(dayPath)
				if err != nil {
					log.Fatal(err)
					return err
				}

				if len(histories) > 0 {
					lastHistory := histories[len(histories)-1]

					if lastHistory.Duration == 0 {
						lastHistory.Duration = lastHistory.FinishedAt.Sub(lastHistory.StartedAt)
						newHistories := append(histories[:len(histories)-1], lastHistory)

						if err := Write(dayPath, newHistories); err != nil {
							return err
						}
					}

					break
				}
			}
		}
	}

	if _, err := os.Stat(TaskPath); os.IsNotExist(err) {
		fp, err := os.Create(TaskPath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}

	return nil
}

func TaskRead() ([]*Task, error) {
	jsonBytes, err := ioutil.ReadFile(TaskPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(jsonBytes, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks, nil
}

func TrackRead(trackPath string) ([]*Track, error) {
	jsonBytes, err := ioutil.ReadFile(trackPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return tracks, nil
	}

	err = json.Unmarshal(jsonBytes, &tracks)
	if err != nil {
		log.Fatal(err)
	}

	return tracks, nil
}

func Write(trackPath string, histories []*Track) error {
	buf, err := json.Marshal(histories)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(trackPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
