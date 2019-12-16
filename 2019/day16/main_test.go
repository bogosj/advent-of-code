package main

import (
	"reflect"
	"testing"
)

func Test_pattern(t *testing.T) {
	type args struct {
		elt    int
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name:    "first",
			args:    args{elt: 1, length: 12},
			wantRet: []int{1, 0, -1, 0, 1, 0, -1, 0, 1, 0, -1, 0},
		}, {
			name:    "second",
			args:    args{elt: 2, length: 12},
			wantRet: []int{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0},
		}, {
			name:    "third",
			args:    args{elt: 3, length: 12},
			wantRet: []int{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1, 0},
		}, {
			name:    "eighth",
			args:    args{elt: 8, length: 12},
			wantRet: []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := pattern(tt.args.elt, tt.args.length); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("pattern() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_fft(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "phase 1",
			args: args{in: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			want: []int{4, 8, 2, 2, 6, 1, 5, 8},
		}, {
			name: "phase 2",
			args: args{in: []int{4, 8, 2, 2, 6, 1, 5, 8}},
			want: []int{3, 4, 0, 4, 0, 4, 3, 8},
		}, {
			name: "phase 3",
			args: args{in: []int{3, 4, 0, 4, 0, 4, 3, 8}},
			want: []int{0, 3, 4, 1, 5, 5, 1, 8},
		}, {
			name: "phase 4",
			args: args{in: []int{0, 3, 4, 1, 5, 5, 1, 8}},
			want: []int{0, 1, 0, 2, 9, 4, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fft(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstEight(t *testing.T) {
	type args struct {
		in     []int
		phases int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{in: []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}, phases: 100},
			want: []int{2, 4, 1, 7, 6, 1, 7, 6},
		}, {
			name: "test 2",
			args: args{in: []int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7}, phases: 100},
			want: []int{7, 3, 7, 4, 5, 4, 1, 8},
		}, {
			name: "test 3",
			args: args{in: []int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3}, phases: 100},
			want: []int{5, 2, 4, 3, 2, 1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstEight(tt.args.in, tt.args.phases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("firstEight() = %v, want %v", got, tt.want)
			}
		})
	}
}
