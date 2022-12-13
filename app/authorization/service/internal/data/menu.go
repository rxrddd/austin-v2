package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/utils/convertHelper"
	"github.com/ZQCard/kratos-base-project/pkg/utils/redisHelper"
	"gorm.io/gorm"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

const childModuleMenu = "Menu"

func (a AuthorizationRepo) GetMenuAll(ctx context.Context) ([]*biz.Menu, error) {
	var response []*biz.Menu
	// 缓存key
	cacheParams := map[string]interface{}{
		"type": "all",
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleMenu+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, nil
		} else {
			a.log.Error("GetMenuAll()", err)
		}
	}

	var menus []entity.Menu
	// 获取所有根菜单
	err := a.data.db.Model(entity.Menu{}).Preload("MenuBtns").Find(&menus).Error
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
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   btn.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		response = append(response, &biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			ParentIds: convertHelper.StringToInt64ArrayNoErr(v.ParentIds, "-"),
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

func (a AuthorizationRepo) GetMenuTree(ctx context.Context) ([]*biz.Menu, error) {
	var response []*biz.Menu
	// 缓存key
	cacheParams := map[string]interface{}{
		"type": "tree",
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleMenu+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, nil
		} else {
			a.log.Error("GetMenuTree()", err)
		}
	}

	var menus []entity.Menu
	// 获取所有根菜单
	err := a.data.db.Model(entity.Menu{}).Where("parent_id = 0").Preload("MenuBtns").Order("sort ASC").Find(&menus).Error
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
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   btn.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		response = append(response, &biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			ParentId:  v.ParentId,
			ParentIds: convertHelper.StringToInt64ArrayNoErr(v.ParentIds, "-"),
			Path:      v.Path,
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
	for k := range response {
		err := a.findChildrenMenu(response[k])
		if err != nil {
			return response, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
	}
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleMenu+":")
	return response, nil
}

func (a AuthorizationRepo) findChildrenMenu(menu *biz.Menu) (err error) {
	var tmp []entity.Menu
	err = a.data.db.Model(entity.Menu{}).Where("parent_id = ?", menu.Id).Preload("MenuBtns").Find(&tmp).Error
	menu.Children = []biz.Menu{}
	for _, v := range tmp {
		btns := []biz.MenuBtn{}
		for _, btn := range v.MenuBtns {
			btns = append(btns, biz.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   btn.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		menu.Children = append(menu.Children, biz.Menu{
			Id:        v.Id,
			Name:      v.Name,
			Path:      v.Path,
			ParentId:  v.ParentId,
			ParentIds: convertHelper.StringToInt64ArrayNoErr(v.ParentIds, "-"),
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
			err = a.findChildrenMenu(&menu.Children[k])
		}
	}
	return err
}

func (a AuthorizationRepo) CreateMenu(ctx context.Context, reqData *biz.Menu) (*biz.Menu, error) {
	btns := []*entity.MenuBtn{}
	for _, v := range reqData.MenuBtns {
		btns = append(btns, &entity.MenuBtn{
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
		})
	}
	var menu entity.Menu
	now := time.Now()
	menu = entity.Menu{
		Name:      reqData.Name,
		ParentId:  reqData.ParentId,
		ParentIds: convertHelper.Int64ArrayToString(reqData.ParentIds, "-"),
		Path:      reqData.Path,
		Hidden:    reqData.Hidden,
		Component: reqData.Component,
		Sort:      reqData.Sort,
		Title:     reqData.Title,
		Icon:      reqData.Icon,
		CreatedAt: &now,
		UpdatedAt: &now,
		MenuBtns:  btns,
	}
	err := a.data.db.Model(entity.Menu{}).Create(&menu).Error
	if err != nil {
		return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	btns2 := []biz.MenuBtn{}
	for _, v := range menu.MenuBtns {
		btns2 = append(btns2, biz.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	res := &biz.Menu{
		Id:        menu.Id,
		Name:      menu.Name,
		Path:      menu.Path,
		ParentId:  menu.ParentId,
		ParentIds: convertHelper.StringToInt64ArrayNoErr(menu.ParentIds, "-"),
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
		MenuBtns:  btns2,
	}
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleMenu+":")
	return res, nil
}

func (a AuthorizationRepo) UpdateMenu(ctx context.Context, reqData *biz.Menu) (*biz.Menu, error) {
	btns := []*entity.MenuBtn{}
	for _, v := range reqData.MenuBtns {
		btns = append(btns, &entity.MenuBtn{
			Id:          v.Id,
			MenuId:      reqData.Id,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
		})
	}
	var menu entity.Menu
	a.data.db.Model(entity.Menu{}).Where("id = ?", reqData.Id).First(&menu)
	menu.Id = reqData.Id
	menu.Name = reqData.Name
	menu.ParentId = reqData.ParentId
	menu.ParentIds = convertHelper.Int64ArrayToString(reqData.ParentIds, "-")
	menu.Path = reqData.Path
	menu.Hidden = reqData.Hidden
	menu.Component = reqData.Component
	menu.Sort = reqData.Sort
	menu.Title = reqData.Title
	menu.Icon = reqData.Icon
	menu.MenuBtns = btns
	// 关联数据更新
	tx := a.data.db.Begin()
	err := tx.Model(entity.Menu{}).Where("id = ?", menu.Id).Session(&gorm.Session{FullSaveAssociations: true}).Save(&menu).Error
	if err != nil {
		tx.Rollback()
		return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 先删除,后添加
	if err = tx.Where("menu_id  = ?", menu.Id).Unscoped().Delete(&entity.MenuBtn{}).Error; err != nil {
		tx.Rollback()
		return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 保存按钮
	for _, v := range menu.MenuBtns {
		if err = tx.Model(entity.MenuBtn{}).Where("id = ?", v.Id).Create(&v).Error; err != nil {
			tx.Rollback()
			return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
	}
	tx.Commit()
	a.data.db.Model(entity.Menu{}).Where("id = ?", menu.Id).Preload("MenuBtns").Find(&menu)
	btns2 := []biz.MenuBtn{}
	for _, v := range menu.MenuBtns {
		btns2 = append(btns2, biz.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	res := &biz.Menu{
		Id:        menu.Id,
		Name:      menu.Name,
		Path:      menu.Path,
		ParentId:  menu.ParentId,
		ParentIds: convertHelper.StringToInt64ArrayNoErr(menu.ParentIds, "-"),
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		MenuBtns:  btns2,
		CreatedAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleMenu+":")
	return res, nil
}

func (a AuthorizationRepo) DeleteMenu(ctx context.Context, id int64) error {
	var menu entity.Menu
	// 查看菜单是否存在
	err := a.data.db.Model(entity.Menu{}).Where("id = ?", id).First(&menu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errResponse.SetErrByReason(errResponse.ReasonRecordNotFound)
	}
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	// 如果有角色使用菜单无法删除
	roleMenu := entity.RoleMenu{}
	a.data.db.Model(entity.RoleMenu{}).Where("menu_id = ?", id).First(&roleMenu)
	if roleMenu.Id != 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "菜单已被使用,无法删除")
	}

	tx := a.data.db.Begin()
	err = tx.Model(entity.Menu{}).Where("id = ?", id).Delete(&menu).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	// 删除菜单与按钮的关联关系
	err = tx.Model(entity.MenuBtn{}).Where("menu_id = ?", id).Delete(&entity.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	tx.Commit()
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleMenu+":")
	return nil
}
