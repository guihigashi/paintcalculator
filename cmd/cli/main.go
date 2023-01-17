package main

import (
	"log"

	"drcodechallenge/pkg/domain"

	"github.com/samber/lo"
)

func main() {
	r, err := domain.NewRectangle(2, 2)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("a valid rectangle: %#v\n", *r)

	w, err := domain.NewWall(5, 5, 1, 1)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("a valid wall: %#v\n", *w)

	walls := lo.Times(4, func(index int) domain.Wall {
		w, _ := domain.NewWall(5, 5, 1, 1)
		return *w
	})
	room, err := domain.NewRoom(walls)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("a valid room: %#v\n", *room)
}
