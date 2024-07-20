package service

import (
	"encoding/json"
	"net/http"
	"sync"
	"todo-list/internal/service/timego"
	"todo-list/pkg/models"
)

var Storage sync.Map

// createTaskHandler godoc
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
func (app *Application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskReq models.TaskRequest

	if err := app.readJSON(w, r, &taskReq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	task := models.NewTask(taskReq)

	if err := app.Cache.Insert(task.ID, task); err != nil {
		app.notFoundResponse(w, r)
	}

	err := app.writeJSON(w, http.StatusCreated, envelope{"id": task.ID}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// updateTaskHandler godoc
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
func (app *Application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var taskReq models.TaskRequest

	if err := app.readJSON(w, r, &taskReq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	task := models.UpdatedTask(taskReq, id)

	if err := app.Cache.Update(task.ID, task); err != nil {
		app.notFoundResponse(w, r)
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteTaskHandler godoc
//	@Summary		Delete a task
//	@Description	Delete a task by ID
//	@Tags			tasks
//	@Param			id	path	string	true	"Task ID"
//	@Success		204
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id} [delete]
func (app *Application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := app.Cache.Delete(id); err != nil {
		app.notFoundResponse(w, r)
	}

	w.WriteHeader(http.StatusNoContent)
}

// checkTaskHandler godoc
//	@Summary		Mark task as done
//	@Description	Update a task status by ID
//	@Tags			tasks
//	@Param			id	path	string	true	"Task ID"
//	@Success		204
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id}/done [put]
func (app *Application) checkTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	val, err := app.Cache.Get(id)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	task := val.(models.Task)
	task.Check()

	if err := app.Cache.Update(task.ID, task); err != nil {
		app.notFoundResponse(w, r)
	}

	w.WriteHeader(http.StatusNoContent)
}

// listTaskHandler godoc
//	@Summary		Get all tasks
//	@Description	Get all tasks by status
//	@Tags			tasks
//	@Produce		json
//	@Param			status	query		string	false	"Status Filter"	Enum(active,done)
//	@Success		200		{array}		models.Task
//	@Failure		404		{string}	string	"Not Found"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/todo-list/tasks [get]
func (app *Application) listTaskHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if len(status) == 0 {
		status = "active"
	}

	if !(status == "active" || status == "done") {
		app.notFoundResponse(w, r)
		return
	}

	var tasks, response []*models.Task

	for _, val := range app.Cache.GetAll() {
		tasks = append(tasks, val.(*models.Task))
	}

	for _, task := range tasks {
		isWeekend, err := timego.IsWeekend(task.ActiveAt)
		if err != nil {
			app.serverErrorResponse(w, r, err)
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
				app.serverErrorResponse(w, r, err)
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
