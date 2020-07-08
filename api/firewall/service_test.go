package firewall

import (
	"context"
	"testing"

	"github.com/shellhub-io/shellhub/api/store/memory"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetFirewallRule(t *testing.T) {
	store := memory.NewStore()
	svc := NewService(store)
	ctx := context.TODO()

	rule := &models.FirewallRule{
		ID:       "id",
		TenantID: "tenantid",
		FirewallRuleFields: models.FirewallRuleFields{
			Action:   "action",
			Priority: 1,
			Active:   true,
			SourceIP: "127.0.0.1",
			Username: "username",
			Hostname: "hostname",
		},
	}

	err := store.CreateFirewallRule(ctx, rule)
	assert.NoError(t, err)

	r, err := svc.GetRule(ctx, rule.ID)
	assert.NoError(t, err)
	assert.Equal(t, rule, r)
}
