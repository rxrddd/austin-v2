package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthorizationService) GetApiAll(ctx context.Context, empty *emptypb.Empty) (*v1.GetApiAllReply, error) {
	res, err := s.authorizationUsecase.GetApiAll(ctx)
	if err != nil {
		return nil, err
	}
	list := []*v1.ApiInfo{}
	for k := range res {
		res := &v1.ApiInfo{
			Id:        res[k].Id,
			Group:     res[k].Group,
			Name:      res[k].Name,
			Path:      res[k].Path,
			Method:    res[k].Method,
			CreatedAt: res[k].CreatedAt,
			UpdatedAt: res[k].UpdatedAt,
		}
		list = append(list, res)
	}

	return &v1.GetApiAllReply{
		List: list,
	}, nil

}

func (s *AuthorizationService) GetApiList(ctx context.Context, req *v1.GetApiListRequest) (*v1.GetApiListReply, error) {
	params := map[string]interface{}{}
	params["group"] = req.Group
	params["name"] = req.Name
	params["path"] = req.Path
	params["method"] = req.Method

	res, count, err := s.authorizationUsecase.GetApiList(ctx, req.Page, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	response := &v1.GetApiListReply{}
	response.Total = count
	for k := range res {
		res := &v1.ApiInfo{
			Id:        res[k].Id,
			Group:     res[k].Group,
			Name:      res[k].Name,
			Path:      res[k].Path,
			Method:    res[k].Method,
			CreatedAt: res[k].CreatedAt,
			UpdatedAt: res[k].UpdatedAt,
		}
		response.List = append(response.List, res)
	}
	return response, nil

}

func (s *AuthorizationService) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.ApiInfo, error) {
	bc := &biz.Api{
		Group:  req.Group,
		Name:   req.Name,
		Path:   req.Path,
		Method: req.Method,
	}

	Api, err := s.authorizationUsecase.CreateApi(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ApiInfo{
		Id:        Api.Id,
		Group:     Api.Group,
		Name:      Api.Name,
		Path:      Api.Path,
		Method:    Api.Method,
		CreatedAt: Api.CreatedAt,
		UpdatedAt: Api.UpdatedAt,
	}, nil
}

func (s *AuthorizationService) UpdateApi(ctx context.Context, req *v1.UpdateApiRequest) (*v1.ApiInfo, error) {
	bc := &biz.Api{
		Id:     req.Id,
		Group:  req.Group,
		Name:   req.Name,
		Path:   req.Path,
		Method: req.Method,
	}

	Api, err := s.authorizationUsecase.UpdateApi(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ApiInfo{
		Id:        Api.Id,
		Group:     Api.Group,
		Name:      Api.Name,
		Path:      Api.Path,
		Method:    Api.Method,
		CreatedAt: Api.CreatedAt,
		UpdatedAt: Api.UpdatedAt,
	}, nil
}

func (s *AuthorizationService) DeleteApi(ctx context.Context, req *v1.DeleteApiRequest) (*v1.CheckReply, error) {
	err := s.authorizationUsecase.DeleteApi(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: true,
	}, nil
}
