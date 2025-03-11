package test

import (
	"testing"
	"xiaoyun/backend/utils"
)

func TestGenRandomNumber(t *testing.T) {
	type args struct {
		length uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11",
			args: args{
				length: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GenRandomNumber(tt.args.length); got != tt.want {
				t.Errorf("GenRandomNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
