package test

import (
	"testing"
	"xiaoyun/backend/utils"
)

func TestGenerateRandomJwtSecret(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "11",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GenerateRandomJwtSecret()
			t.Logf("GenerateRandomJwtSecret() output: %v\n", got)
		})
	}
}

func TestGenJWTToken(t *testing.T) {
	type args struct {
		username string
		role     int64
		expires  int64
		secret   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试1",
			args: args{
				username: "小云",
				role:     1,
				expires:  1,
				secret:   "test",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.GenJWTToken(tt.args.username, uint8(tt.args.role), tt.args.expires, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GenJWTToken() output: %v\n", got)
		})
	}
}
