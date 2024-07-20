package service

import (
	"net/http"

	_ "todo-list/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Todo List API
//	@version		1.0
//	@description	This is a simple Todo List API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@BasePath
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
