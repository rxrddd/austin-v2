package ctxdata

import "context"

const AdministratorIdKey = "kratos-AdministratorId"
const AdministratorUsername = "kratos-AdministratorUsername"

func GetAdminId(ctx context.Context) int64 {
	return ctx.Value(AdministratorIdKey).(int64)
}

func GetAdminName(ctx context.Context) string {
	return ctx.Value(AdministratorUsername).(string)
}
