package domain

import (
	"errors"

	"github.com/samber/lo"
)

type Wall struct {
	Rectangle
	Windows []Rectangle
	Doors   []Rectangle
}

func NewWall(width, height float64, qtyWindows, qtyDoors int) (*Wall, error) {
	wall := &Wall{
		Rectangle: Rectangle{Width: width, Height: height},
		Windows: lo.Times(qtyWindows, func(index int) Rectangle {
			return Window()
		}),
		Doors: lo.Times(qtyDoors, func(index int) Rectangle {
			return Door()
		}),
	}

	if err := wall.HasErrors(); err != nil {
		return nil, err
	}

	return wall, nil
}

func (w Wall) HasErrors() error {
	wallArea := w.Area()

	// validação 1
	if wallArea < 1 || wallArea > 50 {
		return errors.New("nenhuma parede pode ter menos de 1 metro quadrado nem mais de 50 metros quadrados")
	}

	windowsDoorsTotalArea := lo.Reduce(append(w.Windows, w.Doors...),
		func(acc float64, cur Rectangle, _ int) float64 {
			return acc + cur.Area()
		}, 0)
	if windowsDoorsTotalArea > wallArea*0.5 {
		return errors.New("o total de área das portas e janelas deve ser no máximo 50%% da área de parede")
	}

	door := Door()
	if len(w.Doors) > 0 && w.Height < 0.03+door.Height {
		return errors.New("a altura de paredes com porta deve ser, no mínimo, 30 centímetros maior que a altura da porta")
	}

	return nil
}

func (w Wall) PaintableArea() float64 {
	windowsDoorsTotalArea := lo.Reduce(append(w.Windows, w.Doors...),
		func(acc float64, cur Rectangle, _ int) float64 {
			return acc + cur.Area()
		}, 0)
	return w.Area() - windowsDoorsTotalArea
}
