package main

import (
	"strings"
	"testing"
)

func Test_departureTime(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{in: strings.Split("7,13,x,x,59,x,31,19", ",")},
			want: 1068781,
		},
		{
			name: "2",
			args: args{in: strings.Split("67,7,59,61", ",")},
			want: 754018,
		},
		{
			name: "3",
			args: args{in: strings.Split("67,x,7,59,61", ",")},
			want: 779210,
		},
		{
			name: "4",
			args: args{in: strings.Split("67,7,x,59,61", ",")},
			want: 1261476,
		},
		{
			name: "5",
			args: args{in: strings.Split("1789,37,47,1889", ",")},
			want: 1202161486,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := departureTime(tt.args.in); got != tt.want {
				t.Errorf("departureTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
