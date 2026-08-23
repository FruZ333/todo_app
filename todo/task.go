package todo

import "time"

type task struct {
	title  string
	text   string
	isDone bool

	createdAt time.Time
	doneAt    *time.Time // nil
}

func NewTask(title string, text string) task {
	return task{
		title:  title,
		text:   text,
		isDone: false,

		createdAt: time.Now(),
		doneAt:    nil,
	}
}
