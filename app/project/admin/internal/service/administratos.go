package service

import (
	"austin-v2/api/project/admin/v1"
	"austin-v2/app/project/admin/pkg/ctxdata"
	"austin-v2/pkg/utils/timeHelper"
	"context"
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

	// 生成token 将真正的token置于redis中 key为md5的值, 以便冻结用户时，禁止用户登陆
	token, err := s.administratorRepo.GenerateAdministratorToken(ctx, administrator)
	if err != nil {
		return nil, err
	}
	// 更新
	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (s *AdminInterface) LoginSuccess(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	administrator, err := s.administratorRepo.GetAdministrator(ctx, ctxdata.GetAdminId(ctx))
	if err != nil {
		return nil, err
	}
	ip := ctx.Value("RemoteAddr")
	if ip == nil {
		return nil, nil
	}
	loginTime := timeHelper.CurrentTimeYMDHIS()
	administrator.LastLoginTime = loginTime
	administrator.LastLoginIp = ip.(string)
	// 更新用户登陆信息
	if err = s.administratorRepo.AdministratorLoginSuccess(ctx, administrator); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *AdminInterface) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	res, err := s.administratorRepo.GetAdministrator(ctx, ctxdata.GetAdminId(ctx))
	if err != nil {
		return nil, err
	}
	_ = s.administratorRepo.DestroyAdministratorToken(ctx, res)
	return nil, nil
}

func (s *AdminInterface) GetAdministratorList(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	return s.administratorRepo.ListAdministrator(ctx, req)
}

func (s *AdminInterface) GetAdministratorInfo(ctx context.Context, empty *emptypb.Empty) (*v1.AdministratorInfoResponse, error) {
	res, err := s.administratorRepo.GetAdministrator(ctx, ctxdata.GetAdminId(ctx))
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *AdminInterface) GetAdministrator(ctx context.Context, req *v1.GetAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	res, err := s.administratorRepo.GetAdministrator(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return res, err
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

func (s *AdminInterface) ForbidAdministrator(ctx context.Context, req *v1.ForbidAdministratorRequest) (*v1.CheckReply, error) {
	return s.administratorRepo.ForbidAdministrator(ctx, req.Id)
}

func (s *AdminInterface) ApproveAdministrator(ctx context.Context, req *v1.ApproveAdministratorRequest) (*v1.CheckReply, error) {
	return s.administratorRepo.ApproveAdministrator(ctx, req.Id)
}
