package google_photos_api_client_go_test

import (
	"github.com/denysvitali/google-photos-api-client-go/v2"
	"testing"
)

func TestNewOAuthConfig(t *testing.T) {
	creds := google_photos_api_client_go.APIAppCredentials{
		ClientID:     "test-client-ID",
		ClientSecret: "test-client-secret",
	}

	got := google_photos_api_client_go.NewOAuthConfig(creds)

	if got.ClientID != creds.ClientID {
		t.Errorf("client ID should be equal: got=%s, want=%s", got.ClientID, creds.ClientID)
	}

	if got.ClientSecret != creds.ClientSecret {
		t.Errorf("client secret should be equal: got=%s, want=%s", got.ClientSecret, creds.ClientSecret)
	}
}
