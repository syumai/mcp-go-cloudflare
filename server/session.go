//go:build !js

package server

import "sync"

type sessionIDStore struct {
	sessions sync.Map
}

func newSessionIDStore() SessionIDStore {
	return &sessionIDStore{}
}

func (s *sessionIDStore) Store(sessionID string) {
	s.sessions.Store(sessionID, true)
}

func (s *sessionIDStore) Range(f func(sessionID string) bool) {
	s.sessions.Range(func(sessionID, value any) bool {
		return f(sessionID.(string))
	})
}

func (s *sessionIDStore) Load(sessionID string) bool {
	_, ok := s.sessions.Load(sessionID)
	return ok
}

func (s *sessionIDStore) Delete(sessionID string) {
	s.sessions.Delete(sessionID)
}
