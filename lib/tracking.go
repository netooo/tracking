package tracking

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/rkoesters/xdg/basedir"
)

var (
	configDir  = filepath.Join(basedir.ConfigHome, "tracking")
	ConfigPath = filepath.Join(configDir, "config.json")
	SecretPath = filepath.Join(configDir, "secret.json")

	cacheDir  = filepath.Join(basedir.CacheHome, "tracking")
	TaskPath  = filepath.Join(cacheDir, "task.json")
	today     = time.Now().Format("20060102")
	TodayPath = filepath.Join(cacheDir, today+".json")
)

func GetConfigString(s string) (string, error) {
	configBlob, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return "", errors.New("command failed")
	}

	var configJson interface{}
	err = json.Unmarshal(configBlob, &configJson)
	if err != nil {
		return "", errors.New("command failed")
	}
	sheetParam := configJson.(map[string]interface{})[s].(string)

	return sheetParam, nil
}
