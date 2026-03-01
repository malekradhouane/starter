package oidc

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// GeneratePKCE returns code_verifier and code_challenge (S256)
func GeneratePKCE() (verifier string, challenge string) {
	// verifier: 43-128 characters
	b := make([]byte, 64)
	_, _ = rand.Read(b)
	verifier = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(sum[:])
	return verifier, challenge
}

// GenerateNonce returns a cryptographically-random base64url-encoded string suitable for OIDC nonce
func GenerateNonce() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}
