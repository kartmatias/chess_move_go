package main

import (
	"github.com/kartmatias/chess_move_go/controller"
	"github.com/kartmatias/chess_move_go/model"
	"net/http"
	"os"
	"sync"
)

func main() {

	port := os.Getenv("CHESS_PORT")
	if port == "" {
		port = "8080"
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/home", &controller.HomeHandler{})
	serveMux.Handle("/knight", &controller.KnightHandler{})
	serveMux.Handle("/knight/", &controller.KnightHandler{
		Store:&model.PositionStore{
			List: map[string]model.Position{
			},
			RWMutex: &sync.RWMutex{},
		},
	})
	http.ListenAndServe(":"+port, serveMux)

}
