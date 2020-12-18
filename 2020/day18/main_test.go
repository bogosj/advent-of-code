package main

import (
	"testing"
)

func Test_solveSub(t *testing.T) {
	type args struct {
		eq string
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			name:    "basic",
			args:    args{eq: "2 + 3 * 4"},
			wantRet: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := solveSub(tt.args.eq); gotRet != tt.wantRet {
				t.Errorf("solveSub() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_reduce(t *testing.T) {
	type args struct {
		eq string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{eq: "2 + (4 * 5)"},
			want: "2 + 20",
		},
		{
			name: "complex",
			args: args{eq: "2 + (4 * (3 + 5))"},
			want: "2 + (4 * 8)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce(tt.args.eq); got != tt.want {
				t.Errorf("reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_problem_solve(t *testing.T) {
	type fields struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "basic",
			fields: fields{p: "2 + 4"},
			want:   6,
		},
		{
			name:   "parens",
			fields: fields{p: "2 + (4 * 5)"},
			want:   22,
		},
		{
			name:   "nested parens",
			fields: fields{p: "2 + (4 * (6 + 5))"},
			want:   46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &problem{
				p: tt.fields.p,
			}
			if got := p.solve(); got != tt.want {
				t.Errorf("problem.solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
