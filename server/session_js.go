//go:build js && wasm

package server

// implement SessionStore interface using Durable Objects
type sessionStore struct{}

func newSessionStore() SessionStore {
	// TODO
	return nil
}

func (s *sessionStore) Store(sessionID string, session *sseSession) {
	// TODO
}

func (s *sessionStore) Range(f func(key, value *sseSession) bool) {
	// TODO
}

func (s *sessionStore) Load(sessionID string) (*sseSession, bool) {
	// TODO
	return nil, false
}

func (s *sessionStore) Delete(sessionID string) {
	// TODO
}
