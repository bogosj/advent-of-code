package main

import (
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
