package minesweeper

import (
	"github.com/guillermodoghel/minesweeper/internal/api"
)

type Persistence struct {
}

type PersistenceService interface {
	WriteGame(game api.Game) error
	ReadGame(gameId string) (*api.Game, error)
}
