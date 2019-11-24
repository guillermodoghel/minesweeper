package main

import (
	"github.com/guillermodoghel/minesweeper/internal/api"
	"github.com/guillermodoghel/minesweeper/internal/api/minesweeper"
	"github.com/guillermodoghel/minesweeper/internal/api/minesweeper/local_kvs"
)

func setupMinesweeperService() api.MinesweeperService {
	return minesweeper.NewMinesweeperService(local_kvs.NewPersistenceService());
}
