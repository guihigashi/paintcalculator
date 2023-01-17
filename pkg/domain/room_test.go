package domain

import (
	"reflect"
	"testing"

	"github.com/samber/lo"
)

func TestNewRoom(t *testing.T) {
	type args struct {
		walls []Wall
	}
	validRoom := lo.Times(4, func(index int) Wall {
		w, _ := NewWall(5, 5, 1, 1)
		return *w
	})

	tests := []struct {
		name    string
		args    args
		want    *Room
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{validRoom},
			want:    &Room{validRoom},
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{validRoom[0:3]},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoom(tt.args.walls)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
