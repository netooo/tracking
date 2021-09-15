package tracking

import (
	"context"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentLine int    `json:"content_line"`
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
			t.ContentLine,
		},
	}); err != nil {
		panic(err)
	}

	return nil
}
