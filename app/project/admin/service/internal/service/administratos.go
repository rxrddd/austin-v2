package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminInterface) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.authUseCase.Login(ctx, req)
}

func (s *AdminInterface) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}

func (s *AdminInterface) GetAdministratorInfo(ctx context.Context, empty *emptypb.Empty) (*v1.GetAdministratorInfoReply, error) {
	res, err := s.administratorUseCase.GetAdministrator(ctx, ctx.Value("kratos-AdministratorId").(int64))
	if err != nil {
		return nil, err
	}
	return &v1.GetAdministratorInfoReply{
		Id:        res.Id,
		Username:  res.Username,
		Mobile:    res.Mobile,
		Nickname:  res.Mobile,
		Avatar:    res.Avatar,
		Status:    res.Status,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}, nil
}
