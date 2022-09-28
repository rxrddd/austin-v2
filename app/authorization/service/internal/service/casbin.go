package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
)

func (s *AuthorizationService) AddRolesForUser(ctx context.Context, req *v1.AddRolesForUserRequest) (*v1.CheckReply, error) {
	success, err := s.authorizationUsecase.AddRolesForUser(ctx, req.Username, req.Roles)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, nil
}

func (s *AuthorizationService) GetRolesForUser(ctx context.Context, req *v1.GetRolesForUserRequest) (*v1.GetRolesForUserReply, error) {
	roles, err := s.authorizationUsecase.GetRolesForUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.GetRolesForUserReply{
		Roles: roles,
	}, nil
}

func (s *AuthorizationService) GetUsersForRole(ctx context.Context, req *v1.GetUsersForRoleRequest) (*v1.GetUsersForRoleReply, error) {
	users, err := s.authorizationUsecase.GetUsersForRole(ctx, req.Role)
	if err != nil {
		return nil, err
	}
	return &v1.GetUsersForRoleReply{
		Users: users,
	}, nil
}

func (s *AuthorizationService) DeleteRoleForUser(ctx context.Context, req *v1.DeleteRoleForUserRequest) (*v1.CheckReply, error) {
	success, err := s.authorizationUsecase.DeleteRoleForUser(ctx, req.Username, req.Role)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, nil
}

func (s *AuthorizationService) DeleteRolesForUser(ctx context.Context, req *v1.DeleteRolesForUserRequest) (*v1.CheckReply, error) {
	success, err := s.authorizationUsecase.DeleteRolesForUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, nil
}

func (s *AuthorizationService) GetPolicies(ctx context.Context, req *v1.GetPoliciesRequest) (*v1.GetPoliciesReply, error) {

	rules, err := s.authorizationUsecase.GetPolicies(ctx, req.Role)
	if err != nil {
		return nil, err
	}
	reply := []*v1.PolicyRules{}
	for _, v := range rules {
		reply = append(reply, &v1.PolicyRules{
			Path:   v.Path,
			Method: v.Method,
		})
	}
	return &v1.GetPoliciesReply{
		PolicyRules: reply,
	}, nil
}

func (s *AuthorizationService) UpdatePolicies(ctx context.Context, req *v1.UpdatePoliciesRequest) (*v1.CheckReply, error) {
	rules := []biz.PolicyRules{}
	for _, v := range req.PolicyRules {
		rules = append(rules, biz.PolicyRules{
			Path:   v.Path,
			Method: v.Method,
		})
	}
	success, err := s.authorizationUsecase.UpdatePolicies(ctx, req.Role, rules)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: success,
	}, nil
}
