package http

import (
	"awesomeLearnigGO/ToDoList"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *ToDoList.List
}

func ResponseWithError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	// Used errors.As - this is modern standard GO
	// He is work, even if the error is passed as a value OR as a pointer!

	var customErr ToDoList.ErrorTask
	if errors.As(err, &customErr) {
		w.WriteHeader(customErr.StatusCode()) // auto fetch error status code
		json.NewEncoder(w).Encode(customErr)  // ToString() used standard encoder
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
}
func NewHTTPHandlers(todoList *ToDoList.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP request body

succeed:
  - status code : 201 Created
  - response body: JSON represent created task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with Error + time
*/

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		err := ToDoList.ErrorTask{
			Message: "[ERROR]: Validation",
			Type:    ToDoList.ErrValidation,
		}
		ResponseWithError(w, err)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		ResponseWithError(w, err)
		return
	}
	if err := h.todoList.CreateTask(taskDTO.Title, taskDTO.Description); err != nil {
		ResponseWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskDTO)
}

/*
pattern: /tasks/{title}
method: GET
info: pattern

succeed:

 - status code: 400, 404, 500, ...
 - response body: JSON with Error + time

*/

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pathTitle := vars["title"]
	//pathTitle := r.PathValue("title")
	task, err := h.todoList.GetTask(pathTitle)
	if err != nil {
		ResponseWithError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

/*
pattern: /tasks
method: GET
info: -
succeed:
 - status code: 200 Ok
 - response body: JSON represented found tasks
failed:
 - status code: 400, 500, ...
 - response body: JSON with Error + time
*/

func (h HTTPHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	if r.URL.RawQuery != "" && title == "" {
		err := ToDoList.ErrorTask{
			Message: "Invalid query arguments",
			Type:    ToDoList.ErrNotFound,
		}
		ResponseWithError(w, err)
		return
	}

	filter := ToDoList.TaskFilter{
		Title:     title,
		Completed: nil,
	}
	tasks := h.todoList.GetTasks(filter)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)

}

/*
pattern : /tasks/{title}?compeleted=true
method: PUTCH
info: pattern + JSON in Request body

succeed:
 - status code: 200 Ok
 - response body: JSON represented change tasks

failed:
 - status code: 400, 409,  500, ...
 - response body: JSON with Error + time

*/

func (h *HTTPHandlers) HandleChangeTask(w http.ResponseWriter, r *http.Request) {
	//pathTitle := r.PathValue("title")
	vars := mux.Vars(r)
	pathTitle := vars["title"]

	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		err := ToDoList.ErrorTask{
			Message: "[ERROR]: Validation",
			Type:    ToDoList.ErrValidation,
		}
		ResponseWithError(w, err)
		return
	}

	if err := taskDTO.ValidateForUpdate(); err != nil {
		ResponseWithError(w, err)
		return
	}

	if err := h.todoList.ChangeTask(
		pathTitle,
		taskDTO.Title,
		taskDTO.Description,
		taskDTO.Completed); err != nil {
		ResponseWithError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskDTO)

}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
 - status code: 204 No content
 - response body: -
failed:
 - status code: 400, 404, 409, ...
 - response body: JSON with Error + time
*/

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	//pathTitle := r.PathValue("title")
	vars := mux.Vars(r)
	pathTitle := vars["title"]

	if err := h.todoList.Delete(pathTitle); err != nil {
		ResponseWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
