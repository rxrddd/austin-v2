package data

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/utils/convertHelper"
	"gorm.io/gorm"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

func (a AuthorizationRepo) GetRoleList(ctx context.Context) ([]*biz.Role, error) {
	var res []*biz.Role
	var roles []entity.Role
	// 获取所有根角色
	err := a.data.db.Model(entity.Role{}).Where("parent_id = 0").Find(&roles).Error
	if err != nil {
		return res, err
	}
	for _, v := range roles {
		res = append(res, &biz.Role{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			ParentIds: convertHelper.StringToInt64ArrayNoErr(v.ParentIds, "-"),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	for k := range res {
		err := a.findChildrenRole(res[k])
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func (a AuthorizationRepo) findChildrenRole(role *biz.Role) (err error) {
	var tmp []entity.Role
	err = a.data.db.Model(entity.Role{}).Where("parent_id = ?", role.Id).Find(&tmp).Error
	role.Children = []biz.Role{}
	for _, v := range tmp {
		role.Children = append(role.Children, biz.Role{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			ParentIds: convertHelper.StringToInt64ArrayNoErr(v.ParentIds, "-"),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	if len(role.Children) > 0 {
		for k := range role.Children {
			err = a.findChildrenRole(&role.Children[k])
		}
	}
	return err
}

func (a AuthorizationRepo) GetRole(ctx context.Context, params map[string]interface{}) (*biz.Role, error) {
	var res *biz.Role
	var role entity.Role

	db := a.data.db.Model(entity.Role{})
	if id, ok := params["id"]; ok && id != nil {
		db = db.Where("id = ?", id)
	}
	if name, ok := params["name"]; ok && name != nil {
		db = db.Where("name = ?", name)
	}
	err := db.First(&role).Error

	if err != nil {
		return res, err
	}
	res = &biz.Role{
		Id:        role.Id,
		Name:      role.Name,
		ParentId:  role.ParentId,
		ParentIds: convertHelper.StringToInt64ArrayNoErr(role.ParentIds, "-"),
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (a AuthorizationRepo) CreateRole(ctx context.Context, reqData *biz.Role) (*biz.Role, error) {
	var role entity.Role
	// 查看角色名是否存在
	err := a.data.db.Model(entity.Role{}).Where("name = ?", reqData.Name).First(&role).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Role{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationRoleExist)
	}
	now := time.Now()
	role = entity.Role{
		Name:      reqData.Name,
		ParentId:  reqData.ParentId,
		ParentIds: convertHelper.Int64ArrayToString(reqData.ParentIds, "-"),
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err = a.data.db.Model(entity.Role{}).Create(&role).Error
	if err != nil {
		return &biz.Role{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	res := &biz.Role{
		Id:        role.Id,
		Name:      role.Name,
		ParentId:  role.ParentId,
		ParentIds: convertHelper.StringToInt64ArrayNoErr(role.ParentIds, "-"),
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (a AuthorizationRepo) UpdateRole(ctx context.Context, reqData *biz.Role) (*biz.Role, error) {
	var role entity.Role
	// 查看角色名是否存在
	err := a.data.db.Model(entity.Role{}).Where("name = ? AND id != ?", reqData.Name, reqData.Id).First(&role).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Role{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationRoleExist)
	}
	err = a.data.db.Model(entity.Role{}).Where("id = ?", reqData.Id).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Role{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationRoleNotFound)
	}
	role.Name = reqData.Name
	role.ParentId = reqData.ParentId
	role.ParentIds = convertHelper.Int64ArrayToString(reqData.ParentIds, "-")
	err = a.data.db.Model(entity.Role{}).Where("id = ?", role.Id).Save(&role).Error
	if err != nil {
		return &biz.Role{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	a.data.db.Model(entity.Role{}).Where("id = ?", role.Id).Find(&role)
	res := &biz.Role{
		Id:        role.Id,
		Name:      role.Name,
		ParentId:  role.ParentId,
		ParentIds: convertHelper.StringToInt64ArrayNoErr(role.ParentIds, "-"),
		CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

// 检查角色是否存在
func (a AuthorizationRepo) checkRoleExist(role []string) bool {
	roleCount := int64(len(role))
	if roleCount == 0 {
		return false
	}
	var count int64
	if len(role) == 1 {
		a.data.db.Model(entity.Role{}).Where("name = ?", role[0]).Count(&count)
	} else {
		a.data.db.Model(entity.Role{}).Where("name IN (?)", role).Count(&count)
	}
	return count == roleCount
}

func (a AuthorizationRepo) DeleteRole(ctx context.Context, id int64) error {
	var role entity.Role
	// 查看角色是否存在
	err := a.data.db.Model(entity.Role{}).Where("id = ?", id).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errResponse.SetErrByReason(errResponse.ReasonAuthorizationRoleNotFound)
	}
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	users, err := a.data.enforcer.GetUsersForRole(role.Name)
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if len(users) > 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "角色已被使用,无法删除")
	}
	tx := a.data.db.Begin()
	// 删除角色
	err = tx.Model(entity.Role{}).Where("id = ?", id).Delete(&role).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 删除角色关联的菜单
	err = tx.Where("role_id = ?", id).Delete(&entity.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	tx.Commit()
	// 删除策略
	a.data.enforcer.RemoveFilteredPolicy(0, role.Name)
	return nil
}
