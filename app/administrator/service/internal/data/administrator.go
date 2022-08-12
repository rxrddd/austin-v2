package data

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/data/entity"
	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/ZQCard/kratos-base-project/app/administrator/service/internal/biz"
)

type administratorRepo struct {
	data *Data
	log  *log.Helper
}

func (a administratorRepo) VerifyPassword(ctx context.Context, id int64, password string) (bool, error) {
	administrator := entity.AdministratorEntity{}
	if err := a.data.db.Model(&entity.AdministratorEntity{}).Where("id = ?", id).First(&administrator).Error; err != nil {
		return false, errors.New(500, "SYSTEM_ERROR", err.Error())

	}
	if administrator.Id != id {
		return false, errors.New(400, "ADMINISTRATOR_MOBILE_EXIST", "ADMINISTRATOR_RECORD_NOT_FOUND")
	}
	return util.CheckPasswordHash(password, administrator.Salt, administrator.Password), nil
}

// searchParam 搜索条件
func (a administratorRepo) searchParam(params map[string]interface{}) *gorm.DB {
	conn := a.data.db.Model(&entity.AdministratorEntity{})
	if id, ok := params["id"]; ok && id.(int64) != 0 {
		conn = conn.Where("id = ?", id)
	}
	if notId, ok := params["notId"]; ok && notId.(int64) != 0 {
		conn = conn.Where("id != ?", notId)
	}
	if username, ok := params["username"]; ok && username.(string) != "" {
		conn = conn.Where("username = ?", username)
	}
	if mobile, ok := params["mobile"]; ok && mobile.(string) != "" {
		conn = conn.Where("mobile = ?", mobile)
	}
	// 包含删除
	if isDeleted, ok := params["is_deleted"]; ok && isDeleted.(string) == entity.AdministratorDeleted {
		conn = conn.Scopes(entity.HasDelete())
	}
	if notDeleted, ok := params["not_deleted"]; ok && notDeleted.(string) == entity.AdministratorUnDeleted {
		conn = conn.Scopes(entity.UnDelete())
	}
	return conn
}

// GetAdministratorByParams 根据条件获取数据
func (a administratorRepo) GetAdministratorByParams(params map[string]interface{}) (record entity.AdministratorEntity, err error) {
	if len(params) == 0 {
		return entity.AdministratorEntity{}, errors.New(http.StatusBadRequest, "MISSING_CONDITION", "缺少搜索条件")
	}
	conn := a.searchParam(params)
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.AdministratorEntity{}, errors.New(http.StatusBadRequest, "ADMINISTRATOR_RECORD_NOT_FOUND", biz.ErrRecordNotFound)
		}
		return record, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	return record, nil
}

func (a administratorRepo) CreateAdministrator(ctx context.Context, reqData *biz.Administrator) (*biz.Administrator, error) {
	modelTable := entity.AdministratorEntity{}
	// 查看用户名是否存在
	record, _ := a.GetAdministratorByParams(map[string]interface{}{
		"username": reqData.Username,
	})
	if record.Id != 0 {
		return nil, errors.New(400, "ADMINISTRATOR_USERNAME_EXIST", "用户名已存在")
	}

	// 查看手机号是否存在
	record, _ = a.GetAdministratorByParams(map[string]interface{}{
		"mobile": reqData.Mobile,
	})
	if record.Id != 0 {
		return nil, errors.New(400, "ADMINISTRATOR_MOBILE_EXIST", "管理员手机号已存在")
	}
	modelTable.Username = reqData.Username
	modelTable.Salt, modelTable.Password = util.HashPassword(reqData.Password)
	modelTable.Mobile = reqData.Mobile
	modelTable.Nickname = reqData.Nickname
	modelTable.Avatar = reqData.Avatar
	modelTable.Status = entity.AdministratorStatusOK

	if err := a.data.db.Model(&modelTable).Create(&modelTable).Error; err != nil {
		return nil, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	// 返回数据
	response := ModelToResponse(modelTable)
	return &response, nil
}

func (a administratorRepo) UpdateAdministrator(ctx context.Context, reqData *biz.Administrator) (*biz.Administrator, error) {
	// 根据id查找记录
	record, err := a.GetAdministratorByParams(map[string]interface{}{
		"id": reqData.Id,
	})
	if err != nil {
		return nil, err
	}
	if record.Id != reqData.Id {
		return nil, errors.New(http.StatusBadRequest, "ADMINISTRATOR_RECORD_NOT_FOUND", biz.ErrRecordNotFound)
	}
	// 查看用户名是否存在
	recordTmp, _ := a.GetAdministratorByParams(map[string]interface{}{
		"username": reqData.Username,
		"notId":    reqData.Id,
	})
	if recordTmp.Id != 0 {
		return nil, errors.New(400, "ADMINISTRATOR_MOBILE_EXIST", "管理员用户名已存在")
	}

	// 查看手机号是否存在
	recordTmp, _ = a.GetAdministratorByParams(map[string]interface{}{
		"mobile": reqData.Mobile,
		"notId":  reqData.Id,
	})
	if recordTmp.Id != 0 {
		return nil, errors.New(400, "ADMINISTRATOR_MOBILE_EXIST", "管理员手机号已存在")
	}
	// 更新记录
	record.Username = reqData.Username
	record.Password = util.HashSaltPassword(record.Salt, reqData.Password)
	record.Mobile = reqData.Mobile
	record.Nickname = reqData.Nickname
	record.Avatar = reqData.Avatar
	if err := a.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return nil, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (a administratorRepo) GetAdministrator(ctx context.Context, params map[string]interface{}) (*biz.Administrator, error) {
	// 根据id查找记录
	record, err := a.GetAdministratorByParams(params)
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (a administratorRepo) ListAdministrator(ctx context.Context, pageNum, pageSize int64) ([]*biz.Administrator, int64, error) {
	list := []entity.AdministratorEntity{}
	conn := a.searchParam(map[string]interface{}{})
	err := conn.Scopes(entity.Paginate(pageNum, pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, errors.New(500, "SYSTEM_ERROR", err.Error())
	}

	count := int64(0)
	conn.Count(&count)
	rv := make([]*biz.Administrator, 0, len(list))
	for _, record := range list {
		serviceName := ModelToResponse(record)
		rv = append(rv, &serviceName)
	}
	return rv, count, nil
}

func (a administratorRepo) DeleteAdministrator(ctx context.Context, id int64) error {
	// 根据id查找记录
	record, err := a.GetAdministratorByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	err = a.data.db.Model(&record).Where("id = ?", id).UpdateColumn("deleted_at", time.Now().Format("2006-01-02 15:04:05")).Error
	if err != nil {
		return errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	return nil
}

// ModelToResponse 转换 administrator 表中所有字段的值
func ModelToResponse(administrator entity.AdministratorEntity) biz.Administrator {
	administratorInfoRsp := biz.Administrator{}
	administratorInfoRsp.Id = administrator.Id
	administratorInfoRsp.Username = administrator.Username
	administratorInfoRsp.Mobile = administrator.Mobile
	administratorInfoRsp.Nickname = administrator.Nickname
	administratorInfoRsp.Avatar = administrator.Avatar
	administratorInfoRsp.Status = administrator.Status
	administratorInfoRsp.CreatedAt = administrator.CreatedAt.Format("2006-01-02 15:04:05")
	administratorInfoRsp.UpdatedAt = administrator.UpdatedAt.Format("2006-01-02 15:04:05")
	administratorInfoRsp.DeletedAt = administrator.DeletedAt
	return administratorInfoRsp
}

func NewAdministratorRepo(data *Data, logger log.Logger) biz.AdministratorRepo {
	return &administratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
