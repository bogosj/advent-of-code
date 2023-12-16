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

func TestSqrt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "perfect",
			args: args{i: 4},
			want: 2,
		},
		{
			name: "floored",
			args: args{i: 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.i); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "prime",
			args: args{n: 7},
			want: []int{1, 7},
		},
		{
			name: "nine",
			args: args{n: 9},
			want: []int{1, 3, 9},
		},
		{
			name: "100",
			args: args{n: 100},
			want: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
		{
			name: "999999",
			args: args{n: 999999},
			want: []int{1, 3, 7, 9, 11, 13, 21, 27, 33, 37, 39, 63, 77, 91, 99, 111, 117, 143, 189, 231, 259,
				273, 297, 333, 351, 407, 429, 481, 693, 777, 819, 999, 1001, 1221, 1287, 1443, 2079, 2331, 2457,
				2849, 3003, 3367, 3663, 3861, 4329, 5291, 6993, 8547, 9009, 10101, 10989, 12987, 15873, 25641,
				27027, 30303, 37037, 47619, 76923, 90909, 111111, 142857, 333333, 999999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factors() = %v, want %v", got, tt.want)
			}
		})
	}
}
