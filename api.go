package apisdkgo

import "github.com/LumeWeb/siacentral-api/sia"

// NewSiaClient intializes a new Sia Central API client
func NewSiaClient() *sia.APIClient {
	return sia.NewClient()
}
