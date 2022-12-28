package service

import (
	"austin-v2/api/authorization/v1"
	"austin-v2/app/authorization/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthorizationService struct {
	v1.UnimplementedAuthorizationServer
	authorizationUsecase *biz.AuthorizationUsecase
	log                  *log.Helper
}

func NewAuthorizationService(AdministratorUseCase *biz.AuthorizationUsecase,
	logger log.Logger) *AuthorizationService {
	return &AuthorizationService{
		log:                  log.NewHelper(log.With(logger, "module", "service/interface")),
		authorizationUsecase: AdministratorUseCase,
	}
}

func (s *AuthorizationService) CheckAuthorization(ctx context.Context, req *v1.CheckAuthorizationRequest) (*v1.CheckReply, error) {
	bc := &biz.Authorization{
		Sub: req.Sub,
		Obj: req.Obj,
		Act: req.Act,
	}

	success, err := s.authorizationUsecase.CheckAuthorization(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, err
}
