package accountHelper

import (
	"austin-v2/app/msgpusher-common/model"
	"context"
	"encoding/json"
	"fmt"
)

func GetAccount(ctx context.Context, sc IAccount, sendAccount int64, v interface{}) error {
	one, err := sc.One(ctx, sendAccount)
	if err != nil {
		return err
	}
	if one == nil {
		return fmt.Errorf("获取账号异常 sendAccount: %d", sendAccount)
	}

	err = json.Unmarshal([]byte(one.Config), &v)
	if err != nil {
		return err
	}
	return nil
}

type IAccount interface {
	One(ctx context.Context, id int64) (item *model.SendAccount, err error)
}
