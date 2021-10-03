package tracking

import (
	"github.com/rkoesters/xdg/basedir"
	"path/filepath"
)

var (
	configDir  = filepath.Join(basedir.ConfigHome, "tracking")
	configPath = filepath.Join(configDir, "config.json")
	secretPath = filepath.Join(configDir, "secret.json")

	cacheDir  = filepath.Join(basedir.CacheHome, "tracking")
	taskPath  = filepath.Join(cacheDir, "task.json")
	trackPath = filepath.Join(cacheDir, "track.json")
)
