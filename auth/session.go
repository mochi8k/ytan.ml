package auth

import (
	"github.com/satori/go.uuid"
)

type sessionManger struct {
	sessionsMap map[uuid.UUID]*session
}

func (sm *sessionManger) addSession(s *session) {
	sm.sessionsMap[s.getSessionID()] = s
}

func (sm sessionManger) GetSession(sessionID uuid.UUID) *session {
	return sm.sessionsMap[sessionID]
}

type session struct {
	sessionID uuid.UUID
	userName  string
}

func (s session) getSessionID() uuid.UUID {
	return s.sessionID
}

func newSession(userName string) *session {
	session := &session{
		sessionID: uuid.NewV4(),
		userName:  userName,
	}
	return session
}
