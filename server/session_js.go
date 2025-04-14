//go:build js && wasm

package server

// implement SessionIDStore interface using Durable Objects
type sessionIDStore struct{}

func newSessionIDStore() SessionIDStore {
	// TODO
	return nil
}

func (s *sessionIDStore) Store(sessionID string) {
	// TODO
}

func (s *sessionIDStore) Range(f func(sessionID string) bool) {
	// TODO
}

func (s *sessionIDStore) Load(sessionID string) bool {
	// TODO
	return false
}

func (s *sessionIDStore) Delete(sessionID string) {
	// TODO
}
