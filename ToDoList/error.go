package ToDoList

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrNotFound     ErrorType = "NOT_FOUND"
	ErrValidation   ErrorType = "VALIDATION_FAILED"
	ErrInternal     ErrorType = "INTERNAL"
	ErrExistAlready ErrorType = "EXIST_ALREADY"
)

type ErrorTask struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

func (e ErrorTask) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}
func (e ErrorTask) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}
func (e ErrorTask) StatusCode() int {
	switch e.Type {
	case ErrValidation:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	case ErrExistAlready:
		return http.StatusConflict
	case ErrInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
