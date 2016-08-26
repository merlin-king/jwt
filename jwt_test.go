package jwt

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestEncodeAndVerifyToken(t *testing.T) {
	secret := "secret"
	signingHash := HmacSha256(secret)

	payload := NewClaim()
	payload.Set("nbf", fmt.Sprintf("%d", time.Now().Add(-1*time.Hour).Unix()))
	payload.Set("exp", fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()))

	err := json.Unmarshal([]byte(`{"sub":"1234567890","name":"John Doe","admin":true}`), &payload)
	if err != nil {
		t.Fatal(err)
	}

	token, err := Encode(signingHash, payload)
	if err != nil {
		t.Fatal(err)
	}

	err = IsValid(signingHash, token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyToken(t *testing.T) {
	secret := "secret"
	signingHash := HmacSha256(secret)

	payload := NewClaim()
	err := json.Unmarshal([]byte(`{"sub":"1234567890","name":"John Doe","admin":true}`), &payload)
	if err != nil {
		t.Fatal(err)
	}

	token, err := Encode(signingHash, payload)
	if err != nil {
		t.Fatal(err)
	}

	tokenComponents := strings.Split(token, ".")

	invalidSignature := "cBab30RMHrHDcEfxjoYZgeFONFh7Hg"
	invalidToken := tokenComponents[0] + "." + tokenComponents[1] + "." + invalidSignature

	err = IsValid(signingHash, invalidToken)
	if err == nil {
		t.Fatal(err)
	}
}

func TestVerifyTokenExp(t *testing.T) {
	secret := "secret"
	signingHash := HmacSha256(secret)

	payload := NewClaim()
	payload.Set("exp", fmt.Sprintf("%d", time.Now().Add(-1*time.Hour).Unix()))

	err := json.Unmarshal([]byte(`{"sub":"1234567890","name":"John Doe","admin":true}`), &payload)
	if err != nil {
		t.Fatal(err)
	}

	token, err := Encode(signingHash, payload)
	if err != nil {
		t.Fatal(err)
	}

	err = IsValid(signingHash, token)
	if err == nil {
		t.Fatal(err)
	}
}

func TestVerifyTokenNbf(t *testing.T) {
	secret := "secret"
	signingHash := HmacSha256(secret)

	payload := NewClaim()
	payload.Set("nbf", fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()))

	err := json.Unmarshal([]byte(`{"sub":"1234567890","name":"John Doe","admin":true}`), &payload)
	if err != nil {
		t.Fatal(err)
	}

	token, err := Encode(signingHash, payload)
	if err != nil {
		t.Fatal(err)
	}

	err = IsValid(signingHash, token)
	if err == nil {
		t.Fatal(err)
	}
}
