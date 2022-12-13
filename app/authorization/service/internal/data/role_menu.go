package data

import (
	"context"
	"encoding/json"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/utils/redisHelper"
	kerrors "github.com/go-kratos/kratos/v2/errors"
)

const childModuleRoleMenu = "RoleMenu"

func (a AuthorizationRepo) GetRoleMenuTree(ctx context.Context, role string) ([]*biz.Menu, error) {
	// 查询角色拥有哪些菜单
	var response []*biz.Menu
	// 缓存key
	cacheParams := map[string]interface{}{
		"role": role,
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleRoleMenu+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, nil
		} else {
			a.log.Error("GetRoleMenuTree()", err)
		}
	}

	var menus []entity.Menu
	tmpRole, err := a.GetRole(ctx, map[string]interface{}{
		"name": role,
	})
	if tmpRole.Id == 0 {
		return response, nil
	}
	// 查看角色拥有菜单id
	menuIds := a.getMenuIdsByRoleId(tmpRole.Id)
	if len(menuIds) == 0 {
		return response, nil
	}

	// 获取所有根菜单
	err = a.data.db.Model(entity.Menu{}).Where("parent_id = 0 AND id IN (?)", menuIds).Preload("MenuBtns").Order("sort ASC").Find(&menus).Error
	if err != nil {
		return response, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	for _, v := range menus {
		response = append(response, &biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			Path:      v.Path,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	for k := range response {
		err := a.findChildrenRoleMenu(response[k], menuIds)
		if err != nil {
			return response, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
	}
	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(a.data.redisCli, cacheKey, string(jsonResponse))
	return response, nil
}

func (a AuthorizationRepo) findChildrenRoleMenu(menu *biz.Menu, menuIds []int64) (err error) {
	var tmp []entity.Menu
	err = a.data.db.Model(entity.Menu{}).Where("parent_id = ? AND id IN (?)", menu.Id, menuIds).Preload("MenuBtns").Find(&tmp).Error
	menu.Children = []biz.Menu{}
	for _, v := range tmp {
		btns := []biz.MenuBtn{}
		for _, btn := range v.MenuBtns {
			btns = append(btns, biz.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				CreatedAt:   btn.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   btn.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		menu.Children = append(menu.Children, biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			Path:      v.Path,
			ParentId:  v.ParentId,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
			MenuBtns:  btns,
		})
	}
	if len(menu.Children) > 0 {
		for k := range menu.Children {
			err = a.findChildrenRoleMenu(&menu.Children[k], menuIds)
		}
	}
	return err
}

func (a AuthorizationRepo) GetRoleMenu(ctx context.Context, role string) ([]*biz.Menu, error) {
	// 查询角色拥有哪些菜单
	var response []*biz.Menu
	// 缓存key
	cacheParams := map[string]interface{}{
		"type": "GetRoleMenu",
		"role": role,
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleRoleMenu+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, nil
		} else {
			a.log.Error("GetRoleMenuTree()", err)
		}
	}

	var menus []entity.Menu
	tmpRole, err := a.GetRole(ctx, map[string]interface{}{
		"name": role,
	})
	if tmpRole.Id == 0 {
		return response, nil
	}
	// 查看角色拥有菜单id
	menuIds := a.getMenuIdsByRoleId(tmpRole.Id)
	if len(menuIds) == 0 {
		return response, nil
	}

	// 获取所有根菜单
	err = a.data.db.Model(entity.Menu{}).Where("id IN (?) ", menuIds).Preload("MenuBtns").Find(&menus).Error
	if err != nil {
		return response, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	for _, v := range menus {
		btns := []biz.MenuBtn{}
		for _, btn := range v.MenuBtns {
			btns = append(btns, biz.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				CreatedAt:   btn.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   btn.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		response = append(response, &biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
			MenuBtns:  btns,
		})
	}
	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(a.data.redisCli, cacheKey, string(jsonResponse))
	return response, nil
}

func (a AuthorizationRepo) SaveRoleMenu(ctx context.Context, roleId int64, menuIds []int64) error {
	tx := a.data.db.Begin()
	// 先删除数据
	err := tx.Where("role_id = ?", roleId).Delete(&entity.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if len(menuIds) == 0 {
		tx.Commit()
		return nil
	}
	// 批量插入数据
	roleMenu := []entity.RoleMenu{}
	for _, v := range menuIds {
		roleMenu = append(roleMenu, entity.RoleMenu{
			RoleId: roleId,
			MenuId: v,
		})
	}
	if err := tx.Create(&roleMenu).Error; err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	tx.Commit()
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleRoleMenu+":")
	return nil
}

func (a AuthorizationRepo) getMenuIdsByRoleId(roleId int64) (menuIds []int64) {
	// 查询角色拥有哪些菜单
	var roleMenu []entity.RoleMenu

	a.data.db.Model(entity.RoleMenu{}).Where("role_id = ?", roleId).Find(&roleMenu)
	// 查看角色拥有菜单id
	for _, v := range roleMenu {
		menuIds = append(menuIds, v.MenuId)
	}
	return menuIds
}

func (a AuthorizationRepo) GetRoleMenuBtn(ctx context.Context, roleId int64, roleName string, menuId int64) (response []*biz.MenuBtn, err error) {
	// 如果角色名称不为空， 则根据名称查找角色id
	if roleName != "" {
		roleInfo, err := a.GetRole(ctx, map[string]interface{}{
			"name": roleName,
		})
		if err != nil {
			return nil, err
		}

		if roleId != 0 && roleInfo.Id != roleId {
			return nil, kerrors.BadRequest(errResponse.ReasonParamsError, "角色参数错误")
		}
		roleId = roleInfo.Id
	}
	// 查询角色拥有哪些菜单
	// 缓存key
	cacheParams := map[string]interface{}{
		"roleId":   roleId,
		"roleName": roleName,
		"menuId":   menuId,
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleRoleMenu+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, nil
		} else {
			a.log.Error("GetRoleMenuTree()", err)
		}
	}

	// 查询角色拥有哪些菜单按钮
	var roleMenuBtn []entity.RoleMenuBtn
	conn := a.data.db.Model(entity.RoleMenuBtn{})
	if menuId != 0 {
		conn = conn.Where("menu_id = ?", menuId)
	}
	if roleId != 0 {
		conn = conn.Where("role_id = ?", roleId)
	}
	err = conn.Find(&roleMenuBtn).Error
	var btnIds []int64
	// 查看角色拥有菜单id
	for _, v := range roleMenuBtn {
		btnIds = append(btnIds, v.BtnId)
	}

	btnResponse := []*entity.MenuBtn{}
	if len(btnIds) == 0 {
		return []*biz.MenuBtn{}, nil
	}
	a.data.db.Model(&entity.MenuBtn{}).Where("id IN (?)", btnIds).Find(&btnResponse)
	for _, v := range btnResponse {
		response = append(response, &biz.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(a.data.redisCli, cacheKey, string(jsonResponse))
	return response, err
}

func (a AuthorizationRepo) SetRoleMenuBtn(ctx context.Context, roleId int64, menuId int64, btnIds []int64) error {
	tx := a.data.db.Begin()
	// 先删除数据
	err := tx.Where("role_id = ? AND menu_id = ?", roleId, menuId).Delete(&entity.RoleMenuBtn{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if len(btnIds) == 0 {
		tx.Commit()
		return nil
	}
	// 批量插入数据
	roleMenuBtn := []entity.RoleMenuBtn{}
	for _, v := range btnIds {
		roleMenuBtn = append(roleMenuBtn, entity.RoleMenuBtn{
			RoleId: roleId,
			MenuId: menuId,
			BtnId:  v,
		})
	}
	if err := tx.Create(&roleMenuBtn).Error; err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	tx.Commit()
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleRoleMenu+":")
	return nil
}
