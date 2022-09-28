package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/data"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminInterface) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	administrator, err := s.administratorRepo.FindLoginAdministratorByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	// 验证密码
	err = s.administratorRepo.VerifyPassword(ctx, administrator.Id, req.Password)
	if err != nil {
		return nil, err
	}
	// 生成token
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"AdministratorId":       administrator.Id,
			"AdministratorUsername": administrator.Username,
			"AdministratorRole":     administrator.Role,
		})
	// 获取jwt key
	signedString, _ := claims.SignedString([]byte(data.GetAuthApiKey()))
	return &v1.LoginReply{
		Token: signedString,
	}, nil
}

func (s *AdminInterface) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AdminInterface) GetAdministratorInfo(ctx context.Context, empty *emptypb.Empty) (*v1.AdministratorInfoResponse, error) {
	res, err := s.administratorRepo.GetAdministrator(ctx, ctx.Value("kratos-AdministratorId").(int64))
	if err != nil {
		return nil, err
	}
	return &v1.AdministratorInfoResponse{
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

func (s *AdminInterface) CreateAdministrator(ctx context.Context, req *v1.CreateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	return s.administratorRepo.CreateAdministrator(ctx, req)
}

func (s *AdminInterface) UpdateAdministrator(ctx context.Context, req *v1.UpdateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	return s.administratorRepo.UpdateAdministrator(ctx, req)
}

func (s *AdminInterface) DeleteAdministrator(ctx context.Context, req *v1.DeleteAdministratorRequest) (*v1.CheckReply, error) {
	return s.administratorRepo.DeleteAdministrator(ctx, req.Id)
}

func (s *AdminInterface) RecoverAdministrator(ctx context.Context, req *v1.RecoverAdministratorRequest) (*v1.CheckReply, error) {
	return s.administratorRepo.RecoverAdministrator(ctx, req.Id)
}

func (s *AdminInterface) GetAdministratorList(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	return s.administratorRepo.ListAdministrator(ctx, req)
}
