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
	ConfigDir  = filepath.Join(basedir.ConfigHome, "tracking")
	ConfigPath = filepath.Join(ConfigDir, "config.json")
	SecretPath = filepath.Join(ConfigDir, "secret.json")

	CacheDir  = filepath.Join(basedir.CacheHome, "tracking")
	TaskPath  = filepath.Join(CacheDir, "task.json")
	today     = time.Now().Format("2006-01-02")
	TodayPath = filepath.Join(CacheDir, today+".json")
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
