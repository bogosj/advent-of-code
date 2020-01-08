package main

import (
	"testing"
)

func Test_substrToFind(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "start",
			args:    args{s: "aaaaaaaabbbc"},
			want:    "aaaaa",
			wantErr: false,
		},
		{
			name:    "middle",
			args:    args{s: "abbbc"},
			want:    "bbbbb",
			wantErr: false,
		},
		{
			name:    "end",
			args:    args{s: "abbccc"},
			want:    "ccccc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := substrToFind(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("substrToFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("substrToFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMd5(t *testing.T) {
	type args struct {
		s       string
		i       int
		stretch bool
		cache   map[int]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abc0 - normal",
			args: args{s: "abc", i: 0, stretch: false, cache: map[int]string{}},
			want: "577571be4de9dcce85a041ba0410f29f",
		},
		{
			name: "abc0 - stretched",
			args: args{s: "abc", i: 0, stretch: true, cache: map[int]string{}},
			want: "a107ff634856bb300138cac6568c0f24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMd5(tt.args.s, tt.args.i, tt.args.stretch, tt.args.cache); got != tt.want {
				t.Errorf("getMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
