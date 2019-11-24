package minesweeper

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermodoghel/minesweeper-API/internal/api"
	"github.com/guillermodoghel/minesweeper-API/internal/server"
	"github.com/guillermodoghel/minesweeper-API/pkg/rest"
	"net/http"
)

type PlayRequest struct {
	GameId string `json:"game_id"`
	Token  string `json:"token"`
	X      int    `json:"X"`
	Y      int    `json:"Y"`
	Mark   string `json:"mark"`
}

type Setup struct {
	Cols  int `json:"cols"`
	Rows  int `json:"rows"`
	Mines int `json:"mines"`
}

type minesweeperController struct {
	minesweeperService api.MinesweeperService
}

func NewMinesweeperController(minesweeperService api.MinesweeperService) server.MinesweeperController {
	return &minesweeperController{
		minesweeperService: minesweeperService,
	}
}
func (s *minesweeperController) PostMarkGame(c *gin.Context) {
	var playrequest PlayRequest
	err := c.Bind(&playrequest)
	if err != nil {
		apiError := rest.NewBadRequestApiError(err.Error())
		c.JSON(apiError.Status(), apiError)
		return;
	}
	if !((playrequest.Mark == "F") || (playrequest.Mark == "?") || (playrequest.Mark == "")) {
		apiError := rest.NewBadRequestApiError("Only 'F' or '?' or '' marks supported")
		c.JSON(apiError.Status(), apiError)
		return;
	}
	gameResult, apierr := s.minesweeperService.MarkCell(playrequest.GameId, playrequest.Token, playrequest.X, playrequest.Y, playrequest.Mark)
	if apierr != nil {
		c.JSON(apierr.Status(), apierr)
	} else {
		c.JSON(http.StatusOK, gameResult)
	}
}

func (s *minesweeperController) PostNewGame(c *gin.Context) {
	var gameSetup Setup
	err := c.Bind(&gameSetup)
	if err != nil {
		apiError := rest.NewBadRequestApiError(err.Error())
		c.JSON(apiError.Status(), apiError)
		return;
	}
	if gameSetup.Cols < 1 || gameSetup.Rows < 1 || gameSetup.Mines < 0 || gameSetup.Mines > gameSetup.Rows*gameSetup.Cols {
		apiError := rest.NewBadRequestApiError("invalid config")
		c.JSON(apiError.Status(), apiError)
		return;
	}
	gameResult, err := s.minesweeperService.InitGame(gameSetup.Rows,gameSetup.Cols, gameSetup.Mines)
	if err != nil {
		apiError := rest.NewInternalServerApiError("Get init a game failed", err)
		c.JSON(apiError.Status(), apiError)
	} else {
		c.JSON(http.StatusCreated, &gameResult)
	}
}

func (s *minesweeperController) PostPlayGame(c *gin.Context) {
	var playrequest PlayRequest
	err := c.Bind(&playrequest)
	if err != nil {
		apiError := rest.NewBadRequestApiError(err.Error())
		c.JSON(apiError.Status(), apiError)
		return;
	}
	gameResult, apierr := s.minesweeperService.PlayGame(playrequest.GameId, playrequest.Token, playrequest.X, playrequest.Y)
	if apierr != nil {
		c.JSON(apierr.Status(), apierr)
	} else {
		c.JSON(http.StatusOK, gameResult)
	}
}

func (s *minesweeperController) GetGameStatus(c *gin.Context) {

	gameId := c.Request.URL.Query().Get("game_id")
	if gameId == "" {
		apiError := rest.NewBadRequestApiError("missing game_id")
		c.JSON(apiError.Status(), apiError)
		return
	}
	token := c.Request.URL.Query().Get("token")
	if token == "" {
		apiError := rest.NewBadRequestApiError("missing token")
		c.JSON(apiError.Status(), apiError)
		return
	}

	gameResult, err := s.minesweeperService.GameStatus(gameId, token)
	if err != nil {
		c.JSON(err.Status(), err)
	} else {
		c.JSON(http.StatusOK, gameResult)
	}
}
