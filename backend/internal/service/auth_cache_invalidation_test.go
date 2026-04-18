//go:build unit

package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsageService_InvalidateUsageCaches(t *testing.T) {
	invalidator := &authCacheInvalidatorStub{}
	svc := &UsageService{authCacheInvalidator: invalidator}

	svc.invalidateUsageCaches(context.Background(), 7)
	require.Equal(t, []int64{7}, invalidator.userIDs)
}
