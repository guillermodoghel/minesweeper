package minesweeper

import (
	"github.com/guillermodoghel/minesweeper-API/internal/api"
	"github.com/guillermodoghel/minesweeper-API/pkg/rest"
)

type minesweeperService struct {
	persistenceService PersistenceService
}

func NewMinesweeperService(persistenceService PersistenceService) api.MinesweeperService {
	return &minesweeperService{persistenceService}
}

func (s *minesweeperService) InitGame(cols, rows, mines int) (*api.Game, rest.ApiError) {
	var g api.Game
	g.NewGame(cols, rows, mines)

	err := s.persistenceService.WriteGame(g);
	if err != nil {
		return nil, rest.NewInternalServerApiError("whoops", err);
	}
	return &g, nil
}

func (s *minesweeperService) MarkCell(gameId string, token string, col int, row int, mark string) (*api.Game, rest.ApiError) {
	game, apierr := s.getGame(gameId, token)
	if apierr != nil {
		return nil, apierr
	}
	apierr = game.Mark(row, col, mark)
	if apierr != nil {
		return nil, apierr
	}
	game.Refresh()
	err := s.persistenceService.WriteGame(*game);
	if err != nil {
		return nil, rest.NewInternalServerApiError("whoops", err);
	}
	return game, nil;
}

func (s *minesweeperService) PlayGame(gameId string, token string, row int, col int) (*api.Game, rest.ApiError) {
	game, apierr := s.getGame(gameId, token)
	if apierr != nil {
		return nil, apierr
	}
	apierr = game.Play(row, col)
	if apierr != nil {
		return nil, apierr
	}
	game.Refresh()
	err := s.persistenceService.WriteGame(*game);
	if err != nil {
		return nil, rest.NewInternalServerApiError("whoops", err);
	}
	return game, nil;
}

func (s *minesweeperService) GameStatus(gameId string, token string) (*api.Game, rest.ApiError) {
	return s.getGame(gameId, token)
}

func (s *minesweeperService) getGame(gameId string, token string) (*api.Game, rest.ApiError) {
	game, err := s.persistenceService.ReadGame(gameId);
	if err != nil {
		if err.Error() == "skv: key not found" {
			return nil, rest.NewNotFoundApiError("game not found")
		}
		return nil, rest.NewInternalServerApiError("whoops", err);
	}
	if game.Token != token {
		return nil, rest.NewUnauthorizedApiError("wrong token")
	}
	return game, nil
}
