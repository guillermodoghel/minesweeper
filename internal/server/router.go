package server

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/guillermodoghel/minesweeper-API/pkg/rest"
	"net/http"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(mswpCrtrl MinesweeperController) *Router {
	router := gin.New()

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Use(RecoverFromPanics())
	router.NoRoute(noRouteHandler)

	router.GET("/ping", func(context *gin.Context) {
		context.Header("x-mphome-ping", "this.is.a.genuine.pong.response")
		context.String(http.StatusOK, "pong")
	})

	minesweeperGroup := router.Group("/minesweeper")
	minesweeperGroup.Use(
		LogRequestAndResponse(),
	)
	minesweeperGroup.POST("/init", mswpCrtrl.PostNewGame)
	minesweeperGroup.POST("/play", mswpCrtrl.PostPlayGame)
	minesweeperGroup.GET("/status", mswpCrtrl.GetGameStatus)
	minesweeperGroup.POST("/mark", mswpCrtrl.PostMarkGame)

	return &Router{
		Engine: router,
	}
}

func (r Router) Run(addr string) error {
	return r.Engine.Run(addr)
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, rest.NewNotFoundApiError(fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)))
}
