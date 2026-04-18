//go:build unit

package service

import (
	"context"
	"sync"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/pkg/pagination"
	"github.com/stretchr/testify/require"
)

// --- mock: UserRepository ---

type mockUserRepo struct {
}

func (m *mockUserRepo) Create(context.Context, *User) error               { return nil }
func (m *mockUserRepo) GetByID(context.Context, int64) (*User, error)     { return &User{}, nil }
func (m *mockUserRepo) GetByEmail(context.Context, string) (*User, error) { return &User{}, nil }
func (m *mockUserRepo) GetFirstAdmin(context.Context) (*User, error)      { return &User{}, nil }
func (m *mockUserRepo) Update(context.Context, *User) error               { return nil }
func (m *mockUserRepo) Delete(context.Context, int64) error               { return nil }
func (m *mockUserRepo) List(context.Context, pagination.PaginationParams) ([]User, *pagination.PaginationResult, error) {
	return nil, nil, nil
}
func (m *mockUserRepo) ListWithFilters(context.Context, pagination.PaginationParams, UserListFilters) ([]User, *pagination.PaginationResult, error) {
	return nil, nil, nil
}
func (m *mockUserRepo) UpdateConcurrency(context.Context, int64, int) error { return nil }
func (m *mockUserRepo) ExistsByEmail(context.Context, string) (bool, error) { return false, nil }
func (m *mockUserRepo) RemoveGroupFromAllowedGroups(context.Context, int64) (int64, error) {
	return 0, nil
}
func (m *mockUserRepo) AddGroupToAllowedGroups(context.Context, int64, int64) error { return nil }
func (m *mockUserRepo) UpdateTotpSecret(context.Context, int64, *string) error      { return nil }
func (m *mockUserRepo) EnableTotp(context.Context, int64) error                     { return nil }
func (m *mockUserRepo) DisableTotp(context.Context, int64) error                    { return nil }
func (m *mockUserRepo) RemoveGroupFromUserAllowedGroups(context.Context, int64, int64) error {
	return nil
}

// --- mock: APIKeyAuthCacheInvalidator ---

type mockAuthCacheInvalidator struct {
	invalidatedUserIDs []int64
	mu                 sync.Mutex
}

func (m *mockAuthCacheInvalidator) InvalidateAuthCacheByKey(context.Context, string)    {}
func (m *mockAuthCacheInvalidator) InvalidateAuthCacheByGroupID(context.Context, int64) {}
func (m *mockAuthCacheInvalidator) InvalidateAuthCacheByUserID(_ context.Context, userID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.invalidatedUserIDs = append(m.invalidatedUserIDs, userID)
}

// --- 测试 ---

func TestNewUserService_FieldsAssignment(t *testing.T) {
	repo := &mockUserRepo{}
	auth := &mockAuthCacheInvalidator{}
	svc := NewUserService(repo, nil, auth, nil)
	require.NotNil(t, svc)
	require.Equal(t, repo, svc.userRepo)
	require.Equal(t, auth, svc.authCacheInvalidator)
	require.Nil(t, svc.billingCache)
}
