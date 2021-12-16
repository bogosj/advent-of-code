package main

import (
	"testing"
)

/*
func Test_versionSum(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "8A004A801A8002F478",
			args: args{message: "8A004A801A8002F478"},
			want: 16,
		},
		{
			name: "620080001611562C8802118E34",
			args: args{message: "620080001611562C8802118E34"},
			want: 12,
		},
		{
			name: "C0015000016115A2E0802F182340",
			args: args{message: "C0015000016115A2E0802F182340"},
			want: 23,
		},
		{
			name: "A0016C880162017C3686B18A3D4780",
			args: args{message: "A0016C880162017C3686B18A3D4780"},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got := versionSum(tt.args.message, false); got != tt.want {
				t.Errorf("versionSum() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func Test_versionSum(t *testing.T) {
	type args struct {
		message  string
		inBinary bool
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{

		{
			name: "D2FE28",
			args: args{message: "D2FE28", inBinary: false},
			want: 6, want1: "000",
		},
		{
			name: "38006F45291200",
			args: args{message: "38006F45291200", inBinary: false},
			want: 9, want1: "0000000",
		},
		{
			name: "EE00D40C823060",
			args: args{message: "EE00D40C823060", inBinary: false},
			want: 14, want1: "00000",
		},
		{
			name: "8A004A801A8002F478",
			args: args{message: "8A004A801A8002F478", inBinary: false},
			want: 16, want1: "000",
		},
		{
			name: "620080001611562C8802118E34",
			args: args{message: "620080001611562C8802118E34", inBinary: false},
			want: 12, want1: "00",
		},
		{
			name: "C0015000016115A2E0802F182340",
			args: args{message: "C0015000016115A2E0802F182340", inBinary: false},
			want: 23, want1: "000000",
		},
		{
			name: "A0016C880162017C3686B18A3D4780",
			args: args{message: "A0016C880162017C3686B18A3D4780", inBinary: false},
			want: 31, want1: "0000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := versionSum(tt.args.message, tt.args.inBinary)
			if got != tt.want {
				t.Errorf("versionSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("versionSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
