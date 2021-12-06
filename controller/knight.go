package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kartmatias/chess_move_go/model"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type KnightHandler struct {
	Store *model.PositionStore
}

type coordinate struct {
	x int
	y int
}

var (
	//https://regex101.com/
	getRequest = regexp.MustCompile(`^\/knight\/[a-h][1-8]$`)
	listRequest = regexp.MustCompile(`\/knight[\/]*$`)
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

	position := strings.ReplaceAll(getRequest.FindString(r.URL.Path), "/knight/", "")

	if !IsValidPosition(position) {
		ReturnError(w, http.StatusInternalServerError, "INVALID POSITION")
		return
	}

	fmt.Println(position)
	h.generate()

	h.Store.RLock()
	positions := make([]model.Position, 0, len(h.Store.List))

	for _, value := range h.Store.List {
		if isValidMove(position, value.ID) {
			positions = append(positions, value)
		}
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

func (h *KnightHandler) Post(w http.ResponseWriter, r *http.Request)  {
	var p model.Position

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnError(w, http.StatusInternalServerError, "EMPTY BODY")
	}

	if !IsValidPosition(p.ID) {
		ReturnError(w, http.StatusInternalServerError, "INVALID POSITION")
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

	if len(h.Store.List) != 0 {
		return
	}

	list := GenerateValidPositions()
	for _, position := range list {
		h.Store.Lock()
		h.Store.List[position.ID] = position
		h.Store.Unlock()
	}
}

func GetX(pos string) int {
	return int(CharCodeAt(pos, 0) - 96)
}

func GetY(pos string) int {
	y, err := strconv.Atoi( Substr(pos,1,1) )
	if err != nil {
		return 0
	}
	return y
}

func IsValidPosition(position string) bool {
	var validPos = regexp.MustCompile(`(?m)[a-h][1-8]$`)
	matches := validPos.FindStringSubmatch(position)
	if len(matches) < 1 {
		return false
	}
	return true
}

func (h *KnightHandler) ValidMoves(w http.ResponseWriter, r *http.Request) {

	h.Store.RLock()
	positions := make([]model.Position, 0, len(h.Store.List))

	for _, value := range h.Store.List {
		if isValidMove("a1", "a2") {
			positions = append(positions, value)
		}
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

func isValidMove(origin string, destination string) bool {
	var p1 = coordinate{x: GetX(origin), y: GetY(origin)}
	var p2 = coordinate{x: GetX(destination), y: GetY(destination)}

	return (math.Abs(float64(p2.x - p1.x)) == 1) &&
		(math.Abs(float64(p2.y - p1.y)) == 2) ||
		(math.Abs(float64(p2.x - p1.x)) == 2) &&
			(math.Abs(float64(p2.y - p1.y)) == 1)

}

func GenerateValidPositions() []model.Position {
	positions := make([]model.Position, 0, 63)
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			strPos := string(96 + i) + strconv.Itoa(j)
			positions = append(positions,model.Position{ID: strPos, Position: strPos})
		}
	}
	return positions
}
