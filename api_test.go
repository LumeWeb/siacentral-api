package apisdkgo

import (
	"reflect"
	"testing"

	"github.com/siacentral/apisdkgo/scprime"
	"github.com/siacentral/apisdkgo/sia"
)

func TestNewSiaClient(t *testing.T) {
	tests := []struct {
		name string
		want *sia.APIClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSiaClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSiaClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewScPrimeClient(t *testing.T) {
	tests := []struct {
		name string
		want *scprime.APIClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScPrimeClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScPrimeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
