package service

import (
	"net/http"
	"proxy-server/pkg/models"
	"sync"
)

var Storage sync.Map

func (app *Application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskReq models.TaskRequest
	var task models.Task

	if err := app.readJSON(w, r, &taskReq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := task.Set(taskReq); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.Cache.Insert(task.ID, task)

	err := app.writeJSON(w, http.StatusCreated, envelope{"id": task.ID}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		app.notFoundResponse(w, r)
		return
	}

	var taskReq models.TaskRequest

	if err := app.readJSON(w, r, &taskReq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		app.notFoundResponse(w, r)
		return
	}

	app.Cache.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) checkTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		app.notFoundResponse(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) listTaskHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if len(status) == 0 {
		status = "active"
	}

	if !(status == "active" || status == "done") {
		app.notFoundResponse(w, r)
		return
	}

	// return code = http.StatusOK
}
