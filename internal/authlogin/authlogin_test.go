package authlogin

import (
	"testing"
)

func TestProcessAuthLogin(t *testing.T) {
	tests := []struct {
		name     string
		request  AuthLoginRequest
		expected AuthLoginResponse
	}{
		{
			name: "Valid admin login",
			request: AuthLoginRequest{
				Username: "admin",
				Password: "admin",
			},
			expected: AuthLoginResponse{
				Valid:  true,
				Ticket: "1234567890", // This in random, so it will not be the same
			},
		},
		{
			name: "Invalid login",
			request: AuthLoginRequest{
				Username: "user",
				Password: "password",
			},
			expected: AuthLoginResponse{
				Valid:      false,
				ReasonCode: "INV-100",
				ReasonText: "Oh Dear",
				ReasonUrl:  "http://www.google.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := processAuthLogin(tt.request)

			if tt.expected.Valid {
				if !response.Valid  {
					t.Errorf("expected valid response with ticket %s, got %s", tt.expected.Ticket, response)
				}
			} else {
				if response.Valid || response.ReasonCode != tt.expected.ReasonCode || response.ReasonText != tt.expected.ReasonText || response.ReasonUrl != tt.expected.ReasonUrl {
					t.Errorf("expected invalid response with reason code %s, reason text %s, reason url %s, got %s", tt.expected.ReasonCode, tt.expected.ReasonText, tt.expected.ReasonUrl, response)
				}
			}
		})
	}
}