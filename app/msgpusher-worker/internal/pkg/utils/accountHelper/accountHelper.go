package accountHelper

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"context"
	"encoding/json"
	"fmt"
)

func GetAccount(ctx context.Context, sc *biz.SendAccountUseCase, sendAccount int, v interface{}) error {
	one, err := sc.One(ctx, int64(sendAccount))
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
