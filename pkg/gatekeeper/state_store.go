package gatekeeper

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

var (
	ErrStateNotFound = errors.New("state not found")
	ErrStateExpired  = errors.New("state expired")
)

// StateData holds the state information for the OAuth2 flow
type StateData struct {
	CodeVerifier string    `json:"code_verifier"`
	RedirectURL  string    `json:"redirect_url"`
	ContinueURL  string    `json:"continue_url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	Nonce        string    `json:"nonce,omitempty"`
	Provider     string    `json:"provider,omitempty"`
	CSRFToken    string    `json:"csrf_token"`
}

type memoryStateStore struct {
	mu     sync.Mutex
	states map[string]memoryStateEntry
}

type memoryStateEntry struct {
	data      StateData
	expiresAt time.Time
}

func NewMemoryStateStore() StateStore {
	return &memoryStateStore{
		states: make(map[string]memoryStateEntry),
	}
}

func (s *memoryStateStore) Save(ctx context.Context, key string, value StateData, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.states[key] = memoryStateEntry{
		data:      value,
		expiresAt: time.Now().Add(ttl),
	}
	return nil
}

func (s *memoryStateStore) Load(ctx context.Context, key string) (StateData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.states[key]
	if !exists {
		return StateData{}, ErrStateNotFound
	}

	if time.Now().After(entry.expiresAt) {
		delete(s.states, key)
		return StateData{}, ErrStateExpired
	}

	return entry.data, nil
}

func (s *memoryStateStore) Delete(ctx context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.states, key)
	return nil
}

func (s *memoryStateStore) VerifyAndLoad(ctx context.Context, key string) (StateData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.states[key]
	if !exists {
		return StateData{}, ErrStateNotFound
	}

	if time.Now().After(entry.expiresAt) {
		delete(s.states, key)
		return StateData{}, ErrStateExpired
	}

	// Delete after successful load (atomic operation)
	delete(s.states, key)
	return entry.data, nil
}

func (s *memoryStateStore) GenerateState(ctx context.Context) (string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
