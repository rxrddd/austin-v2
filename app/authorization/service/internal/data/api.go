package data

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"gorm.io/gorm"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

func (a AuthorizationRepo) GetApiAll(ctx context.Context) ([]*biz.Api, error) {
	var res []*biz.Api
	var list []entity.Api
	err := a.data.db.Model(&entity.Api{}).Order("`group` ASC").Find(&list).Error
	if err != nil {
		return nil, errResponse.SetErrByReason(errResponse.ReasonSystemError)
	}
	for _, v := range list {
		res = append(res, &biz.Api{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, nil
}

func (a AuthorizationRepo) GetApiList(ctx context.Context, page int64, pageSize int64, group, name, method, path string) ([]*biz.Api, int64, error) {
	var res []*biz.Api
	var list []entity.Api
	conn := a.data.db.Model(&entity.Api{})

	if name != "" {
		conn = conn.Where("name LIKE ?", "%"+name+"%")
	}
	if method != "" {
		conn = conn.Where("method LIKE ?", "%"+method+"%")
	}
	if path != "" {
		conn = conn.Where("path LIKE ?", "%"+path+"%")
	}
	if group != "" {
		conn = conn.Where("group LIKE ?", "%"+group+"%")
	}
	err := conn.Scopes(entity.Paginate(page, pageSize)).Order("id ASC").Find(&list).Error
	if err != nil {
		return nil, 0, errResponse.SetErrByReason(errResponse.ReasonSystemError)
	}
	count := int64(0)
	conn.Count(&count)
	for _, v := range list {
		res = append(res, &biz.Api{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, count, nil
}

func (a AuthorizationRepo) CreateApi(ctx context.Context, reqData *biz.Api) (*biz.Api, error) {
	var api entity.Api
	// 查看Api是否存在
	err := a.data.db.Model(entity.Api{}).Where("path = ? AND method = ?", reqData.Path, reqData.Method).First(&api).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Api{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiExist)
	}
	now := time.Now()
	api = entity.Api{
		Group:     reqData.Group,
		Name:      reqData.Name,
		Path:      reqData.Path,
		Method:    reqData.Method,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err = a.data.db.Model(entity.Api{}).Create(&api).Error
	if err != nil {
		return &biz.Api{}, err
	}
	res := &biz.Api{
		Id:        api.Id,
		Group:     api.Group,
		Name:      api.Name,
		Path:      api.Path,
		Method:    api.Method,
		CreatedAt: api.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: api.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (a AuthorizationRepo) UpdateApi(ctx context.Context, reqData *biz.Api) (*biz.Api, error) {
	var api entity.Api
	// 查看Api名是否存在
	err := a.data.db.Model(entity.Api{}).Where("path = ? AND method = ? AND id != ?", reqData.Path, reqData.Method, reqData.Id).First(&api).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Api{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiExist)
	}

	err = a.data.db.Model(entity.Api{}).Where("id = ?", reqData.Id).First(&api).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Api{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiNotFound)
	}
	if err != nil {
		return &biz.Api{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	api.Group = reqData.Group
	api.Name = reqData.Name
	api.Method = reqData.Method
	api.Path = reqData.Path
	err = a.data.db.Model(entity.Api{}).Where("id = ?", api.Id).Save(&api).Error
	if err != nil {
		return &biz.Api{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	a.data.db.Model(entity.Api{}).Where("id = ?", api.Id).Find(&api)
	res := &biz.Api{
		Id:        api.Id,
		Group:     api.Group,
		Name:      api.Name,
		Path:      reqData.Path,
		Method:    reqData.Method,
		CreatedAt: api.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: api.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (a AuthorizationRepo) DeleteApi(ctx context.Context, id int64) error {
	var api entity.Api
	// 查看Api是否存在
	err := a.data.db.Model(entity.Api{}).Where("id = ?", id).First(&api).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiNotFound)
	}
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	// 检查api是否被使用，使用中无法删除
	policies := a.data.enforcer.GetFilteredPolicy(0, "", api.Name, api.Method)
	if len(policies) != 0 {
		return kerrors.BadRequest(errResponse.ReasonParamsError, "API已被使用,无法删除")
	}
	err = a.data.db.Model(entity.Api{}).Where("id = ?", id).Delete(&api).Error
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	a.data.enforcer.RemoveFilteredPolicy(0, "", api.Name, api.Method)

	return nil
}
