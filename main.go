package main

import (
	"github.com/kartmatias/chess_move_go/controller"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("CHESS_PORT")
	if port == "" {
		port = "8080"
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/home", &controller.HomeHandler{})
	http.ListenAndServe(":"+port, serveMux)

}
