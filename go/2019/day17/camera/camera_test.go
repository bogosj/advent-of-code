package camera

import (
	"reflect"
	"testing"
)

func Test_moveInst(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name:    "simple",
			args:    args{in: "ABC"},
			wantRet: []int{65, 66, 67, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := moveInst(tt.args.in); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("moveInst() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
