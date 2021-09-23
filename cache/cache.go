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
	contents  []Content
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

func Read() (*[]Content, error) {
	logs, err := c.Read()
	return logs, err
}
func (c *Cache) Read() (*[]Content, error) {
	jsonBytes, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return &contents, nil
	}

	err = json.Unmarshal(jsonBytes, &contents)
	if err != nil {
		log.Fatal(err)
	}

	return &contents, nil
}

func Write() { c.Write() }
func (c *Cache) Write() error {
	logs, err := Read()
	newLogs := append(*logs, c.Content)

	buf, err := json.Marshal(newLogs)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(c.Filename, os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}

	return nil
}
