package credentials

import "github.com/google/uuid"

var S = NewSessions()

type Sessions interface {
	CreateSession(userUuid uuid.UUID) uuid.UUID
	CheckSession(token uuid.UUID) bool
	DeleteSession(token uuid.UUID)
}

type SessionsImpl struct {
	tokenToSession map[uuid.UUID]Session
}

type Session struct {
	userUuid uuid.UUID
	token uuid.UUID
}

func NewSessions() Sessions {
	return &SessionsImpl{
		tokenToSession: map[uuid.UUID]Session{},
	}
}

func (s *SessionsImpl) CreateSession(userUuid uuid.UUID) uuid.UUID {
	session := Session{
		userUuid: userUuid,
		token: uuid.New(),
	}

	s.tokenToSession[session.token] = session
	return session.token
}

func (s *SessionsImpl) CheckSession(token uuid.UUID) bool {
	_, exists := s.tokenToSession[token]
	return exists
}

func (s *SessionsImpl) DeleteSession(token uuid.UUID) {
	if s.tokenToSession[token].token == token {
		delete(s.tokenToSession, token)
	}
}
