package todo

type list struct {
	tasks map[string]task
}

func NewList() *list {
	return &list{
		tasks: make(map[string]task),
	}
}

func (l *list) addTask(task task) {
	// создает новую задачу
	l.tasks[task.title] = task
}

func (l *list) listTasks() map[string]task {
	// выводит все задачи
	return l.tasks
}

func (l *list) delTask(title string) string {
	// удаляет задачу по заголовку
	_, ok := l.tasks[title]
	if !ok {
		return errorNotFound
	}

	delete(l.tasks, title)

	return ""
}

func (l *list) doneTask(title string) string {
	// меняет статус выполнения и создает время выполнения по заголовку задачи
	task, ok := l.tasks[title]
	if !ok {
		return errorNotFound
	}
	
	task.done()

	l.tasks[title] = task   // обновляем задачу

	return ""
}


