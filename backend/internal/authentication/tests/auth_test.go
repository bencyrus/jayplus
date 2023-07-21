package tests

import (
	"backend/pkg/authentication"
	"testing"
	"time"
)

func TestGenerateTokenPair(t *testing.T) {
	// Setting up the auth object.
	auth := &authentication.Auth{
		Issuer:             "testIssuer",
		Audience:           "testAudience",
		Secret:             "testSecret",
		AccessTokenExpiry:  1 * time.Hour,
		RefreshTokenExpiry: 24 * time.Hour,
	}

	// Setting up the test user.
	user := &authentication.JWTUser{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
	}

	// Expected results
	expectedAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	expectedRefreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

	cases := []struct {
		name    string
		user    *authentication.JWTUser
		wantErr bool
	}{
		{
			name:    "Valid User",
			user:    user,
			wantErr: false,
		},
		{
			name:    "Nil User",
			user:    nil,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tp, err := auth.GenerateTokenPair(tc.user)
			if (err != nil) != tc.wantErr {
				t.Errorf("GenerateTokenPair() error = %v, wantErr %v", err, tc.wantErr)
			}

			if err == nil {
				if tp.AccessToken[:len(expectedAccessToken)] != expectedAccessToken {
					t.Errorf("Unexpected Access Token Got = %v, want %v", tp.AccessToken, expectedAccessToken)
				}

				if tp.RefreshToken[:len(expectedRefreshToken)] != expectedRefreshToken {
					t.Errorf("Unexpected Refresh Token Got = %v, want %v", tp.RefreshToken, expectedRefreshToken)
				}
			}
		})
	}
}
