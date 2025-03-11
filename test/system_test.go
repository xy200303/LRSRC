package main

import (
	"reflect"
	"testing"
)

func TestGetSysConfigByGroupMap(t *testing.T) {
	type args struct {
		sysGroup string
	}
	tests := []struct {
		name    string
		args    args
		want    *map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				sysGroup: "sys_smtp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := system.GetSysConfigMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSysConfigByGroupMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSysConfigByGroupMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
