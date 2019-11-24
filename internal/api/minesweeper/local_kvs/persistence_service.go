package local_kvs

import (
	"github.com/guillermodoghel/minesweeper/internal/api"
	"github.com/guillermodoghel/minesweeper/internal/api/minesweeper"
	"github.com/guillermodoghel/skv"
)

type persistenceService struct {
}

const db_path = "games.db"

func NewPersistenceService() minesweeper.PersistenceService {
	return &persistenceService{}
}

func (s *persistenceService) WriteGame(game api.Game) error {
	store, err := skv.Open(db_path)
	if err != nil {
		return err
	}
	store.Put(game.GameId, game);
	store.Close()
	return nil
}
func (s *persistenceService) ReadGame(gameId string) (*api.Game, error) {
	store, err := skv.Open(db_path)
	if err != nil {
		return nil, err
	}
	var game = api.Game{}
	err = store.Get(gameId, &game)
	store.Close()
	if err != nil {
		return nil, err
	}
	return &game, nil
}
