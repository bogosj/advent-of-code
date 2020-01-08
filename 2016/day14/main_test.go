package main

import "testing"

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
