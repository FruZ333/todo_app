package todo

import "time"

type Task struct {
	// структура задачи
	title  string
	text   string
	isDone bool

	createdAt time.Time
	doneAt    *time.Time // nil
}

func NewTask(title string, text string) Task {
	// конструктор задачи 
	return Task{
		title:  title,
		text:   text,
		isDone: false,

		createdAt: time.Now(),
		doneAt:    nil,
	}
}

func (t *Task) done() {
	// меняет статус и добавляет время выполнения задачи	
	t.isDone = true

	now := time.Now()
	t.doneAt = &now  
}
