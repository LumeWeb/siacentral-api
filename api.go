package apisdkgo

import (
	"github.com/siacentral/apisdkgo/sia"
)

// NewSiaClient intializes a new Sia Central API client
func NewSiaClient() *sia.APIClient {
	return sia.NewClient()
}
