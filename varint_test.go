package gokit

import (
	"reflect"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"111", args{uint64(12345)}, []byte{192, 196, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeVarint(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeVarint() = %v, want %v", got, tt.want)
			}
		})
	}
}
