package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/biz"

	v1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
)

func (s *AdministratorService) GetLoginAdministratorByUsername(ctx context.Context, req *v1.GetLoginAdministratorByUsernameRequest) (*v1.GetLoginAdministratorByUsernameReply, error) {
	return s.administratorCase.FindLoginAdministratorByUsername(ctx, req)
}

func (s *AdministratorService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordReply, error) {
	res, err := s.administratorCase.VerifyAdministratorPassword(ctx, req)
	return &v1.VerifyPasswordReply{
		Success: res,
	}, err
}

func (s *AdministratorService) CreateAdministrator(ctx context.Context, req *v1.CreateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	bc := &biz.Administrator{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	}
	AdministratorInfo, err := s.administratorCase.Create(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.AdministratorInfoResponse{
		Id:        AdministratorInfo.Id,
		Username:  AdministratorInfo.Username,
		Mobile:    AdministratorInfo.Mobile,
		Avatar:    AdministratorInfo.Avatar,
		Nickname:  AdministratorInfo.Nickname,
		Status:    AdministratorInfo.Status,
		CreatedAt: AdministratorInfo.CreatedAt,
		UpdatedAt: AdministratorInfo.UpdatedAt,
		DeletedAt: AdministratorInfo.DeletedAt,
	}, nil
}
func (s *AdministratorService) UpdateAdministrator(ctx context.Context, req *v1.UpdateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	bc := &biz.Administrator{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   req.Status,
	}
	AdministratorInfo, err := s.administratorCase.Update(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.AdministratorInfoResponse{
		Id:        AdministratorInfo.Id,
		Username:  AdministratorInfo.Username,
		Mobile:    AdministratorInfo.Mobile,
		Avatar:    AdministratorInfo.Avatar,
		Nickname:  AdministratorInfo.Nickname,
		Status:    AdministratorInfo.Status,
		CreatedAt: AdministratorInfo.CreatedAt,
		UpdatedAt: AdministratorInfo.UpdatedAt,
		DeletedAt: AdministratorInfo.DeletedAt,
	}, nil
}

func (s *AdministratorService) DeleteAdministrator(ctx context.Context, req *v1.DeleteAdministratorRequest) (*v1.AdministratorCheckResponse, error) {
	err := s.administratorCase.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.AdministratorCheckResponse{
		IsSuccess: success,
	}, err
}

func (s *AdministratorService) GetAdministrator(ctx context.Context, req *v1.GetAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	params := map[string]interface{}{}
	params["id"] = req.Id
	params["username"] = req.Username
	params["mobile"] = req.Mobile
	AdministratorInfo, err := s.administratorCase.Get(ctx, params)
	if err != nil {
		return nil, err
	}
	response := &v1.AdministratorInfoResponse{
		Id:        AdministratorInfo.Id,
		Username:  AdministratorInfo.Username,
		Mobile:    AdministratorInfo.Mobile,
		Avatar:    AdministratorInfo.Avatar,
		Nickname:  AdministratorInfo.Nickname,
		Status:    AdministratorInfo.Status,
		CreatedAt: AdministratorInfo.CreatedAt,
		UpdatedAt: AdministratorInfo.UpdatedAt,
		DeletedAt: AdministratorInfo.DeletedAt,
	}
	return response, nil
}

func (s *AdministratorService) ListAdministrator(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	AdministratorInfoList, count, err := s.administratorCase.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	response := &v1.ListAdministratorReply{}
	response.Total = count
	for _, v := range AdministratorInfoList {
		administratorInfo := &v1.AdministratorInfoResponse{
			Id:        v.Id,
			Username:  v.Username,
			Mobile:    v.Mobile,
			Avatar:    v.Avatar,
			Nickname:  v.Nickname,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
		response.List = append(response.List, administratorInfo)
	}
	return response, nil
}
