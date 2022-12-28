package data

import (
	"austin-v2/app/authorization/internal/biz"
	entity2 "austin-v2/app/authorization/internal/data/entity"
	"austin-v2/pkg/errResponse"
	"austin-v2/pkg/utils/redisHelper"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

const childModuleAPI = "API"

func (a AuthorizationRepo) GetApiAll(ctx context.Context) ([]*biz.Api, error) {
	// 缓存key
	cacheParams := map[string]interface{}{
		"type": "all",
	}
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleAPI+":", cacheParams)
	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		res := []*biz.Api{}
		if err := json.Unmarshal([]byte(cache), &res); err == nil {
			return res, nil
		} else {
			a.log.Error("GetApiAll()", err)
		}
	}

	var response []*biz.Api
	var list []entity2.Api
	err := a.data.db.Model(&entity2.Api{}).Order("`id` ASC").Find(&list).Error
	if err != nil {
		return nil, errResponse.SetErrByReason(errResponse.ReasonSystemError)
	}
	for _, v := range list {
		response = append(response, &biz.Api{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(a.data.redisCli, cacheKey, string(jsonResponse))
	return response, nil
}

func (a AuthorizationRepo) GetApiList(ctx context.Context, page int64, pageSize int64, params map[string]interface{}) ([]*biz.Api, int64, error) {
	// 缓存key
	cacheParams := params
	cacheParams["page"] = page
	cacheParams["pageSize"] = pageSize
	cacheKey := redisHelper.GetRedisCacheKeyByParams(a.data.Module+":"+childModuleAPI+":", cacheParams)
	countCacheKey := cacheKey + ":count"

	var response []*biz.Api
	var list []entity2.Api

	// 查看缓存
	if cache := a.GetRedisCache(cacheKey); cache != "" {
		countStr := a.GetRedisCache(countCacheKey)
		count, _ := strconv.ParseInt(countStr, 10, 64)
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, count, nil
		} else {
			a.log.Error("ListAdministrator()", err)
		}
	}

	conn := a.data.db.Model(&entity2.Api{})

	if name, ok := params["name"]; ok && name != nil && name.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+name.(string)+"%")
	}

	if method, ok := params["method"]; ok && method != nil && method.(string) != "" {
		conn = conn.Where("method LIKE ?", "%"+method.(string)+"%")
	}

	if path, ok := params["path"]; ok && path != nil && path.(string) != "" {
		conn = conn.Where("path LIKE ?", "%"+path.(string)+"%")
	}

	if group, ok := params["group"]; ok && group != nil && group.(string) != "" {
		conn = conn.Where("group LIKE ?", "%"+group.(string)+"%")
	}

	err := conn.Scopes(entity2.Paginate(page, pageSize)).Order("id ASC").Find(&list).Error
	if err != nil {
		return nil, 0, errResponse.SetErrByReason(errResponse.ReasonSystemError)
	}
	count := int64(0)
	conn.Count(&count)
	for _, v := range list {
		response = append(response, &biz.Api{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(a.data.redisCli, cacheKey, string(jsonResponse))
	_ = redisHelper.SaveRedisCache(a.data.redisCli, countCacheKey, strconv.FormatInt(count, 10))
	return response, count, nil
}

func (a AuthorizationRepo) CreateApi(ctx context.Context, reqData *biz.Api) (*biz.Api, error) {
	var api entity2.Api
	// 查看Api是否存在
	err := a.data.db.Model(entity2.Api{}).Where("path = ? AND method = ?", reqData.Path, reqData.Method).First(&api).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Api{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiExist)
	}
	now := time.Now()
	api = entity2.Api{
		Group:     reqData.Group,
		Name:      reqData.Name,
		Path:      reqData.Path,
		Method:    reqData.Method,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err = a.data.db.Model(entity2.Api{}).Create(&api).Error
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
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleAPI+":")
	return res, nil
}

func (a AuthorizationRepo) UpdateApi(ctx context.Context, reqData *biz.Api) (*biz.Api, error) {
	var api entity2.Api
	// 查看Api名是否存在
	err := a.data.db.Model(entity2.Api{}).Where("path = ? AND method = ? AND id != ?", reqData.Path, reqData.Method, reqData.Id).First(&api).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.Api{}, errResponse.SetErrByReason(errResponse.ReasonAuthorizationApiExist)
	}

	err = a.data.db.Model(entity2.Api{}).Where("id = ?", reqData.Id).First(&api).Error
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
	err = a.data.db.Model(entity2.Api{}).Where("id = ?", api.Id).Save(&api).Error
	if err != nil {
		return &biz.Api{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	a.data.db.Model(entity2.Api{}).Where("id = ?", api.Id).Find(&api)
	res := &biz.Api{
		Id:        api.Id,
		Group:     api.Group,
		Name:      api.Name,
		Path:      reqData.Path,
		Method:    reqData.Method,
		CreatedAt: api.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: api.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleAPI+":")
	return res, nil
}

func (a AuthorizationRepo) DeleteApi(ctx context.Context, id int64) error {
	var api entity2.Api
	// 查看Api是否存在
	err := a.data.db.Model(entity2.Api{}).Where("id = ?", id).First(&api).Error
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
	err = a.data.db.Model(entity2.Api{}).Where("id = ?", id).Delete(&api).Error
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	a.data.enforcer.RemoveFilteredPolicy(0, "", api.Name, api.Method)
	redisHelper.BatchDeleteRedisCache(a.data.redisCli, a.data.Module+":"+childModuleAPI+":")
	return nil
}
