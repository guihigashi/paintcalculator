package domain

import (
	"reflect"
	"testing"
)

func TestCansNeeded(t *testing.T) {
	type args struct {
		canner      Canner
		paintNeeded float64
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "exemplo da referencia",
			args: args{
				canner:      RealCans{},
				paintNeeded: 19,
			},
			want: map[string]int{
				"18.0L": 1,
				"0.5L":  2,
			},
		},
		{
			name: "deve arredondar pra cima",
			args: args{
				canner:      RealCans{},
				paintNeeded: 0.8,
			},
			want: map[string]int{
				"0.5L": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := CansNeeded(tt.args.canner)

			got, _ := solver(tt.args.paintNeeded)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CansNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
