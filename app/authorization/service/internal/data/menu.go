package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	entity2 "github.com/ZQCard/kratos-base-project/app/authorization/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"gorm.io/gorm"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

func (a AuthorizationRepo) GetMenuAll(ctx context.Context) ([]*biz.Menu, error) {
	var res []*biz.Menu
	var menus []entity2.Menu
	// 获取所有根菜单
	err := a.data.db.Model(entity2.Menu{}).Preload("MenuBtns").Find(&menus).Error
	if err != nil {
		return res, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
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
		res = append(res, &biz.Menu{
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
	return res, nil
}

func (a AuthorizationRepo) GetMenuTree(ctx context.Context) ([]*biz.Menu, error) {
	var res []*biz.Menu
	var menus []entity2.Menu
	// 获取所有根菜单
	err := a.data.db.Model(entity2.Menu{}).Where("parent_id = 0").Preload("MenuBtns").Order("sort ASC").Find(&menus).Error
	if err != nil {
		return res, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
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
		res = append(res, &biz.Menu{
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
			MenuBtns:  btns,
		})
	}
	for k := range res {
		err := a.findChildrenMenu(res[k])
		fmt.Println("res[k].MenuBtns")
		fmt.Println(res[k].MenuBtns)
		if err != nil {
			return res, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
	}

	return res, nil
}

func (a AuthorizationRepo) findChildrenMenu(menu *biz.Menu) (err error) {
	var tmp []entity2.Menu
	err = a.data.db.Model(entity2.Menu{}).Where("parent_id = ?", menu.Id).Preload("MenuBtns").Find(&tmp).Error
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
			err = a.findChildrenMenu(&menu.Children[k])
		}
	}
	return err
}

func (a AuthorizationRepo) CreateMenu(ctx context.Context, reqData *biz.Menu) (*biz.Menu, error) {
	btns := []*entity2.MenuBtn{}
	for _, v := range reqData.MenuBtns {
		btns = append(btns, &entity2.MenuBtn{
			Name:        v.Name,
			Description: v.Description,
		})
	}
	var menu entity2.Menu
	now := time.Now()
	menu = entity2.Menu{
		Name:      reqData.Name,
		ParentId:  reqData.ParentId,
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
	err := a.data.db.Model(entity2.Menu{}).Create(&menu).Error
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
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	res := &biz.Menu{
		Id:        menu.Id,
		Name:      menu.Name,
		Path:      menu.Path,
		ParentId:  menu.ParentId,
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
		MenuBtns:  btns2,
	}
	return res, nil
}

func (a AuthorizationRepo) UpdateMenu(ctx context.Context, reqData *biz.Menu) (*biz.Menu, error) {
	btns := []*entity2.MenuBtn{}
	for _, v := range reqData.MenuBtns {
		btns = append(btns, &entity2.MenuBtn{
			Id:          v.Id,
			MenuId:      reqData.Id,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	var menu entity2.Menu
	menu = entity2.Menu{
		Id:        reqData.Id,
		Name:      reqData.Name,
		ParentId:  reqData.ParentId,
		Path:      reqData.Path,
		Hidden:    reqData.Hidden,
		Component: reqData.Component,
		Sort:      reqData.Sort,
		Title:     reqData.Title,
		Icon:      reqData.Icon,
		MenuBtns:  btns,
	}
	// 关联数据更新
	tx := a.data.db.Begin()
	err := tx.Model(entity2.Menu{}).Where("id = ?", menu.Id).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&menu).Error
	if err != nil {
		tx.Rollback()
		return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 先删除,后添加
	if err = tx.Where("menu_id  = ?", menu.Id).Unscoped().Delete(&entity2.MenuBtn{}).Error; err != nil {
		tx.Rollback()
		return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 保存按钮
	for _, v := range menu.MenuBtns {
		if err = tx.Model(entity2.MenuBtn{}).Where("id = ?", v.Id).Create(&v).Error; err != nil {
			tx.Rollback()
			return &biz.Menu{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
		}
	}
	tx.Commit()
	a.data.db.Model(entity2.Menu{}).Where("id = ?", menu.Id).Preload("MenuBtns").Find(&menu)
	btns2 := []biz.MenuBtn{}
	for _, v := range menu.MenuBtns {
		btns2 = append(btns2, biz.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	res := &biz.Menu{
		Id:        menu.Id,
		Name:      menu.Name,
		Path:      menu.Path,
		ParentId:  menu.ParentId,
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		MenuBtns:  btns2,
		CreatedAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (a AuthorizationRepo) DeleteMenu(ctx context.Context, id int64) error {
	var menu entity2.Menu
	// 查看菜单是否存在
	err := a.data.db.Model(entity2.Menu{}).Where("id = ?", id).First(&menu).Error
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errResponse.SetErrByReason(errResponse.ReasonRecordNotFound)
	}
	tx := a.data.db.Begin()
	err = tx.Model(entity2.Menu{}).Where("id = ?", id).Delete(&menu).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 删除菜单与角色的关联关系
	err = tx.Model(entity2.RoleMenu{}).Where("menu_id = ?", id).Delete(&entity2.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 删除菜单与按钮的关联关系
	err = tx.Model(entity2.MenuBtn{}).Where("menu_id = ?", id).Delete(&entity2.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	tx.Commit()
	return nil
}
