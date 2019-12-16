package intmath

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "many, including negative",
			args: args{[]int{1, 2, 45, 12345, -123, -8888}},
			want: -8888,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.in...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "many, including negative",
			args: args{[]int{1, 2, 45, 12345, -123, -8888}},
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.in...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{4},
			want: 4,
		},
		{
			name: "negative",
			args: args{-4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.i); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	type args struct {
		a        int
		b        int
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two",
			args: args{a: 5, b: 6},
			want: 30,
		},
		{
			name: "more",
			args: args{a: 5, b: 6, integers: []int{5, 7}},
			want: 210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.args.a, tt.args.b, tt.args.integers...); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small",
			args: args{a: 5, b: 6},
			want: 1,
		},
		{
			name: "big",
			args: args{a: 555555555, b: 111},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
