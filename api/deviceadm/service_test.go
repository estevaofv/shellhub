package deviceadm

import (
	"context"
	"testing"
	"time"

	"github.com/shellhub-io/shellhub/api/store/memory"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetDevice(t *testing.T) {
	store := memory.NewStore()
	svc := NewService(store)
	ctx := context.TODO()

	device := models.Device{
		UID:      "uid",
		Identity: &models.DeviceIdentity{MAC: "mac"},
		Info: &models.DeviceInfo{
			ID:         "id",
			PrettyName: "pretty-name",
			Version:    "version",
		},
		PublicKey: "public-key",
		TenantID:  "tenant-id",
		LastSeen:  time.Now(),
	}

	err := store.AddDevice(ctx, device)
	assert.NoError(t, err)

	d, err := svc.GetDevice(ctx, models.UID(device.UID))
	assert.NoError(t, err)
	assert.Equal(t, device, *d)
}
