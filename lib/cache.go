package tracking

var c *Cache

type Cache struct {
	Filename string `json:"filename"`
	TaskID   int    `json:"task_id"`
	Task     *Task
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Duration int64  `json:"duration"`
}

func New(filename string) *Cache {
	c = new(Cache)
	c.Filename = filename
	return c
}

func (c *Cache) Init() error {
	if err := c.Read(); err != nil {
		if err = c.Write(); err != nil {
			return err
		}
	}
	return nil
}
