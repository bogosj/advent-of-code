package main

import (
	"strings"
	"testing"
)

var (
	sampleText = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`
)

func Test_countSteps(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{in: strings.Fields(sampleText)},
			want: 58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSteps(tt.args.in); got != tt.want {
				t.Errorf("countSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
