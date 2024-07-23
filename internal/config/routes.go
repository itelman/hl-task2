package config

import (
	"net/http"
	"todo-list/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "todo-list/docs"
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

// @BasePath
func Routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our app receives.
	standardMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)

	router := mux.NewRouter()

	router.HandleFunc("/api/todo-list/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/todo-list/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/todo-list/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/todo-list/tasks/{id}/done", handlers.CheckTaskHandler).Methods("PUT")
	router.HandleFunc("/api/todo-list/tasks", handlers.ListTaskHandler).Methods("GET")

	router.HandleFunc("/api/todo-list/tasks/{id}", handlers.ShowTaskHandler).Methods("GET")

	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedResponse)

	return standardMiddleware.Then(router)
}
