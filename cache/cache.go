package cache

import (
	"github.com/rkoesters/xdg/basedir"
	"log"
	"os"
	"path/filepath"
)

var (
	tr       *TrackCache
	ta       *TaskCache
	cacheDir = filepath.Join(basedir.CacheHome, "tracking")
)

type TrackCache struct {
	Filename string `json:"filename"`
	Track    Track
}

type TaskCache struct {
	Filename string `json:"filename"`
	Task     Task
}

func New() {
	TrackNew()
	TaskNew()
}
func TrackNew() *TrackCache {
	tr = new(TrackCache)

	trackPath := filepath.Join(cacheDir, trackFile)
	tr.Filename = trackPath
	return tr
}

func TaskNew() *TaskCache {
	ta = new(TaskCache)

	taskPath := filepath.Join(cacheDir, taskFile)
	ta.Filename = taskPath
	return ta
}

func Init() error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDir, 0777); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if _, err := os.Stat(tr.Filename); os.IsNotExist(err) {
		fp, err := os.Create(tr.Filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}
	if _, err := os.Stat(ta.Filename); os.IsNotExist(err) {
		fp, err := os.Create(ta.Filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}

	return nil
}
