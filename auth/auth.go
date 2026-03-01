package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	maxAge = 86400 * 30 // 30 days
)

// Init sets up the Goth authentication providers and session store.
func Init() {
	// Setup session store (for goth)
	sessionKey := os.Getenv("GOTH_SESSION_KEY")
	if sessionKey == "" {
		sessionKey = "secret-key-for-goth" // fallback for development
		log.Println("WARNING: Using default session key. Set GOTH_SESSION_KEY in production!")
	}

	store := sessions.NewCookieStore([]byte(sessionKey))
	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = os.Getenv("ENVIRONMENT") == "production"
	gothic.Store = store

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if googleClientID == "" || googleClientSecret == "" || callbackURL == "" {
		log.Println("WARNING: Missing Google OAuth environment variables!")
	}

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, callbackURL),
	)
}
