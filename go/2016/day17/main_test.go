package main

import "testing"

func Test_shortestPath(t *testing.T) {
	type args struct {
		passcode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ihgpwlah",
			args: args{passcode: "ihgpwlah"},
			want: "DDRRRD",
		},
		{
			name: "kglvqrro",
			args: args{passcode: "kglvqrro"},
			want: "DDUDRLRRUDRD",
		},
		{
			name: "ulqzkmiv",
			args: args{passcode: "ulqzkmiv"},
			want: "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.passcode); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPath(t *testing.T) {
	type args struct {
		passcode string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			name:    "ihgpwlah",
			args:    args{passcode: "ihgpwlah"},
			wantRet: 370,
		},
		{
			name:    "kglvqrro",
			args:    args{passcode: "kglvqrro"},
			wantRet: 492,
		},
		{
			name:    "ulqzkmiv",
			args:    args{passcode: "ulqzkmiv"},
			wantRet: 830,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := longestPath(tt.args.passcode); gotRet != tt.wantRet {
				t.Errorf("longestPath() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
