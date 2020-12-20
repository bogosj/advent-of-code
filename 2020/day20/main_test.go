package main

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			name:    "basic",
			args:    args{s: "asdf"},
			wantRet: "fdsa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := reverse(tt.args.s); gotRet != tt.wantRet {
				t.Errorf("reverse() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_idsFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name:    "basic",
			args:    args{s: "##...#.#"},
			wantRet: []int{197, 163},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := idsFromString(tt.args.s); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("idsFromString() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
