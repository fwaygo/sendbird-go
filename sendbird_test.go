package sendbird

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	applicationID := "application_id"
	apiToken := "api_token"
	version := "v3"

	_, err := NewClient(ClientConfig{
		ApplicationID: applicationID,
		APIToken:      apiToken,
		Version:       version,
	})
	if err != nil {
		t.Fatal(err)
	}
}
