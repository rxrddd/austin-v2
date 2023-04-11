package transaction

import (
	"austin-v2/common/dal/query"
	"context"
	"gorm.io/gorm"
)

type transactionMgrCtx struct {
}

type tranMgr struct {
	db *gorm.DB
}

type ITranMgr interface {
	Transaction(ctx context.Context, fc func(txCtx context.Context) error) error
	DB(ctx context.Context) *gorm.DB
	Query(ctx context.Context) *query.Query
}

func NewTranMgr(db *gorm.DB) ITranMgr {
	return &tranMgr{
		db: db,
	}
}

func (t *tranMgr) Transaction(ctx context.Context, fc func(txCtx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, transactionMgrCtx{}, tx)
		return fc(ctx)
	})
}

// DB 包裹原生gorm在事务中
func (t *tranMgr) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(transactionMgrCtx{}).(*gorm.DB)
	if ok {
		return tx.WithContext(ctx)
	}
	return t.db.WithContext(ctx)
}

// Query 包裹gorm gen的代码在事务中
func (t *tranMgr) Query(ctx context.Context) *query.Query {
	tx, ok := ctx.Value(transactionMgrCtx{}).(*gorm.DB)
	if ok {
		return query.Use(tx.WithContext(ctx))
	}
	return query.Use(t.db.WithContext(ctx))
}
