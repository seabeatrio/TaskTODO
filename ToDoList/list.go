package ToDoList

type TaskFilter struct {
	Title     string
	Completed *bool
}
type List struct {
	Tasks map[string]*Task
}

func NewList() *List {
	return &List{
		make(map[string]*Task),
	}
}
func (l *List) CreateTask(title string, description string) error {
	if _, ok := l.Tasks[title]; ok {
		return ErrorTask{
			Message: "Task Exist Already",
			Type:    ErrExistAlready,
		}
	}
	l.Tasks[title] = NewTask(title, description)
	return nil

}
func (l *List) GetTask(title string) (*Task, error) {
	v, ok := l.Tasks[title]
	if !ok {
		return nil, ErrorTask{
			Message: "Task not found",
			Type:    ErrNotFound,
		}
	}
	return v, nil
}
func (l *List) GetTasks(filter TaskFilter) map[string]*Task {
	if filter.Title == "" {
		return l.Tasks
	}
	result := make(map[string]*Task)

	for key, task := range l.Tasks {
		if filter.Title == task.Title {
			result[key] = task
		} else if *filter.Completed == task.IsDone {
			result[key] = task
		}
	}
	return result
}
func (l *List) DoneTask(title string) error {
	v, ok := l.Tasks[title]
	if !ok {
		return ErrorTask{
			Message: "Task not found",
			Type:    ErrNotFound,
		}
	}
	v.ToggleStatus()
	return nil
}
func (l *List) ChangeTask(
	title string,
	newTitle string,
	description string,
	completed *bool) error {
	task, err := l.GetTask(title)

	if err != nil {
		return ErrorTask{
			Message: "Task not found",
			Type:    ErrNotFound,
		}
	}
	if newTitle == "" {
		return ErrorTask{
			Message: "Task Validated Invalid",
			Type:    ErrValidation,
		}
	}
	if completed != nil {
		l.Tasks[title].ToggleStatus()
	}
	task.ChangeTask(newTitle, description)

	if title != newTitle {
		delete(l.Tasks, title)
		l.Tasks[newTitle] = task
	}
	return nil
}
func (l *List) Delete(title string) error {
	_, ok := l.Tasks[title]
	if !ok {
		return ErrorTask{
			Message: "Task not Found",
			Type:    ErrNotFound,
		}
	}
	delete(l.Tasks, title)
	return nil
}
