package domain

import (
	"reflect"
	"testing"
)

func TestNewWall(t *testing.T) {
	type args struct {
		width      float64
		height     float64
		qtyWindows int
		qtyDoors   int
	}
	tests := []struct {
		name    string
		args    args
		want    *Wall
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				width:      10,
				height:     3,
				qtyWindows: 1,
				qtyDoors:   1,
			},
			want: &Wall{
				Rectangle: Rectangle{Width: 10, Height: 3},
				Windows:   []Rectangle{Window()},
				Doors:     []Rectangle{Door()},
			},
			wantErr: false,
		},
		{
			name: "nenhuma parede pode ter menos de 1 metro quadrado",
			args: args{
				width:      0.9,
				height:     0.9,
				qtyWindows: 0,
				qtyDoors:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nenhuma parede pode ter mais de 50 metros quadrados",
			args: args{
				width:      11,
				height:     5,
				qtyWindows: 0,
				qtyDoors:   0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "o total de área das portas e janelas deve ser no máximo 50% da área de parede",
			// parede 12m^2 de área e portas e janelas com 6.92m^2, ou seja, 58% de 12m^2
			args: args{
				width:      4,
				height:     3,
				qtyWindows: 2,
				qtyDoors:   1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "a altura de paredes com porta deve ser, no mínimo, 30 centímetros maior que a altura da porta",
			args: args{
				width:      6,
				height:     1.929,
				qtyWindows: 0,
				qtyDoors:   1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWall(tt.args.width, tt.args.height, tt.args.qtyWindows, tt.args.qtyDoors)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWall() = %v, want %v", got, tt.want)
			}
		})
	}
}
