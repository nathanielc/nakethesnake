package main

import (
	"log"
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
		food := foods[0]
		head := decoded.You.Body[0]
		dx := food.X - head.X
		dy := food.Y - head.Y
		log.Println("food", food)
		log.Println("head", head)

		if dx > dy {
			if dx > 0 {
				move = LEFT
			} else {
				move = RIGHT
			}
		} else {
			if dy > 0 {
				move = DOWN
			} else {
				move = UP
			}
		}
	}

	respond(res, api.MoveResponse{
		Move: move,
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
