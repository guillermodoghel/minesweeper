package api

import (
	"fmt"
	"github.com/guillermodoghel/minesweeper-API/pkg/rest"
	"github.com/pborman/uuid"
	"math/rand"
	"time"
)

type MinesweeperService interface {
	InitGame(rows, cols, mines int) (*Game, rest.ApiError)
	PlayGame(gameId string, token string, col int, row int) (*Game, rest.ApiError)
	MarkCell(gameId string, token string, col int, row int, mark string) (*Game, rest.ApiError)
	GameStatus(gameId string, token string) (*Game, rest.ApiError)
}

type Cell struct {
	HasMine      bool   `json:"-"`
	Clicked      bool   `json:"clicked"`
	MinesArround int    `json:"-"`
	Label        string `json:"label"`
	X            int    `json:"X"`
	Y            int    `json:"Y"`
}

type Game struct {
	GameId             string    `json:"game_id"`
	Token              string    `json:"token,omitempty"`
	HumanFriendlyBoard string    `json:"human_friendly_board"`
	Board              [][]*Cell `json:"board"`
	GameSetup          GameSetup `json:"game_setup,omitempty"`
	Clics              int       `json:"-"`
	Status             string    `json:"game_status"`
	StartTimestamp     int64     `json:"start_timestamp"`
	ElapsedTimestamp   int64     `json:"elapsed_time"`
}
type GameSetup struct {
	Cols  int `json:"cols"`
	Rows  int `json:"rows"`
	Mines int `json:"mines"`
}

const (
	statusActive = "ACTIVE"
	statusLost   = "LOST"
	statusWon    = "WON"
)

type Point struct {
	X int
	Y int
}

func (g *Game) Refresh() {
	g.ElapsedTimestamp = time.Since(time.Unix(g.StartTimestamp, 0)).Milliseconds()
	//	g.Token = uuid.New()
}
func (g *Game) Mark(x, y int, mark string) rest.ApiError {
	if g.Status != statusActive {
		return rest.NewBadRequestApiError("game already " + g.Status)
	}
	if !isOnMap(x, y, g) {
		return rest.NewBadRequestApiError("invalid point dude")
	}
	playedCell := g.Board[y][x]
	if mark == "" {
		mark = " "
	}
	playedCell.Label = fmt.Sprintf("[%s]", mark);
	g.updateBoard()
	return nil
}
func (g *Game) Play(x, y int) rest.ApiError {
	if g.Status != statusActive {
		return rest.NewBadRequestApiError("game already " + g.Status)
	}
	if !isOnMap(x, y, g) {
		return rest.NewBadRequestApiError("invalid point dude")
	}
	playedCell := g.Board[y][x]
	if playedCell.Clicked {
		return rest.NewBadRequestApiError("already played")
	}
	if playedCell.HasMine {
		playedCell.Label = "[*]"
		playedCell.Clicked = true
		g.Status = statusLost
	} else {
		playedCell.Label = fmt.Sprintf("[%d]", playedCell.MinesArround);
		playedCell.Clicked = true
		if playedCell.MinesArround == 0 {
			zeroMinesArroundReveal(x, y, g)
		}
		g.Clics++
		if (g.GameSetup.Cols*g.GameSetup.Rows)-g.GameSetup.Mines <= g.Clics {
			g.Status = statusWon
		}
	}
	g.updateBoard()
	return nil
}
func zeroMinesArroundReveal(x, y int, g *Game) {
	points := getNeighborsOfPoint(x, y, g)
	for _, point := range points {
		if !g.Board[point.Y][point.X].Clicked {
			g.Board[point.Y][point.X].Label = fmt.Sprintf("[%d]", g.Board[point.Y][point.X].MinesArround);
			g.Board[point.Y][point.X].Clicked = true
			g.Clics++
			if g.Board[point.Y][point.X].MinesArround == 0 {
				zeroMinesArroundReveal(point.X, point.Y, g)
			}
		}
	}
}

func (g *Game) NewGame(rows, cols, mines int) {
	g.GameSetup = GameSetup{cols, rows, mines}

	g.StartTimestamp = time.Now().Unix()
	var i, j, k int
	k = 0
	//Initialize the board
	g.Board = make([][]*Cell, rows)
	for i = 0; i < rows; i++ {
		g.Board[i] = make([]*Cell, cols)
		for j = 0; j < cols; j++ {
			if g.Board[i][j] == nil {
				g.Board[i][j] = &Cell{HasMine: false, MinesArround: 0, Label: "[ ]", X: j, Y: i}
			}
		}
	}
	//Set random mines
	rand.Seed(time.Now().UnixNano())
	for k < mines {
		i := rand.Intn(rows)
		j := rand.Intn(cols)
		if !g.Board[i][j].HasMine {
			g.Board[i][j] = &Cell{HasMine: true, Clicked: false, Label: "[ ]", X: j, Y: i}
			k++
		}
	}
	//Set values
	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			if g.Board[i][j].HasMine {
				points := getNeighborsOfPoint(i, j, g)
				for _, point := range points {
					if !g.Board[point.Y][point.X].HasMine {
						g.Board[point.Y][point.X].MinesArround++
					}
				}
			}
		}
	}
	g.Status = statusActive
	g.updateBoard()
	g.GameId = uuid.New()
	g.Token = uuid.New()
	g.Clics = 0
}

func getNeighborsOfPoint(x, y int, game *Game) []Point {
	var points []Point
	for xx := -1; xx <= 1; xx++ {
		for yy := -1; yy <= 1; yy++ {
			if (xx == 0 && yy == 0) {
				continue;
			}
			if (isOnMap(x+xx, y+yy, game)) {
				points = append(points, Point{x + xx, y + yy})
			}
		}
	}
	return points;
}

func isOnMap(x, y int, game *Game) bool {
	return x >= 0 && y >= 0 && x < game.GameSetup.Cols && y < game.GameSetup.Rows;
}

func (game *Game) updateBoard() {
	var i, j int
	s := ""
	for j = 0; j < game.GameSetup.Rows; j++ {
		for i = 0; i < game.GameSetup.Cols; i++ {
			if game.Status == statusActive {
				s += fmt.Sprint(game.Board[j][i].Label)
			} else {
				if game.Board[j][i].HasMine {
					s += fmt.Sprint("[*]")
					game.Board[j][i].Label = "[*]"
				} else {
					s += fmt.Sprintf("[%d]", game.Board[j][i].MinesArround)
					game.Board[j][i].Label = fmt.Sprintf("[%d]", game.Board[j][i].MinesArround)
				}

			}
		}
		s += fmt.Sprint("\n")
	}
	game.HumanFriendlyBoard = s
}
