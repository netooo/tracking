package tracking

import (
	"encoding/json"
	"errors"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"path/filepath"
	"time"
)

var (
	configDir  = filepath.Join(basedir.ConfigHome, "tracking")
	configPath = filepath.Join(configDir, "config.json")
	secretPath = filepath.Join(configDir, "secret.json")

	cacheDir  = filepath.Join(basedir.CacheHome, "tracking")
	taskPath  = filepath.Join(cacheDir, "task.json")
	today     = time.Now().Format("20060102")
	trackPath = filepath.Join(cacheDir, today+".json")
)

func GetConfigString(s string) (string, error) {
	configBlob, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", errors.New("command failed")
	}

	var configJson interface{}
	err = json.Unmarshal(configBlob, &configJson)
	if err != nil {
		return "", errors.New("command failed")
	}
	sheetId := configJson.(map[string]interface{})[s].(string)

	return sheetId, nil
}
