package tracking

import (
	"encoding/json"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	cacheDir = filepath.Join(basedir.CacheHome, "tracking")
)

func CacheInit() error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDir, 0777); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if _, err := os.Stat(trackPath); os.IsNotExist(err) {
		fp, err := os.Create(trackPath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}
	if _, err := os.Stat(taskPath); os.IsNotExist(err) {
		fp, err := os.Create(taskPath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}

	return nil
}

func TaskRead() (*[]Task, error) {
	jsonBytes, err := ioutil.ReadFile(taskPath)
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

func TrackRead() ([]*Track, error) {
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
