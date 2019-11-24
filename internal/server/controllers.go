package server

import "github.com/gin-gonic/gin"

type MinesweeperController interface {
	PostNewGame(c *gin.Context)
	PostPlayGame(c *gin.Context)
	PostMarkGame(c *gin.Context)
	GetGameStatus(c *gin.Context)
}
