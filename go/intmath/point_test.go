package intmath

import (
	"reflect"
	"testing"
)

func TestPoint_Neighbors(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet []Point
	}{
		{
			name:    "origin",
			fields:  fields{X: 0, Y: 0},
			wantRet: []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}},
		},
		{
			name:    "elsewhere",
			fields:  fields{X: 1, Y: 1},
			wantRet: []Point{{1, 0}, {0, 1}, {2, 1}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if gotRet := p.Neighbors(); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Point.Neighbors() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestPoint_ManhattanDistanceTo(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		op Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "basic",
			fields: fields{X: 0, Y: 0},
			args:   args{op: Point{X: 1, Y: 1}},
			want:   2,
		},
		{
			name:   "both positive",
			fields: fields{X: 5, Y: 5},
			args:   args{op: Point{X: 1, Y: 1}},
			want:   8,
		},
		{
			name:   "both negative",
			fields: fields{X: -5, Y: -5},
			args:   args{op: Point{X: -1, Y: -1}},
			want:   8,
		},
		{
			name:   "one negative",
			fields: fields{X: -1, Y: -2},
			args:   args{op: Point{X: 1, Y: 1}},
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.ManhattanDistanceTo(tt.args.op); got != tt.want {
				t.Errorf("Point.ManhattanDistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_AllNeighbors(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet []Point
	}{
		{
			name:    "simple",
			fields:  fields{X: 1, Y: 1},
			wantRet: []Point{{0, 0}, {1, 0}, {2, 0}, {0, 1}, {2, 1}, {0, 2}, {1, 2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if gotRet := p.AllNeighbors(); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Point.AllNeighbors() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestPoint_Neighbor(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		dir rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOp Point
	}{
		{
			name:   "origin up",
			fields: fields{},
			args:   args{dir: 'U'},
			wantOp: Point{X: 0, Y: -1},
		},
		{
			name:   "origin down",
			fields: fields{},
			args:   args{dir: 'S'},
			wantOp: Point{X: 0, Y: 1},
		},
		{
			name:   "origin left",
			fields: fields{},
			args:   args{dir: 'l'},
			wantOp: Point{X: -1, Y: 0},
		},
		{
			name:   "origin right",
			fields: fields{},
			args:   args{dir: 'E'},
			wantOp: Point{X: 1, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if gotOp := p.Neighbor(tt.args.dir); !reflect.DeepEqual(gotOp, tt.wantOp) {
				t.Errorf("Point.Neighbor() = %v, want %v", gotOp, tt.wantOp)
			}
		})
	}
}

func TestPoint_ContainedIn(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		topLeft     Point
		bottomRight Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "simple yes",
			fields: fields{X: 2, Y: 2},
			args:   args{topLeft: Point{X: 1, Y: 3}, bottomRight: Point{X: 3, Y: 1}},
			want:   true,
		},
		{
			name:   "simple no",
			fields: fields{X: 2, Y: 7},
			args:   args{topLeft: Point{X: 1, Y: 3}, bottomRight: Point{X: 3, Y: 1}},
			want:   false,
		},
		{
			name:   "on edge",
			fields: fields{X: 3, Y: 1},
			args:   args{topLeft: Point{X: 1, Y: 3}, bottomRight: Point{X: 3, Y: 1}},
			want:   true,
		},
		{
			name:   "negative box",
			fields: fields{X: 4, Y: -2},
			args:   args{topLeft: Point{X: 3, Y: -1}, bottomRight: Point{X: 5, Y: -3}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.ContainedIn(tt.args.topLeft, tt.args.bottomRight); got != tt.want {
				t.Errorf("Point.ContainedIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
