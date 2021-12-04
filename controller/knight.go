package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kartmatias/chess_move_go/model"
	"net/http"
	"regexp"
)

type KnightHandler struct {
}

var (
	//https://regex101.com/
	getRequest = regexp.MustCompile(`^\/knight\/[abcdefgh](\d)$`)
)


func (h *KnightHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	switch {
	case request.Method == http.MethodGet :
		h.Get(writer, request)
	case request.Method == http.MethodPost :
		h.Post(writer, request)

	}
}

func (h *KnightHandler) Get(w http.ResponseWriter, r *http.Request)  {
	matches := getRequest.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "OK: %v\n", matches)
}

func (h *KnightHandler) Post(w http.ResponseWriter, r *http.Request)  {
	var p model.Position

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnError(w, http.StatusInternalServerError, "EMPTY BODY")
	}

}

