package todo

type List struct {
	tasks map[string]task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]task),
	}
}

func (l *List) AddTask(task task) {
	// создает новую задачу
	l.tasks[task.title] = task
}

func (l *List) ListTasks() map[string]task {
	// выводит все задачи
	return l.tasks
}

func (l *List) DelTask(title string) string {
	// удаляет задачу по заголовку
	_, ok := l.tasks[title]
	if !ok {
		return errorNotFound
	}

	delete(l.tasks, title)

	return ""
}

func (l *List) DoneTask(title string) string {
	// меняет статус выполнения и создает время выполнения по заголовку задачи
	task, ok := l.tasks[title]
	if !ok {
		return errorNotFound
	}
	
	task.done()

	l.tasks[title] = task   // обновляем задачу

	return ""
}


