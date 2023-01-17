package domain

func Latas() []float64 {
	return []float64{0.5, 2.5, 3.6, 18}
}

func Window() Rectangle {
	return Rectangle{Width: 2, Height: 1.2}
}

func Door() Rectangle {
	return Rectangle{Width: 0.8, Height: 1.9}
}
