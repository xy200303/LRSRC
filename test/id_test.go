package utils

import (
	"log"
	"testing"
)

func TestGenerateLRVEID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateLRID("LRVE")
			log.Println(got)
		})
	}
}
