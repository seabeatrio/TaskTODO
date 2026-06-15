package ToDoList

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	Title       string
	Description string
	IsDone      bool

	StartTime time.Time
	EndTime   *time.Time
}

func NewTask(title string, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		IsDone:      false,

		StartTime: time.Now(),
		EndTime:   nil,
	}
}

func (t *Task) ToggleStatus() {
	t.IsDone = !t.IsDone
}
func (t *Task) WhatTime() {
	t.EndTime = new(time.Time)
	*t.EndTime = time.Now()
}
func (t *Task) ChangeTask(title string, description string) {
	t.Description = description
	t.Title = title
}
func (t *Task) Println() {
	_, err := pp.Print(t)
	if err != nil {
		fmt.Println("ERROR: ", "pp library not detection")
	}
}
func (t *Task) Get() *Task {
	return t
}
