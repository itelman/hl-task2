package handlers

import (
	"net/http"
	"todo-list/internal/service/helpers"
)

// healthCheckHandler godoc
//
//	@Summary		Health check
//	@Description	This endpoint checks the health of the server.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"environment": "development",
			"version":     "1.0",
		},
	}

	err := helpers.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		ServerErrorResponse(w, r, err)
	}
}
