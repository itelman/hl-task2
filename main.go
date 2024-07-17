package main

import (
	"flag"
	"fmt"
	"os"
	"proxy-server/internal/service"
	"proxy-server/internal/service/jsonlog"
)

const port = 8080

func main() {
	var cfg service.Config

	flag.IntVar(&cfg.Port, "port", port, "port for api")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &service.Application{
		Config: cfg,
		Logger: logger,
	}

	fmt.Printf("Server starting on http://localhost:%d\n\n", port)
	err := app.Serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
