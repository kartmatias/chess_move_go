package service

import (
	"github.com/kartmatias/chess_move_go/model"
	"strconv"
)

func GetX(pos string) int {
	return int(CharCodeAt(pos, 0) - 96)
}

func GetY(pos string) (int, error) {
	return strconv.Atoi( Substr(pos,1,1) )
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
