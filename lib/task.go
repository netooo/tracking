package tracking

import (
	"context"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ContentID int    `json:"content_id"`
}

func (t *Task) Add(ctx context.Context) error {
	client, err := NewSheetClient(ctx)
	if err != nil {
		panic(err)
	}

	if err := client.Append("Task", [][]interface{}{
		{
			t.ID,
			t.Name,
			t.ContentID,
		},
	}); err != nil {
		panic(err)
	}

	return nil
}
