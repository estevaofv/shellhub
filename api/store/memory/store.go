package memory

import (
	"context"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/pkg/api/paginator"
	"github.com/shellhub-io/shellhub/pkg/models"
)

type Store struct {
	store.Store
	devices  map[string]*models.Device
	sessions map[string]*models.Session
	rules    map[string]*models.FirewallRule
}

func NewStore() *Store {
	return &Store{
		devices:  make(map[string]*models.Device),
		sessions: make(map[string]*models.Session),
		rules:    make(map[string]*models.FirewallRule),
	}
}

func (s *Store) ListDevices(ctx context.Context, pagination paginator.Query, filters []models.Filter) ([]models.Device, int, error) {
	return nil, 0, nil
}

func (s *Store) GetDevice(ctx context.Context, uid models.UID) (*models.Device, error) {
	if d, ok := s.devices[string(uid)]; ok {
		return d, nil
	}

	return nil, nil
}

func (s *Store) DeleteDevice(ctx context.Context, uid models.UID) error {
	return nil
}

func (s *Store) AddDevice(ctx context.Context, d models.Device) error {
	s.devices[d.UID] = &d
	return nil
}

func (s *Store) RenameDevice(ctx context.Context, uid models.UID, name string) error {
	return nil
}

func (s *Store) LookupDevice(ctx context.Context, namespace, name string) (*models.Device, error) {
	return nil, nil
}

func (s *Store) UpdateDeviceStatus(ctx context.Context, uid models.UID, online bool) error {
	return nil
}

func (s *Store) UpdatePendingStatus(ctx context.Context, uid models.UID, status string) error {
	return nil
}

func (s *Store) ListSessions(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error) {
	return nil, 0, nil
}

func (s *Store) GetSession(ctx context.Context, uid models.UID) (*models.Session, error) {
	if s, ok := s.sessions[string(uid)]; ok {
		return s, nil
	}

	return nil, nil
}

func (s *Store) SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error {
	return nil
}

func (s *Store) CreateSession(ctx context.Context, session models.Session) (*models.Session, error) {
	s.sessions[session.UID] = &session
	return &session, nil
}

func (s *Store) GetStats(ctx context.Context) (*models.Stats, error) {
	return nil, nil
}

func (s *Store) KeepAliveSession(ctx context.Context, uid models.UID) error {
	return nil
}

func (s *Store) DeactivateSession(ctx context.Context, uid models.UID) error {
	return nil
}

func (s *Store) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}

func (s *Store) GetUserByTenant(ctx context.Context, tenant string) (*models.User, error) {
	return nil, nil
}

func (s *Store) GetDeviceByName(ctx context.Context, name, tenant string) (*models.Device, error) {
	return nil, nil
}

func (s *Store) GetDeviceByUID(ctx context.Context, uid models.UID, tenant string) (*models.Device, error) {
	return nil, nil
}

func (s *Store) CreateFirewallRule(ctx context.Context, rule *models.FirewallRule) error {
	s.rules[rule.ID] = rule
	return nil
}

func (s *Store) ListFirewallRules(ctx context.Context, pagination paginator.Query) ([]models.FirewallRule, int, error) {
	return nil, 0, nil
}

func (s *Store) GetFirewallRule(ctx context.Context, id string) (*models.FirewallRule, error) {
	if r, ok := s.rules[id]; ok {
		return r, nil
	}

	return nil, nil
}

func (s *Store) UpdateFirewallRule(ctx context.Context, id string, rule models.FirewallRuleUpdate) (*models.FirewallRule, error) {
	return nil, nil

}

func (s *Store) DeleteFirewallRule(ctx context.Context, id string) error {
	return nil
}
