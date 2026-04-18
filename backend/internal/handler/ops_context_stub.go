package handler

import "github.com/gin-gonic/gin"

const (
	opsModelKey     = "ops.request_model"
	opsAccountIDKey = "ops.account_id"
)

// Ops instrumentation has been removed from the lightweight build.
// Keep no-op helpers so gateway handlers stay source-compatible while
// the remaining call sites are phased out incrementally.
func setOpsRequestContext(_ *gin.Context, _ string, _ bool, _ []byte) {}

func setOpsEndpointContext(_ *gin.Context, _ string, _ int16) {}

func setOpsSelectedAccount(_ *gin.Context, _ int64, _ string) {}
