package cache

import (
	"encoding/json"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	c         *Cache
	cacheDir  = filepath.Join(basedir.CacheHome, "tracking")
	cacheFile = "cache.json"
)

type Cache struct {
	Filename string `json:"filename"`
	Content  Content
}

type Content struct {
	TaskID   int    `json:"task_id"`
	TaskName string `json:"task_name"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Duration int64  `json:"duration"`
}

func New() *Cache {
	c = new(Cache)

	cachePath := filepath.Join(cacheDir, cacheFile)
	c.Filename = cachePath
	return c
}

func Init() { c.Init() }
func (c *Cache) Init() error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDir, 0777); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if _, err := os.Stat(c.Filename); os.IsNotExist(err) {
		fp, err := os.Create(c.Filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer fp.Close()
	}
	return nil
}

func Read() { c.Read() }
func (c *Cache) Read() error {
	jsonBytes, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		return err
	}

	return nil
}

func Write() { c.Write() }
func (c *Cache) Write() error {
	buf, err := json.Marshal(c.Content)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(c.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf = append(buf, []byte("\n")...)
	if _, err = f.Write(buf); err != nil {
		panic(err)
	}

	return nil
}
