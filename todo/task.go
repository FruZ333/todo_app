package todo

import "time"

type task struct {
	// структура задачи
	title  string
	text   string
	isDone bool

	createdAt time.Time
	doneAt    *time.Time // nil
}

func NewTask(title string, text string) task {
	// конструктор задачи 
	return task{
		title:  title,
		text:   text,
		isDone: false,

		createdAt: time.Now(),
		doneAt:    nil,
	}
}

func (t *task) done() {
	// меняет статус и добавляет время выполнения задачи	
	t.isDone = true

	now := time.Now()
	t.doneAt = &now  
}
