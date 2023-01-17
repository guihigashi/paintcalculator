package domain

import (
	"errors"

	"github.com/samber/lo"
)

type Room struct {
	Walls []Wall
}

func NewRoom(walls []Wall) (*Room, error) {
	if len(walls) != 4 {
		return nil, errors.New("a sala deve ser composta por 4 paredes")
	}

	return &Room{walls}, nil
}

func (r Room) PaintableArea() float64 {
	return lo.Reduce(r.Walls,
		func(acc float64, cur Wall, _ int) float64 {
			return acc + cur.Area()
		}, 0)
}
