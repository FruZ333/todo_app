package todo

type list struct {
	tasks map[string]task
}

func NewList() *list {
	return &list{
		tasks: make(map[string]task),
	}
}