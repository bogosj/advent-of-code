package main

import "testing"

func Test_boardingPass_id(t *testing.T) {
	type fields struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "BFFFBBFRRR",
			fields: fields{data: "BFFFBBFRRR"},
			want:   567,
		},
		{
			name:   "FFFBBBFRRR",
			fields: fields{data: "FFFBBBFRRR"},
			want:   119,
		},
		{
			name:   "BBFFBBFRLL",
			fields: fields{data: "BBFFBBFRLL"},
			want:   820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boardingPass{
				data: tt.fields.data,
			}
			if got := b.id(); got != tt.want {
				t.Errorf("boardingPass.id() = %v, want %v", got, tt.want)
			}
		})
	}
}
