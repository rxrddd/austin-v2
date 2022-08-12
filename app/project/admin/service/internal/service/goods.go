package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/biz"

	v1 "github.com/ZQCard/kratos-base-project/api/project/admin/v1"
)

func (s *AdminInterface) CreateGoods(ctx context.Context, req *v1.CreateGoodsRequest) (*v1.GoodsInfoResponse, error) {
	bg := &biz.Goods{
		Name:   req.Name,
		Style:  req.Style,
		Size:   req.Size,
		Code:   req.Code,
		Price:  req.Price,
		Number: req.Number,
	}
	GoodsInfo, err := s.GoodsUseCase.Create(ctx, bg)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsInfoResponse{
		Id:        GoodsInfo.Id,
		Name:      GoodsInfo.Name,
		Style:     GoodsInfo.Style,
		Size:      GoodsInfo.Size,
		Code:      GoodsInfo.Code,
		Price:     GoodsInfo.Price,
		Number:    GoodsInfo.Number,
		CreatedAt: GoodsInfo.CreatedAt,
		UpdatedAt: GoodsInfo.UpdatedAt,
		DeletedAt: GoodsInfo.DeletedAt,
	}, nil
}

func (s *AdminInterface) UpdateGoods(ctx context.Context, req *v1.UpdateGoodsRequest) (*v1.GoodsInfoResponse, error) {
	bg := &biz.Goods{
		Id:     req.Id,
		Name:   req.Name,
		Style:  req.Style,
		Size:   req.Size,
		Code:   req.Code,
		Price:  req.Price,
		Number: req.Number,
	}
	GoodsInfo, err := s.GoodsUseCase.Update(ctx, bg)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsInfoResponse{
		Id:        GoodsInfo.Id,
		Name:      GoodsInfo.Name,
		Style:     GoodsInfo.Style,
		Size:      GoodsInfo.Size,
		Code:      GoodsInfo.Code,
		Price:     GoodsInfo.Price,
		Number:    GoodsInfo.Number,
		CreatedAt: GoodsInfo.CreatedAt,
		UpdatedAt: GoodsInfo.UpdatedAt,
		DeletedAt: GoodsInfo.DeletedAt,
	}, nil
}

func (s *AdminInterface) GetGoods(ctx context.Context, req *v1.GetGoodsRequest) (*v1.GoodsInfoResponse, error) {
	GoodsInfo, err := s.GoodsUseCase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	response := &v1.GoodsInfoResponse{
		Id:        GoodsInfo.Id,
		Name:      GoodsInfo.Name,
		Style:     GoodsInfo.Style,
		Size:      GoodsInfo.Size,
		Code:      GoodsInfo.Code,
		Price:     GoodsInfo.Price,
		Number:    GoodsInfo.Number,
		CreatedAt: GoodsInfo.CreatedAt,
		UpdatedAt: GoodsInfo.UpdatedAt,
		DeletedAt: GoodsInfo.DeletedAt,
	}
	return response, nil
}

func (s *AdminInterface) DeleteGoods(ctx context.Context, req *v1.DeleteGoodsRequest) (*v1.GoodsCheckResponse, error) {
	err := s.GoodsUseCase.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.GoodsCheckResponse{
		IsSuccess: success,
	}, err
}

func (s *AdminInterface) ListGoods(ctx context.Context, req *v1.ListGoodsRequest) (*v1.ListGoodsReply, error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	bg := &biz.Goods{
		Id:        req.Id,
		Name:      req.Name,
		Style:     req.Style,
		Size:      req.Size,
		Code:      req.Code,
		DeletedAt: req.DeletedAt,
	}
	GoodsInfoList, count, err := s.GoodsUseCase.List(ctx, req.PageNum, req.PageSize, bg)
	if err != nil {
		return nil, err
	}
	response := &v1.ListGoodsReply{}
	response.Total = count
	for _, v := range GoodsInfoList {
		response.List = append(response.List, &v1.GoodsInfoResponse{
			Id:        v.Id,
			Name:      v.Name,
			Style:     v.Style,
			Size:      v.Size,
			Code:      v.Code,
			Price:     v.Price,
			Number:    v.Number,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		})
	}
	return response, nil
}

func (s *AdminInterface) SaleGoods(ctx context.Context, req *v1.SaleGoodsRequest) (*v1.SaleGoodsReply, error) {
	err := s.GoodsUseCase.Sale(ctx, req.Id, req.Number)
	success := true
	if err != nil {
		success = false
	}
	return &v1.SaleGoodsReply{
		IsSuccess: success,
	}, err
}

func (s *AdminInterface) RecoverGoods(ctx context.Context, req *v1.RecoverGoodsRequest) (*v1.GoodsCheckResponse, error) {
	err := s.GoodsUseCase.Recover(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.GoodsCheckResponse{
		IsSuccess: success,
	}, err
}
