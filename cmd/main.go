package main

import (
	"github.com/guillermodoghel/minesweeper/internal/server"
	"github.com/guillermodoghel/minesweeper/internal/server/minesweeper"
	"github.com/guillermodoghel/minesweeper/pkg/log"
	"github.com/jpmrno/go-meli-toolkit/goutils/logger"
)

const serverAddress = ":8080"

func main() {
	initTheCoolLogger()
	log.SetLevel("info")
	log.Info("Starting reactor core")
	router := server.NewRouter(minesweeper.NewMinesweeperController(setupMinesweeperService()))
	log.Info("reactor core running at port " + serverAddress)
	log.ThrowPanicWithError("Server initialization error", router.Run(serverAddress))
}

func initTheCoolLogger() {
	logger.Log.SetFormatter(&log.MercuryFormatter{
	})
	logger.SetLogLevel("info")
}
