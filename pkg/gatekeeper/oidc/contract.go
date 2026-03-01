package oidc

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidProviderConfig = errors.New("invalid provider configuration")
	ErrTokenExchangeFailed   = errors.New("token exchange failed")
	ErrUserInfoFailed        = errors.New("failed to get user info")
	ErrIDTokenVerification   = errors.New("id token verification failed")
	ErrInvalidSignature      = errors.New("invalid signature")
)

// Token represents the OAuth2 tokens
type Token struct {
	AccessToken  string    `json:"accessToken"`
	TokenType    string    `json:"tokenType"`
	RefreshToken string    `json:"refreshToken,omitempty"`
	Expiry       time.Time `json:"expiration"`
	IDToken      string    `json:"idToken,omitempty"`
}

// VerifiedToken represents a verified OIDC token
type VerifiedToken struct {
	AccessToken string
	Claims      Claims
}

// Claims represents the claims extracted from an OIDC token
type Claims struct {
	Subject        string `json:"sub"`
	Username       string `json:"preferred_username"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
	Issuer         string `json:"iss"`
	Audience       string `json:"aud"`
	ExpirationTime int64  `json:"exp"`
	IssuedAt       int64  `json:"iat"`
	Role           string `json:"role,omitempty"`
	Nonce          string `json:"nonce,omitempty"`
}

// UserInfo represents the user information from the OIDC provider
type UserInfo struct {
	Subject       string `json:"sub"`
	Issuer        string `json:"iss"`
	Username      string `json:"preferred_username,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	Role          string `json:"role,omitempty"`
}

// OIDCProvider interface for OIDC operations
type OIDCProvider interface {
	// BeginAuth starts the authentication flow and returns the auth URL
	BeginAuth(ctx context.Context, state, redirectURL, codeChallenge, codeChallengeMethod string, extraParams map[string]string) (string, error)

	// ExchangeToken exchanges the authorization code for tokens
	ExchangeToken(ctx context.Context, code, verifier, redirectURL string) (Token, error)

	// VerifyIDToken verifies the ID token (for OIDC)
	VerifyIDToken(ctx context.Context, idToken string) (string, error)

	// GetUserInfo gets user information using the access token and may use the ID token to extract roles
	GetUserInfo(ctx context.Context, accessToken string, idToken string) (UserInfo, error)

	GetTokenClaims(ctx context.Context, token string) (Claims, error)
}
