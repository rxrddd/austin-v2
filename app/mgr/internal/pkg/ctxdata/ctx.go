package ctxdata

import "context"

type gvsUidKey struct {
}
type gvsNameKey struct {
}

func SetMgrID(ctx context.Context, id int32) context.Context {
	return context.WithValue(ctx, gvsUidKey{}, id)
}

func MgrID(ctx context.Context) int32 {
	if v, ok := ctx.Value(gvsUidKey{}).(int32); ok {
		return v
	}
	return 0
}
func SetMgrName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, gvsNameKey{}, name)
}

func MgrName(ctx context.Context) string {
	if v, ok := ctx.Value(gvsNameKey{}).(string); ok {
		return v
	}
	return ""
}
