package service

import (
	v1 "austin-v2/api/administrator/v1"
	"austin-v2/app/administrator/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AdministratorService struct {
	v1.UnimplementedAdministratorServer
	administratorUseCase *biz.AdministratorUseCase
	log                  *log.Helper
}

func NewAdministratorService(AdministratorUseCase *biz.AdministratorUseCase,
	logger log.Logger) *AdministratorService {
	return &AdministratorService{
		log:                  log.NewHelper(log.With(logger, "module", "service/interface")),
		administratorUseCase: AdministratorUseCase,
	}
}

func (s *AdministratorService) CreateAdministrator(ctx context.Context, req *v1.CreateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	bc := &biz.Administrator{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Mobile:   req.Mobile,
		Status:   req.Status,
		Avatar:   req.Avatar,
		Role:     req.Role,
	}
	administratorInfo, err := s.administratorUseCase.Create(ctx, bc)
	if err != nil {
		return nil, err
	}
	return bizAdministratorToInfoReply(administratorInfo), nil
}
func (s *AdministratorService) UpdateAdministrator(ctx context.Context, req *v1.UpdateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	bc := &biz.Administrator{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Status:   req.Status,
		Avatar:   req.Avatar,
		Role:     req.Role,
	}
	administratorInfo, err := s.administratorUseCase.Update(ctx, bc)
	if err != nil {
		return nil, err
	}
	return bizAdministratorToInfoReply(administratorInfo), nil
}

func (s *AdministratorService) DeleteAdministrator(ctx context.Context, req *v1.DeleteAdministratorRequest) (*v1.CheckReply, error) {
	err := s.administratorUseCase.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, err
}

func (s *AdministratorService) RecoverAdministrator(ctx context.Context, req *v1.RecoverAdministratorRequest) (*v1.CheckReply, error) {
	err := s.administratorUseCase.Recover(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, err
}

func (s *AdministratorService) GetAdministrator(ctx context.Context, req *v1.GetAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	params := map[string]interface{}{}
	params["id"] = req.Id
	params["username"] = req.Username
	params["mobile"] = req.Mobile
	params["role"] = req.Role
	params["status"] = req.Status
	administratorInfo, err := s.administratorUseCase.Get(ctx, params)
	if err != nil {
		return nil, err
	}
	response := bizAdministratorToInfoReply(administratorInfo)
	return response, nil
}

func (s *AdministratorService) ListAdministrator(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	params := make(map[string]interface{})
	params["status"] = req.Status
	params["username"] = req.Username
	params["nickname"] = req.Nickname
	params["mobile"] = req.Mobile
	params["created_at_start"] = req.CreatedAtStart
	params["created_at_end"] = req.CreatedAtEnd
	AdministratorInfoList, count, err := s.administratorUseCase.List(ctx, req.Page, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	response := &v1.ListAdministratorReply{}
	response.Total = count
	for _, v := range AdministratorInfoList {
		response.List = append(response.List, bizAdministratorToInfoReply(v))
	}
	return response, nil
}

func (s *AdministratorService) VerifyAdministratorPassword(ctx context.Context, req *v1.VerifyAdministratorPasswordRequest) (*v1.CheckReply, error) {
	res, err := s.administratorUseCase.VerifyAdministratorPassword(ctx, req)
	return &v1.CheckReply{
		IsSuccess: res,
	}, err
}

func (s *AdministratorService) AdministratorLoginSuccess(ctx context.Context, req *v1.AdministratorLoginSuccessRequest) (*v1.CheckReply, error) {
	err := s.administratorUseCase.UpdateAdministratorLoginInfo(ctx, req)
	isSuccess := false
	if err == nil {
		isSuccess = true
	}
	return &v1.CheckReply{
		IsSuccess: isSuccess,
	}, nil
}

func (s *AdministratorService) AdministratorStatusChange(ctx context.Context, req *v1.AdministratorStatusChangeRequest) (*v1.CheckReply, error) {
	err := s.administratorUseCase.AdministratorStatusChange(ctx, req.Id, req.Status)
	isSuccess := false
	if err == nil {
		isSuccess = true
	}
	return &v1.CheckReply{
		IsSuccess: isSuccess,
	}, nil
}

func bizAdministratorToInfoReply(info *biz.Administrator) *v1.AdministratorInfoResponse {
	return &v1.AdministratorInfoResponse{
		Id:            info.Id,
		Username:      info.Username,
		Nickname:      info.Nickname,
		Mobile:        info.Mobile,
		Status:        info.Status,
		Avatar:        info.Avatar,
		Role:          info.Role,
		LastLoginTime: info.LastLoginTime,
		LastLoginIp:   info.LastLoginIp,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.CreatedAt,
		DeletedAt:     info.DeletedAt,
	}
}
