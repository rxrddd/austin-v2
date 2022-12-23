package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthorizationService) GetRoleList(ctx context.Context, req *emptypb.Empty) (*v1.GetRoleListReply, error) {

	role, err := s.authorizationUsecase.GetRoleList(ctx)
	if err != nil {
		return nil, err
	}
	list := []*v1.RoleInfo{}
	for k := range role {
		children := findChildrenRole(role[k])
		res := &v1.RoleInfo{
			Id:        role[k].Id,
			ParentId:  role[k].ParentId,
			ParentIds: role[k].ParentIds,
			Name:      role[k].Name,
			CreatedAt: role[k].CreatedAt,
			UpdatedAt: role[k].UpdatedAt,
			Children:  children,
		}
		list = append(list, res)
	}

	return &v1.GetRoleListReply{
		List: list,
	}, nil

}

func findChildrenRole(role *biz.Role) []*v1.RoleInfo {
	children := []*v1.RoleInfo{}
	if len(role.Children) != 0 {
		for k := range role.Children {
			children = append(children, &v1.RoleInfo{
				Id:        role.Children[k].Id,
				Name:      role.Children[k].Name,
				ParentId:  role.Children[k].ParentId,
				ParentIds: role.Children[k].ParentIds,
				CreatedAt: role.Children[k].CreatedAt,
				UpdatedAt: role.Children[k].UpdatedAt,
				Children:  findChildrenRole(&role.Children[k]),
			})
		}
	}
	return children
}

func (s *AuthorizationService) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.RoleInfo, error) {
	bc := &biz.Role{
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
	}

	role, err := s.authorizationUsecase.CreateRole(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.RoleInfo{
		Id:        role.Id,
		ParentId:  role.ParentId,
		ParentIds: role.ParentIds,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}, nil
}

func (s *AuthorizationService) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.RoleInfo, error) {
	bc := &biz.Role{
		Id:        req.Id,
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
	}

	role, err := s.authorizationUsecase.UpdateRole(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.RoleInfo{
		Id:        role.Id,
		ParentId:  role.ParentId,
		ParentIds: role.ParentIds,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}, nil
}

func (s *AuthorizationService) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.CheckReply, error) {
	err := s.authorizationUsecase.DeleteRole(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: true,
	}, nil
}
