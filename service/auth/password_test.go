package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating jwt, %v", err)
	}
	if token == "" {
		t.Errorf("expected token to be not empty")
	}
}
