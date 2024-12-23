package account

import (
	"errors"
	"testing"
)

func TestGetAccount(t *testing.T) {
	repo := FetchUserAccountRepository()

	tests := []struct {
		name     string
		username string
		password string
		expected *UserAccount
		err      error
	}{
		{
			name:     "Valid admin account",
			username: "admin",
			password: "admin",
			expected: &UserAccount{
				Username:   "admin",
				Password:   "admin",
				CustomerID: "1234567890",
			},
			err: nil,
		},
		{
			name:     "Invalid account",
			username: "user",
			password: "password",
			expected: nil,
			err:      errors.New("account not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := repo.GetAccount(tt.username, tt.password)
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("expected error %v, got %v", tt.err, err)
				}
			} else {
				if account == nil || *account != *tt.expected {
					t.Errorf("expected account %v, got %v", tt.expected, account)
				}
			}
		})
	}
}

func TestAddAccount(t *testing.T) {
	repo := FetchUserAccountRepository()

	newAccount := UserAccount{
		Username:   "newuser",
		Password:   "newpassword",
		CustomerID: "0987654321",
	}

	repo.AddAccount(newAccount)

	account, err := repo.GetAccount("newuser", "newpassword")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if account == nil || *account != newAccount {
		t.Errorf("expected account %v, got %v", newAccount, account)
	}
}