package main

import (
	"testing"
)

func Test_explode(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{in: "[[[[[9,8],1],2],3],4]"},
			want: "[[[[0,9],2],3],4]",
		},
		{
			name: "2",
			args: args{in: "[7,[6,[5,[4,[3,2]]]]]"},
			want: "[7,[6,[5,[7,0]]]]",
		},
		{
			name: "3",
			args: args{in: "[[6,[5,[4,[3,2]]]],1]"},
			want: "[[6,[5,[7,0]]],3]",
		},
		{
			name: "4",
			args: args{in: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"},
			want: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name: "5",
			args: args{in: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			want: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name: "doesn't need",
			args: args{in: "[[3,4],5]"},
			want: "[[3,4],5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := explode(tt.args.in); got != tt.want {
				t.Errorf("explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{in: "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
			want: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name: "first",
			args: args{in: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
			want: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.in); got != tt.want {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		in1 string
		in2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{in1: "[1,2]", in2: "[[3,4],5]"},
			want: "[[1,2],[[3,4],5]]",
		},
		{
			name: "needs reduce",
			args: args{in1: "[[[[4,3],4],4],[7,[[8,4],9]]]", in2: "[1,1]"},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.in1, tt.args.in2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_magnitude(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{in: "[[1,2],[[3,4],5]]"},
			want: 143,
		},
		{
			name: "2",
			args: args{in: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
			want: 1384,
		},
		{
			name: "3",
			args: args{in: "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
			want: 445,
		},
		{
			name: "4",
			args: args{in: "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
			want: 791,
		},
		{
			name: "5",
			args: args{in: "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
			want: 1137,
		},
		{
			name: "6",
			args: args{in: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
			want: 3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnitude(tt.args.in); got != tt.want {
				t.Errorf("magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
