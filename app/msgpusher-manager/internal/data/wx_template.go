package data

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/pkg/utils/accountHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type IWxTemplateRepo interface {
	GetOfficialAccountTemplateList(ctx context.Context, sendAccount int64) ([]*message.TemplateItem, error)
}

type wxTemplateRepo struct {
	data *Data
	sc   ISendAccountRepo
	log  *log.Helper
}

func NewWxTemplateRepo(data *Data, sc ISendAccountRepo, logger log.Logger) IWxTemplateRepo {
	return &wxTemplateRepo{
		data: data,
		sc:   sc,
		log:  log.NewHelper(log.With(logger, "module", "data/send_account")),
	}
}

func (s *wxTemplateRepo) GetOfficialAccountTemplateList(ctx context.Context, sendAccount int64) ([]*message.TemplateItem, error) {
	//拼接消息发送
	var acc account.OfficialAccount
	if err := accountHelper.GetAccount(ctx, s.sc, sendAccount, &acc); err != nil {
		return nil, err
	}
	wc := wechat.NewWechat()
	cacheImpl := cache.NewMemory()

	cfg := &offConfig.Config{
		AppID:          acc.AppID,
		AppSecret:      acc.AppSecret,
		Token:          acc.Token,
		EncodingAESKey: acc.EncodingAESKey,
		Cache:          cacheImpl,
	}
	subscribe := wc.GetOfficialAccount(cfg).GetTemplate()
	list, err := subscribe.List()
	return list, err
}
