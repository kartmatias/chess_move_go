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
	serveMux.Handle("/knight", &controller.KnightHandler{})
	serveMux.Handle("/knight/", &controller.KnightHandler{})
	http.ListenAndServe(":"+port, serveMux)

}
