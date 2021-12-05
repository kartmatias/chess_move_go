package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kartmatias/chess_move_go/model"
	"github.com/kartmatias/chess_move_go/service"
	"net/http"
	"regexp"
)

type KnightHandler struct {
	Store *model.PositionStore
}

var (
	//https://regex101.com/
	getRequest = regexp.MustCompile(`^\/knight\/[a-h][1-8]$`)
	listRequest = regexp.MustCompile(`\/knight[\/]*$`)
	validPos = regexp.MustCompile(`(?m)[a-h][1-8]$`)
)


func (h *KnightHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	switch {
	case request.Method == http.MethodGet && getRequest.MatchString(request.URL.Path):
		h.Get(writer, request)
	case request.Method == http.MethodGet && listRequest.MatchString(request.URL.Path):
		h.List(writer, request)
	case request.Method == http.MethodPost :
		h.Post(writer, request)
	}

}

func (h *KnightHandler) Get(w http.ResponseWriter, r *http.Request)  {
	matches := getRequest.FindStringSubmatch(r.URL.Path)

	for i, match := range getRequest.FindAllString(r.URL.Path, -1) {
		fmt.Println(match, "found at index", i)
	}

	if len(matches) < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "GERANDO MOVIMENTOS....: %v\n", matches)

	h.generate()

	fmt.Fprintf(w, "OK: %v\n", matches)
}

func (h *KnightHandler) Post(w http.ResponseWriter, r *http.Request)  {
	var p model.Position

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnError(w, http.StatusInternalServerError, "EMPTY BODY")
	}

	matches := validPos.FindStringSubmatch(p.ID)

	if len(matches) < 1 {
		http.NotFound(w, r)
		return
	}

	h.Store.Lock()
	h.Store.List[p.ID] = p
	h.Store.Unlock()

	jsonBytes, err := json.Marshal(p)

	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *KnightHandler) List(w http.ResponseWriter, r *http.Request) {
	h.Store.RLock()
	positions := make([]model.Position, 0, len(h.Store.List))

	for _, value := range h.Store.List {
		positions = append(positions, value)
	}

	h.Store.RUnlock()
	jbytes, err := json.Marshal(positions)

	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Error when listing")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jbytes)

}

func (h *KnightHandler) generate() {
	lista := service.GenerateValidPositions()

	for _, position := range lista {
		h.Store.Lock()
		h.Store.List[position.ID] = position
		h.Store.Unlock()
	}

}

