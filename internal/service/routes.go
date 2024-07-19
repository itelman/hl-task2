package service

import (
	"net/http"

	_ "proxy-server/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			HTTP Proxy Server API
// @version		1.0
// @description	This is a server to proxy HTTP requests.
// @host			localhost:8080
// @BasePath		/
func (app *Application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/todo-list/tasks", app.createTaskHandler).Methods("POST")
	router.HandleFunc("/api/todo-list/tasks/{id}", app.updateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/todo-list/tasks/{id}", app.deleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/todo-list/tasks/{id}/done", app.checkTaskHandler).Methods("PUT")
	router.HandleFunc("/api/todo-list/tasks", app.listTaskHandler).Methods("GET")

	router.HandleFunc("/health", app.healthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	return router
}
