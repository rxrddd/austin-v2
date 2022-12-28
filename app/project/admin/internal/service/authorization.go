package service

import (
	"austin-v2/api/project/admin/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminInterface) GetRoleList(ctx context.Context, req *emptypb.Empty) (*v1.GetRoleListReply, error) {
	return s.authorizationRepo.GetRoleList(ctx)
}

func (s *AdminInterface) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.RoleInfo, error) {
	return s.authorizationRepo.CreateRole(ctx, req)
}
func (s *AdminInterface) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.RoleInfo, error) {
	return s.authorizationRepo.UpdateRole(ctx, req)
}
func (s *AdminInterface) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.DeleteRole(ctx, req)
}
func (s *AdminInterface) SetRolesForUser(ctx context.Context, req *v1.SetRolesForUserRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.SetRolesForUser(ctx, req)
}
func (s *AdminInterface) GetRolesForUser(ctx context.Context, req *v1.GetRolesForUserRequest) (*v1.GetRolesForUserReply, error) {
	return s.authorizationRepo.GetRolesForUser(ctx, req)
}
func (s *AdminInterface) GetUsersForRole(ctx context.Context, req *v1.GetUsersForRoleRequest) (*v1.GetUsersForRoleReply, error) {
	return s.authorizationRepo.GetUsersForRole(ctx, req)
}

func (s *AdminInterface) DeleteRoleForUser(ctx context.Context, req *v1.DeleteRoleForUserRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.DeleteRoleForUser(ctx, req)
}

func (s *AdminInterface) GetPolicies(ctx context.Context, req *v1.GetPoliciesRequest) (*v1.GetPoliciesReply, error) {
	return s.authorizationRepo.GetPolicies(ctx, req)
}
func (s *AdminInterface) UpdatePolicies(ctx context.Context, req *v1.UpdatePoliciesRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.UpdatePolicies(ctx, req)
}
func (s *AdminInterface) GetApiAll(ctx context.Context, req *emptypb.Empty) (*v1.GetApiAllReply, error) {
	return s.authorizationRepo.GetApiAll(ctx)
}
func (s *AdminInterface) GetApiList(ctx context.Context, req *v1.GetApiListRequest) (*v1.GetApiListReply, error) {
	return s.authorizationRepo.GetApiList(ctx, req)
}
func (s *AdminInterface) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.ApiInfo, error) {
	return s.authorizationRepo.CreateApi(ctx, req)
}
func (s *AdminInterface) UpdateApi(ctx context.Context, req *v1.UpdateApiRequest) (*v1.ApiInfo, error) {
	return s.authorizationRepo.UpdateApi(ctx, req)
}
func (s *AdminInterface) DeleteApi(ctx context.Context, req *v1.DeleteApiRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.DeleteApi(ctx, req)
}
func (s *AdminInterface) GetMenuAll(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	return s.authorizationRepo.GetMenuAll(ctx)
}
func (s *AdminInterface) GetMenuTree(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	return s.authorizationRepo.GetMenuTree(ctx)
}
func (s *AdminInterface) CreateMenu(ctx context.Context, req *v1.CreateMenuRequest) (*v1.MenuInfo, error) {
	return s.authorizationRepo.CreateMenu(ctx, req)
}
func (s *AdminInterface) UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest) (*v1.MenuInfo, error) {
	return s.authorizationRepo.UpdateMenu(ctx, req)
}
func (s *AdminInterface) DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.DeleteMenu(ctx, req)
}
func (s *AdminInterface) GetRoleMenuTree(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {
	return s.authorizationRepo.GetRoleMenuTree(ctx, req)
}
func (s *AdminInterface) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {
	return s.authorizationRepo.GetRoleMenu(ctx, req)
}
func (s *AdminInterface) SetRoleMenu(ctx context.Context, req *v1.SetRoleMenuRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.SetRoleMenu(ctx, req)
}

func (s *AdminInterface) GetRoleMenuBtn(ctx context.Context, req *v1.GetRoleMenuBtnRequest) (*v1.GetRoleMenuBtnReply, error) {
	return s.authorizationRepo.GetRoleMenuBtn(ctx, req)
}
func (s *AdminInterface) SetRoleMenuBtn(ctx context.Context, req *v1.SetRoleMenuBtnRequest) (*v1.CheckReply, error) {
	return s.authorizationRepo.SetRoleMenuBtn(ctx, req)
}
