package service

import "net/http"

// healthCheckHandler godoc
//
//	@Summary		Health check
//	@Description	This endpoint checks the health of the server.
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"OK"
//	@Router			/health [get]
func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
