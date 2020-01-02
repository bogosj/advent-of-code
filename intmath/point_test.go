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
			wantRet: []Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}},
		},
		{
			name:    "elsewhere",
			fields:  fields{X: 1, Y: 1},
			wantRet: []Point{{X: 0, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 2}},
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
