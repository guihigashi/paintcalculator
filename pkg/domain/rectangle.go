package domain

import "errors"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Valid() bool {
	return r.Width > 0 && r.Height > 0
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	r := Rectangle{width, height}

	if !r.Valid() {
		return nil, errors.New("width and height must be greater than zero")
	}

	return &r, nil
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
