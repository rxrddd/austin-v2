package biz

import (
	"context"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

func (uc *AuthorizationUsecase) SetRolesForUser(ctx context.Context, username string, roles []string) (bool, error) {
	return uc.repo.SetRolesForUser(ctx, username, roles)
}

func (uc *AuthorizationUsecase) AddRolesForUser(ctx context.Context, username string, roles []string) (bool, error) {
	return uc.repo.AddRolesForUser(ctx, username, roles)
}

func (uc *AuthorizationUsecase) GetRolesForUser(ctx context.Context, username string) ([]string, error) {
	if username == "" {
		return []string{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingParams)
	}
	return uc.repo.GetRolesForUser(ctx, username)
}

func (uc *AuthorizationUsecase) GetUsersForRole(ctx context.Context, role string) ([]string, error) {
	if role == "" {
		return []string{}, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingParams)
	}
	return uc.repo.GetUsersForRole(ctx, role)
}

func (uc *AuthorizationUsecase) DeleteRoleForUser(ctx context.Context, username string, role string) (bool, error) {
	if role == "" || username == "" {
		return false, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingParams)
	}
	return uc.repo.DeleteRoleForUser(ctx, username, role)
}

func (uc *AuthorizationUsecase) DeleteRolesForUser(ctx context.Context, username string) (bool, error) {
	if username == "" {
		return false, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonMissingParams)
	}
	return uc.repo.DeleteRolesForUser(ctx, username)
}

type PolicyRules struct {
	Path   string
	Method string
}

func (uc *AuthorizationUsecase) GetPolicies(ctx context.Context, role string) ([]*PolicyRules, error) {
	return uc.repo.GetPolicies(ctx, role)
}

func (uc *AuthorizationUsecase) UpdatePolicies(ctx context.Context, role string, rules []PolicyRules) (bool, error) {
	if role == "" {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, "角色不得为空")
	}
	return uc.repo.UpdatePolicies(ctx, role, rules)
}
