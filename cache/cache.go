package cache

import (
	"encoding/json"
	"github.com/netooo/TimeTracking/lib"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"path/filepath"
)

var c *Cache

type Cache struct {
	Filename string `json:"filename"`
	TaskID   int    `json:"task_id"`
	Task     *tracking.Task
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Duration int64  `json:"duration"`
}

func New() *Cache {
	c = new(Cache)

	cachePath := filepath.Join(basedir.CacheHome, "tracking", "cache.json")
	c.Filename = cachePath
	return c
}

func Init() { c.Init() }
func (c *Cache) Init() error {
	if err := c.Read(); err != nil {
		if err = c.Write(); err != nil {
			return err
		}
	}
	return nil
}

func Read() { c.Read() }
func (c *Cache) Read() error {
	jsonString, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonString, &c.Content); err != nil {
		return err
	}
	return nil
}

func Write() { c.Write() }
func (c *Cache) Write() error {
	buf, err := json.MarshalIndent(c.Content, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(c.Filename, buf, os.ModePerm); err != nil {
		return err
	}
	return nil
}
