package main

import (
	"drcodechallenge/pkg/domain"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Wall struct {
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	QtyWindows int     `json:"qtyWindows"`
	QtyDoors   int     `json:"qtyDoors"`
}

type CansNeededRequest struct {
	Room []Wall `json:"room"`
}

func (cn *CansNeededRequest) Bind(r *http.Request) error {
	if cn.Room == nil {
		return errors.New("room é obrigatório")
	}
	return nil
}

type CanLine struct {
	Id    uuid.UUID `json:"id"`
	Label string    `json:"label"`
	Qty   int       `json:"qty"`
}
type CansNeededResponse struct {
	Id         uuid.UUID `json:"id"`
	Target     float64   `json:"target"`
	SolvedFor  float64   `json:"solved_for"`
	CansNeeded []CanLine `json:"cans"`
}

func (cn CansNeededResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func SolveCansNeeded(w http.ResponseWriter, r *http.Request) {
	data := &CansNeededRequest{}
	if err := render.Bind(r, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%#v\n", data)

	walls := []domain.Wall{}
	for _, wall := range data.Room {
		domainWall, err := domain.NewWall(wall.Width, wall.Height, wall.QtyWindows, wall.QtyDoors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		walls = append(walls, *domainWall)

	}
	room, err := domain.NewRoom(walls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	target := room.PaintableArea() / 5

	solver := domain.CansNeeded(domain.RealCans{})

	id, _ := uuid.NewRandom()
	result, solvedFor := solver(target)

	cansNeeded := []CanLine{}

	for k, v := range result {
		id, _ := uuid.NewRandom()
		cansNeeded = append(cansNeeded, CanLine{Id: id, Label: k, Qty: v})
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, CansNeededResponse{
		Id:         id,
		Target:     target,
		SolvedFor:  solvedFor,
		CansNeeded: cansNeeded,
	})

}
