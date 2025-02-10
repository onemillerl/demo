package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) uint32 {
	if ctx.Value(UserIdKey) == nil {
		return 0
	}
	// return uint32(ctx.Value(UserIdKey).(float64))
	return uint32(ctx.Value(UserIdKey).(int32))
}

func GetTokenFromCtx(ctx context.Context) string {
	if ctx.Value(TokenKey) == nil {
		return "0"
	}
	// return uint32(ctx.Value(UserIdKey).(float64))
	return string(ctx.Value(TokenKey).(string))
}
