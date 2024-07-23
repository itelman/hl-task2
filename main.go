package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"todo-list/internal/config"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port for api")
	flag.Parse()

	fmt.Printf("Server starting on http://localhost:%d\n\n", port)
	err := http.ListenAndServe(":8080", config.Routes())
	if err != nil {
		log.Fatal(err, nil)
	}
}
