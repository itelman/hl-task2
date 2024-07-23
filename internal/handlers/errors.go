package handlers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"todo-list/internal/service/helpers"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := map[string]interface{}{"error": message}

	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	if err := log.Output(2, trace); err != nil {
		log.Println("issue with printing error logs", err)
	}

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}
