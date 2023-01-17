package domain

import (
	"reflect"
	"testing"
)

func TestRectangle_Valid(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid",
			fields: fields{2, 2},
			want:   true,
		},
		{
			name:   "invalid width",
			fields: fields{-2, 2},
			want:   false,
		},
		{
			name:   "invalid height",
			fields: fields{2, -2},
			want:   false,
		},
		{
			name:   "invalid width and height",
			fields: fields{-2, -2},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := r.Valid(); got != tt.want {
				t.Errorf("Rectangle.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRectangle(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Rectangle
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{2, 2},
			want:    &Rectangle{2, 2},
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{-2, 2},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRectangle(tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRectangle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "valid",
			fields: fields{2, 2},
			want:   4,
		},
		{
			name:   "invalid",
			fields: fields{0, 0},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
