package main

import "testing"

func Test_cellPower(t *testing.T) {
	type args struct {
		x      int
		y      int
		serial int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{3, 5, 8},
			want: 4,
		},
		{
			name: "2",
			args: args{122, 79, 57},
			want: -5,
		},
		{
			name: "3",
			args: args{217, 196, 39},
			want: 0,
		},
		{
			name: "4",
			args: args{101, 153, 71},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellPower(tt.args.x, tt.args.y, tt.args.serial); got != tt.want {
				t.Errorf("cellPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
