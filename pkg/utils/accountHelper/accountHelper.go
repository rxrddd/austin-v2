package accountHelper

import (
	"austin-v2/app/msgpusher-common/model"
	"context"
	"encoding/json"
	"errors"
)

var (
	AccountNotFindError   = errors.New("未找到账号")
	AccountUnmarshalError = errors.New("账号解析错误")
)

func GetAccount(ctx context.Context, sc IAccount, sendAccount int64, v interface{}) error {
	one, err := sc.One(ctx, sendAccount)
	if err != nil {
		return err
	}
	if one.ID <= 0 {
		return AccountNotFindError
	}

	err = json.Unmarshal([]byte(one.Config), &v)
	if err != nil {
		return AccountUnmarshalError
	}
	return nil
}

type IAccount interface {
	One(ctx context.Context, id int64) (item model.SendAccount, err error)
}
