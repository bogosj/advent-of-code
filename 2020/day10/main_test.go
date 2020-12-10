package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name       string
		args       args
		wantOnes   int
		wantThrees int
	}{
		{
			name:       "basic",
			args:       args{in: []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}},
			wantOnes:   7,
			wantThrees: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOnes, gotThrees := part1(tt.args.in)
			if gotOnes != tt.wantOnes {
				t.Errorf("part1() gotOnes = %v, want %v", gotOnes, tt.wantOnes)
			}
			if gotThrees != tt.wantThrees {
				t.Errorf("part1() gotThrees = %v, want %v", gotThrees, tt.wantThrees)
			}
		})
	}
}
