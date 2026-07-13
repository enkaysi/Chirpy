package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name          string
		password      string
		hash          string
		wantErr       bool
		matchPassword bool
	}{
		{
			name:          "Correct password",
			password:      password1,
			hash:          hash1,
			wantErr:       false,
			matchPassword: true,
		},
		{
			name:          "Incorrect password",
			password:      "wrongPassword",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Password doesn't match different hash",
			password:      password1,
			hash:          hash2,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Empty password",
			password:      "",
			hash:          hash1,
			wantErr:       false,
			matchPassword: false,
		},
		{
			name:          "Invalid hash",
			password:      password1,
			hash:          "invalidhash",
			wantErr:       true,
			matchPassword: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && match != tt.matchPassword {
				t.Errorf("CheckPasswordHash() expects %v, got %v", tt.matchPassword, match)
			}
		})
	}
}

func TestValidToken(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "pie"

	signed, err := MakeJWT(userID, tokenSecret, 5*time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	getID, err := ValidateJWT(signed, tokenSecret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if getID != userID {
		t.Errorf("expected %v, got %v", userID, getID)
	}
}

func TestExpiredToken(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "pie"

	signed, err := MakeJWT(userID, tokenSecret, -time.Second)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(signed, tokenSecret)
	if err == nil {
		t.Error("expired token expected")
	}
}

func TestWrongSecret(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "pie"

	signed, err := MakeJWT(userID, tokenSecret, 5*time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT failed %v", err)
	}

	_, err = ValidateJWT(signed, "pizza")
	if err == nil {
		t.Error("incorrect key expected")
	}
}
