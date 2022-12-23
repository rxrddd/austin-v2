package data

import (
	"context"
	"encoding/json"
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/biz"
	entity2 "github.com/ZQCard/kratos-base-project/app/administrator/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/utils/encryption"
	"github.com/ZQCard/kratos-base-project/pkg/utils/redisHelper"
	"github.com/ZQCard/kratos-base-project/pkg/utils/timeHelper"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type AdministratorRepo struct {
	data *Data
	log  *log.Helper
}

func (s AdministratorRepo) GetAdministratorByParams(params map[string]interface{}) (record entity2.AdministratorEntity, err error) {
	if len(params) == 0 {
		return entity2.AdministratorEntity{}, errResponse.SetErrByReason(errResponse.ReasonMissingParams)
	}
	conn := s.data.db.Model(&entity2.AdministratorEntity{})
	if id, ok := params["id"]; ok && id != nil {
		conn = conn.Where("id = ?", id)
	}
	if nickname, ok := params["nickname_like"]; ok && nickname != nil && nickname.(string) != "" {
		conn = conn.Where("nickname LIKE ?", "%"+nickname.(string)+"%")
	}
	if nickname, ok := params["nickname"]; ok && nickname != nil && nickname.(string) != "" {
		conn = conn.Where("nickname = ?", nickname)
	}
	if username, ok := params["username_like"]; ok && username != nil && username.(string) != "" {
		conn = conn.Where("username LIKE ?", "%"+username.(string)+"%")
	}
	if username, ok := params["username"]; ok && username != nil && username.(string) != "" {
		conn = conn.Where("username = ?", username)
	}
	if mobile, ok := params["mobile"]; ok && mobile != nil && mobile.(string) != "" {
		conn = conn.Where("mobile = ?", mobile)
	}
	if status, ok := params["status"]; ok && status != nil && status.(int64) != 0 {
		conn = conn.Where("status = ?", status)
	}
	if role, ok := params["role"]; ok && role != nil && role.(string) != "" {
		conn = conn.Where("role = ?", role)
	}
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity2.AdministratorEntity{}, errResponse.SetErrByReason(errResponse.ReasonAdministratorNotFound)
		}
		return record, errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}
	return record, nil
}

func (s AdministratorRepo) CreateAdministrator(ctx context.Context, reqData *biz.Administrator) (*biz.Administrator, error) {
	// 查看用户名是否存在
	recordTmp, _ := s.GetAdministratorByParams(map[string]interface{}{
		"username": reqData.Username,
	})
	if recordTmp.Id != 0 {
		return nil, errResponse.SetErrByReason(errResponse.ReasonAdministratorUsernameExist)
	}
	// 查看手机号是否存在
	recordTmp, _ = s.GetAdministratorByParams(map[string]interface{}{
		"mobile": reqData.Mobile,
	})
	if recordTmp.Id != 0 {
		return nil, errResponse.SetErrByReason(errResponse.ReasonAdministratorMobileExist)
	}
	salt, password := encryption.HashPassword(reqData.Password)
	modelTable := entity2.AdministratorEntity{
		Username:  reqData.Username,
		Salt:      salt,
		Password:  password,
		Nickname:  reqData.Nickname,
		Mobile:    reqData.Mobile,
		Status:    entity2.AdministratorStatusOK,
		Role:      reqData.Role,
		Avatar:    reqData.Avatar,
		CreatedAt: timeHelper.GetCurrentTime(),
		UpdatedAt: timeHelper.GetCurrentTime(),
		DeletedAt: "",
	}

	modelTable.Id = reqData.Id
	if err := s.data.db.Model(&modelTable).Create(&modelTable).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())

	}
	response := ModelToResponse(modelTable)
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return &response, nil
}

func (s AdministratorRepo) UpdateAdministrator(ctx context.Context, reqData *biz.Administrator) (*biz.Administrator, error) {
	// 根据id查找记录
	record, err := s.GetAdministratorByParams(map[string]interface{}{
		"id": reqData.Id,
	})
	if err != nil {
		return nil, err
	}
	// 如果 密码不为空，则更新密码和盐
	if reqData.Password != "" {
		salt, password := encryption.HashPassword(reqData.Password)
		record.Salt = salt
		record.Password = password
	}
	record.Avatar = reqData.Avatar
	record.Nickname = reqData.Nickname
	record.Status = reqData.Status
	// 更新字段
	if err := s.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}
	// 返回数据
	response := ModelToResponse(record)
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return &response, nil
}

func (s AdministratorRepo) UpdateAdministratorLoginInfo(ctx context.Context, id int64, loginTime string, loginIp string) error {
	// 更新登陆信息
	err := s.data.db.Model(&entity2.AdministratorEntity{}).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"last_login_ip":   loginIp,
		"last_login_time": loginTime,
	}).Error
	if err != nil {
		return errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return nil
}

func (s AdministratorRepo) GetAdministrator(ctx context.Context, params map[string]interface{}) (*biz.Administrator, error) {
	response := &biz.Administrator{}

	// 缓存key
	cacheKey := redisHelper.GetRedisCacheKeyByParams(s.data.Module, params)
	// 查看缓存
	if cache := redisHelper.GetRedisCache(s.data.redisCli, cacheKey); cache != "" {
		if err := json.Unmarshal([]byte(cache), response); err == nil {
			return response, nil
		} else {
			s.log.Error("GetAdministrator()", err)
		}
	}
	record, err := s.GetAdministratorByParams(params)
	if err != nil {
		return response, err
	}
	// 返回数据
	tmp := ModelToResponse(record)
	response = &tmp
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(s.data.redisCli, cacheKey, string(jsonResponse))
	return response, nil
}

func (s AdministratorRepo) ListAdministrator(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*biz.Administrator, int64, error) {
	response := []*biz.Administrator{}

	// 缓存key
	cacheParams := params
	cacheParams["page"] = page
	cacheParams["pageSize"] = pageSize
	cacheKey := redisHelper.GetRedisCacheKeyByParams(s.data.Module, cacheParams)
	countCacheKey := cacheKey + ":count"
	// 查看缓存
	if cache := redisHelper.GetRedisCache(s.data.redisCli, cacheKey); cache != "" {
		countStr := redisHelper.GetRedisCache(s.data.redisCli, countCacheKey)
		count, _ := strconv.ParseInt(countStr, 10, 64)
		if err := json.Unmarshal([]byte(cache), &response); err == nil {
			return response, count, nil
		} else {
			s.log.Error("ListAdministrator()", err)
		}
	}
	list := []entity2.AdministratorEntity{}
	conn := s.data.db.Model(&entity2.AdministratorEntity{})
	if id, ok := params["id"]; ok && id != nil {
		conn = conn.Where("id = ?", id)
	}
	if nickname, ok := params["nickname"]; ok && nickname != nil && nickname.(string) != "" {
		conn = conn.Where("nickname LIKE ?", "%"+nickname.(string)+"%")
	}

	if username, ok := params["username"]; ok && username != nil && username.(string) != "" {
		conn = conn.Where("username LIKE ?", "%"+username.(string)+"%")
	}

	if mobile, ok := params["mobile"]; ok && mobile != nil && mobile.(string) != "" {
		conn = conn.Where("mobile LIKE ?", "%"+mobile.(string)+"%")
	}

	if status, ok := params["status"]; ok && status != nil && status.(int64) != 0 {
		conn = conn.Where("status = ?", status)
	}
	// 开始时间 检查日期格式
	if start, ok := params["created_at_start"]; ok && start != nil && start.(string) != "" {
		tmp := start.(string)
		if !timeHelper.CheckDateFormat(tmp) {
			return nil, 0, errResponse.SetErrByReason(errResponse.TimeFormatError)
		}
		tmp = tmp + " 00:00:00"
		conn = conn.Where("created_at >= ?", tmp)
	}
	// 结束时间
	if end, ok := params["created_at_end"]; ok && end != nil && end.(string) != "" {
		tmp := end.(string)
		if !timeHelper.CheckDateFormat(tmp) {
			return nil, 0, errResponse.SetErrByReason(errResponse.TimeFormatError)
		}
		tmp = tmp + " 23:59:59"
		conn = conn.Where("created_at <= ?", tmp)
	}

	err := conn.Scopes(entity2.Paginate(page, pageSize)).Find(&list).Error
	if err != nil {
		return response, 0, errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}

	count := int64(0)
	conn.Count(&count)
	for _, record := range list {
		administrator := ModelToResponse(record)
		response = append(response, &administrator)
	}

	// 返回数据
	jsonResponse, _ := json.Marshal(response)
	_ = redisHelper.SaveRedisCache(s.data.redisCli, cacheKey, string(jsonResponse))
	_ = redisHelper.SaveRedisCache(s.data.redisCli, countCacheKey, strconv.FormatInt(count, 10))
	return response, count, nil
}

func (s AdministratorRepo) DeleteAdministrator(ctx context.Context, id int64) error {
	// 根据id查找记录
	record, err := s.GetAdministratorByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	if err := s.data.db.Model(&record).Where("id = ?", id).UpdateColumn("deleted_at", timeHelper.GetCurrentYMDHIS()).Error; err != nil {
		return err
	}
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return nil
}

func (s AdministratorRepo) RecoverAdministrator(ctx context.Context, id int64) error {
	if id == 0 {
		return errResponse.SetErrByReason(errResponse.ReasonMissingParams)
	}
	err := s.data.db.Model(entity2.AdministratorEntity{}).Where("id = ?", id).UpdateColumn("deleted_at", "").Error
	if err != nil {
		return errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return nil
}
func (s AdministratorRepo) AdministratorStatusChange(ctx context.Context, id int64, status int64) error {
	if id == 0 || status == 0 {
		return errResponse.SetErrByReason(errResponse.ReasonMissingParams)
	}
	if status != entity2.AdministratorStatusOK && status != entity2.AdministratorStatusForbid {
		return errResponse.SetErrByReason(errResponse.ReasonParamsError)
	}
	err := s.data.db.Model(entity2.AdministratorEntity{}).Where("id = ?", id).UpdateColumn("status", status).Error
	if err != nil {
		return errors.New(http.StatusInternalServerError, errResponse.ReasonSystemError, err.Error())
	}
	redisHelper.BatchDeleteRedisCache(s.data.redisCli, s.data.Module)
	return nil
}

func (s AdministratorRepo) VerifyAdministratorPassword(ctx context.Context, id int64, password string) (bool, error) {
	administrator := entity2.AdministratorEntity{}
	if err := s.data.db.Model(&entity2.AdministratorEntity{}).Where("id = ?", id).First(&administrator).Error; err != nil {
		return false, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	if administrator.Id != id {
		return false, errors.New(400, "ADMINISTRATOR_MOBILE_EXIST", "ADMINISTRATOR_RECORD_NOT_FOUND")
	}
	return encryption.CheckPasswordHash(password, administrator.Salt, administrator.Password), nil
}

// ModelToResponse 转换 administrator 表中所有字段的值
func ModelToResponse(administrator entity2.AdministratorEntity) biz.Administrator {
	administratorInfoRsp := biz.Administrator{
		Id:            administrator.Id,
		Username:      administrator.Username,
		Salt:          administrator.Salt,
		Password:      administrator.Password,
		Nickname:      administrator.Nickname,
		Mobile:        administrator.Mobile,
		Status:        administrator.Status,
		Avatar:        administrator.Avatar,
		Role:          administrator.Role,
		LastLoginIp:   administrator.LastLoginIp,
		LastLoginTime: administrator.LastLoginTime,
		CreatedAt:     timeHelper.FormatYMDHIS(administrator.CreatedAt),
		UpdatedAt:     timeHelper.FormatYMDHIS(administrator.UpdatedAt),
		DeletedAt:     administrator.DeletedAt,
	}
	return administratorInfoRsp
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &AdministratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
