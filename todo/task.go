package todo

import "time"

type Task struct {
	// структура задачи
	title  string
	text   string
	isDone bool

	CreatedAt time.Time
	DoneAt    *time.Time // nil
}

func NewTask(title string, text string) Task {
	// конструктор задачи 
	return Task{
		title:  title,
		text:   text,
		isDone: false,

		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}

func (t *Task) done() {
	// меняет статус и добавляет время выполнения задачи	
	t.isDone = true

	now := time.Now()
	t.DoneAt = &now  
}
