package credentials

import (
	"github.com/google/uuid"
	"testing"
)

func TestSessionsImpl(t *testing.T) {
	S = NewSessions()
	sessionToken := S.CreateSession(uuid.New())
	sessionExists := S.CheckSession(sessionToken)
	if !sessionExists {
		t.Error("Unable to check created session")
	}
	S.DeleteSession(sessionToken)
	sessionExists = S.CheckSession(sessionToken)
	if sessionExists {
		t.Error("Unable to check deleted session")
	}
}