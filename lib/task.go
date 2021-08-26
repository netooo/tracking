package tracking

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	HashID    int    `json:"hash_id"`
	ContentID int    `json:"content_id"`
}
