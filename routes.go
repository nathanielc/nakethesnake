package main

import (
	"log"
	"math"
	"net/http"

	"github.com/nathanielc/nakethesnake/api"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	//dump(decoded)

	respond(res, api.StartResponse{
		Color: "#A5F3B4",
	})
}

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"
)

var moves = []string{
	UP,
	DOWN,
	LEFT,
	RIGHT,
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	//dump(decoded)

	move := RIGHT
	foods := decoded.Board.Food
	if len(foods) > 0 {
		head := decoded.You.Body[0]
		move = findFood(head, foods)
	}
	move = findSafe(&decoded, move)

	respond(res, api.MoveResponse{
		Move: move,
	})
}

func distance(a, b api.Coord) int {
	return b.X - a.X + b.Y - a.Y
}

func findFood(head api.Coord, foods []api.Coord) (move string) {
	var minIdx, min int
	min = math.MaxInt32
	for i, f := range foods {
		d := distance(head, f)
		if d < min {
			minIdx = i
			min = d
		}
	}

	food := foods[minIdx]

	dx := food.X - head.X
	dy := food.Y - head.Y

	if dx > dy {
		if dx > 0 {
			move = RIGHT
		} else {
			move = LEFT
		}
	} else {
		if dy > 0 {
			move = DOWN
		} else {
			move = UP
		}
	}
	return
}
func findSafe(game *api.SnakeRequest, move string) string {
	fm := makeFlat(game)
	head := game.You.Body[0]
	pos := movePos(head, move)
	if isSafe(fm, pos) {
		return move
	}

	for _, m := range moves {
		pos := movePos(head, m)
		if isSafe(fm, pos) {
			return m
		}
	}
	return move
}

func movePos(pos api.Coord, move string) api.Coord {
	switch move {
	case UP:
		pos.Y -= 1
	case DOWN:
		pos.Y += 1
	case LEFT:
		pos.X -= 1
	case RIGHT:
		pos.X += 1
	}
	return pos
}

func isSafe(fm flatMap, pos api.Coord) bool {
	return !fm.At(pos.X, pos.Y)
}

type flatMap struct {
	w, h int
	data []bool
}

func (m flatMap) At(x, y int) bool {
	if x < 0 || y < 0 || x >= m.w || y >= m.h {
		return true
	}
	return m.data[x+m.w*y]
}

func makeFlat(game *api.SnakeRequest) flatMap {
	m := flatMap{
		w:    game.Board.Width,
		h:    game.Board.Height,
		data: make([]bool, game.Board.Width*game.Board.Height),
	}
	for _, s := range game.Board.Snakes {
		for _, coord := range s.Body {
			i := coord.X + m.w*coord.Y
			m.data[i] = true
		}
	}
	return m
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
