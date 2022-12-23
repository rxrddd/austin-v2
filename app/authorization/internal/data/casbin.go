package data

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/biz"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"strings"
)

func (a AuthorizationRepo) SetRolesForUser(ctx context.Context, username string, roles []string) (bool, error) {
	// 检查角色是否存在
	if !a.checkRoleExist(roles) {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, "角色不存在")
	}

	// 删除用户所有角色
	success, err := a.data.enforcer.DeleteRolesForUser(username)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 添加用户角色
	success, err = a.data.enforcer.AddRolesForUser(username, roles)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	return success, nil
}

func (a AuthorizationRepo) AddRolesForUser(ctx context.Context, username string, roles []string) (bool, error) {
	// 检查角色是否存在
	if !a.checkRoleExist(roles) {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, "角色不存在")
	}

	// 检查用户是否已经拥有角色
	for _, v := range roles {
		exist, err := a.data.enforcer.HasRoleForUser(username, v)
		if err != nil {
			return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
		if exist {
			return false, kerrors.BadRequest(errResponse.ReasonParamsError, username+" 已拥有 "+v+" 角色")
		}
	}

	success, err := a.data.enforcer.AddRolesForUser(username, roles)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	return success, nil
}

func (a AuthorizationRepo) GetRolesForUser(ctx context.Context, username string) ([]string, error) {
	return a.data.enforcer.GetRolesForUser(username)
}

func (a AuthorizationRepo) GetUsersForRole(ctx context.Context, role string) ([]string, error) {
	return a.data.enforcer.GetUsersForRole(role)
}

func (a AuthorizationRepo) DeleteRoleForUser(ctx context.Context, username string, role string) (bool, error) {
	exist, err := a.data.enforcer.HasRoleForUser(username, role)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if !exist {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, username+" 未拥有 "+role+" 角色")
	}
	return a.data.enforcer.DeleteRoleForUser(username, role)
}

func (a AuthorizationRepo) DeleteRolesForUser(ctx context.Context, username string) (bool, error) {
	return a.data.enforcer.DeleteRolesForUser(username)
}

func (a AuthorizationRepo) GetPolicies(ctx context.Context, role string) ([]*biz.PolicyRules, error) {
	rules := []*biz.PolicyRules{}

	// 检查角色是否存在
	if !a.checkRoleExist([]string{role}) {
		return rules, kerrors.BadRequest(errResponse.ReasonParamsError, "角色不存在")
	}

	// 查询已有策略规则
	policies := a.data.enforcer.GetFilteredPolicy(0, role)
	for _, v := range policies {
		rules = append(rules, &biz.PolicyRules{
			Path:   v[1],
			Method: v[2],
		})
	}

	return rules, nil
}

func (a AuthorizationRepo) UpdatePolicies(ctx context.Context, role string, rules []biz.PolicyRules) (bool, error) {
	// 检查角色是否存在
	if !a.checkRoleExist([]string{role}) {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, "角色不存在")
	}
	policies := [][]string{}
	for _, v := range rules {
		// method需要为全部大写
		policies = append(policies, []string{role, v.Path, strings.ToUpper(v.Method)})
	}
	// 移除已有策略规则
	_, err := a.data.enforcer.RemoveFilteredPolicy(0, role)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	success, err := a.data.enforcer.AddPolicies(policies)
	if err != nil {
		return false, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if !success {
		return false, kerrors.BadRequest(errResponse.ReasonParamsError, "存在相同api,添加失败")
	}
	return true, nil
}
