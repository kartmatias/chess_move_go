package model

import "sync"

type Position struct {
	ID string `json:"id"`
	Position string `json:"position"`
}

type PositionStore struct {
	m map[string]Position
	*sync.RWMutex
}
