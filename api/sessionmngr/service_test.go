package sessionmngr

import (
	"context"
	"testing"
	"time"

	"github.com/shellhub-io/shellhub/api/store/memory"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSession(t *testing.T) {
	store := memory.NewStore()
	svc := NewService(store)
	ctx := context.TODO()

	session := models.Session{
		UID:           "uid",
		Username:      "username",
		DeviceUID:     "target",
		IPAddress:     "127.0.0.1",
		StartedAt:     time.Now(),
		LastSeen:      time.Now(),
		Authenticated: true,
		Active:        true,
	}

	_, err := store.CreateSession(ctx, session)
	assert.NoError(t, err)

	s, err := svc.GetSession(ctx, models.UID(session.UID))
	assert.NoError(t, err)
	assert.Equal(t, session, *s)
}
