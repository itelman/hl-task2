package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"todo-list/internal/service/helpers"
	"todo-list/internal/service/timego"
	"todo-list/pkg/models"

	"github.com/gorilla/mux"
)

var store sync.Map

// createTaskHandler godoc
//
//	@Summary		Create a new task
//	@Description	Create a new task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		models.TaskRequest	true	"Task"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		404		{string}	string	"Not Found"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/todo-list/tasks [post]
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskReq models.TaskRequest

	if err := helpers.ReadJSON(w, r, &taskReq); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	task := models.NewTask(taskReq)

	if _, ok := store.Load(task.ID); ok {
		NotFoundResponse(w, r)
	}

	store.Store(task.ID, task)

	err := helpers.WriteJSON(w, http.StatusCreated, map[string]interface{}{"id": task.ID}, nil)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}
}

// updateTaskHandler godoc
//
//	@Summary		Update a task
//	@Description	Update a task by ID
//	@Tags			tasks
//	@Accept			json
//	@Param			id		path	string				true	"Task ID"
//	@Param			task	body	models.TaskRequest	true	"Task"
//	@Success		204
//	@Failure		400	{string}	string	"Bad Request"
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id} [put]
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var taskReq models.TaskRequest

	if _, ok := store.Load(id); !ok {
		NotFoundResponse(w, r)
		return
	}

	if err := helpers.ReadJSON(w, r, &taskReq); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	task := models.UpdatedTask(taskReq, id)

	store.Swap(task.ID, task)

	w.WriteHeader(http.StatusNoContent)
}

// deleteTaskHandler godoc
//
//	@Summary		Delete a task
//	@Description	Delete a task by ID
//	@Tags			tasks
//	@Param			id	path	string	true	"Task ID"
//	@Success		204
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id} [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if _, ok := store.Load(id); !ok {
		NotFoundResponse(w, r)
		return
	}

	store.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

// checkTaskHandler godoc
//
//	@Summary		Mark task as done
//	@Description	Update a task status by ID
//	@Tags			tasks
//	@Param			id	path	string	true	"Task ID"
//	@Success		204
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id}/done [put]
func CheckTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	val, ok := store.Load(id)
	if !ok {
		NotFoundResponse(w, r)
		return
	}

	task := val.(models.Task)
	task.Check()

	store.Swap(task.ID, task)

	w.WriteHeader(http.StatusNoContent)
}

// listTaskHandler godoc
//
//	@Summary		Get all tasks
//	@Description	Get all tasks by status
//	@Tags			tasks
//	@Produce		json
//	@Param			status	query		string	false	"Status Filter"	Enum(active,done)
//	@Success		200		{array}		models.Task
//	@Failure		404		{string}	string	"Not Found"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/todo-list/tasks [get]
func ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if len(status) == 0 {
		status = "active"
	}

	if !(status == "active" || status == "done") {
		NotFoundResponse(w, r)
		return
	}

	var tasks, response []*models.Task
	var arr []interface{}

	f := func(key, value any) bool {
		arr = append(arr, value)

		return true
	}

	store.Range(f)

	for _, val := range arr {
		task := val.(models.Task)
		tasks = append(tasks, &task)
	}

	for _, task := range tasks {
		isWeekend, err := timego.IsWeekend(task.ActiveAt)
		if err != nil {
			ServerErrorResponse(w, r, err)
			return
		}

		if isWeekend {
			task.MarkAsWeekend()
		}
	}

	if status == "active" {
		for _, task := range tasks {
			isActive, err := timego.IsActive(task.ActiveAt)
			if err != nil {
				ServerErrorResponse(w, r, err)
				return
			}

			if isActive {
				response = append(response, task)
			}
		}
	}

	if status == "done" {
		for _, task := range tasks {
			if task.Done {
				response = append(response, task)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ShowTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	val, ok := store.Load(id)
	if !ok {
		NotFoundResponse(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}
