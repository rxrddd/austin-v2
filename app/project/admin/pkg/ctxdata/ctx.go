package ctxdata

import "context"

const AdministratorIdKey = "kratos-AdministratorId"
const AdministratorUsername = "kratos-AdministratorUsername"

func GetAdminId(ctx context.Context) int64 {
	if v, ok := ctx.Value(AdministratorIdKey).(int64); ok {
		return v
	}
	return 0
}

func GetAdminName(ctx context.Context) string {
	if v, ok := ctx.Value(AdministratorUsername).(string); ok {
		return v
	}
	return ""
}
