package http

import "awesomeLearnigGO/ToDoList"

// DTO == Data transfer object

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed,omitempty"`
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return ToDoList.ErrorTask{
			Type:    ToDoList.ErrValidation,
			Message: "Title is empty",
		}
	} else if t.Description == "" {
		return ToDoList.ErrorTask{
			Type:    ToDoList.ErrValidation,
			Message: "Description is empty",
		}
	}
	return nil
}

func (t TaskDTO) ValidateForUpdate() error {
	if t.Description == "" && t.Title == "" && t.Completed != nil {
		return ToDoList.ErrorTask{
			Type:    ToDoList.ErrValidation,
			Message: "Title is empty",
		}
	}
	return nil
}
