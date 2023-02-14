package biz

import (
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type WxTemplateUseCase struct {
	repo data.IWxTemplateRepo
	log  *log.Helper
}

func NewWxTemplateUseCase(repo data.IWxTemplateRepo, logger log.Logger) *WxTemplateUseCase {
	return &WxTemplateUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/send-account-usecase")),
	}
}

func (uc *WxTemplateUseCase) GetOfficialAccountTemplateList(ctx context.Context, req *domain.OfficialAccountTemplateRequest) (*domain.OfficialAccountTemplateResp, error) {
	var (
		items = make([]*domain.OfficialAccountTemplateRow, 0)
	)
	list, err := uc.repo.GetOfficialAccountTemplateList(ctx, req.SendAccount)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		items = append(items, &domain.OfficialAccountTemplateRow{
			TemplateID: item.TemplateID,
			Title:      item.Title,
			Content:    item.Content,
			Example:    item.Example,
		})
	}

	return &domain.OfficialAccountTemplateResp{
		Rows: items,
	}, nil
}
