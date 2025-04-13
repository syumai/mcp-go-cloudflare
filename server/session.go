//go:build !js

package server

import "sync"

type sessionStore struct {
	sessions sync.Map
}

func newSessionStore() SessionStore {
	return &sessionStore{}
}

func (s *sessionStore) Store(sessionID string, session *sseSession) {
	s.sessions.Store(sessionID, session)
}

func (s *sessionStore) Range(f func(sessionID string, value any) bool) {
	s.sessions.Range(func(sessionID, value any) bool {
		return f(sessionID.(string), value.(*sseSession))
	})
}

func (s *sessionStore) Load(sessionID string) (*sseSession, bool) {
	session, ok := s.sessions.Load(sessionID)
	if !ok {
		return nil, false
	}
	return session.(*sseSession), true
}

func (s *sessionStore) Delete(sessionID string) {
	s.sessions.Delete(sessionID)
}
