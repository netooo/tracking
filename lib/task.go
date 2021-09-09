package tracking

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ContentID int    `json:"content_id"`
}
