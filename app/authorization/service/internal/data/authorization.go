package data

import (
	"context"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/conf"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
)

var casbinEnforcer *casbin.Enforcer

type AuthorizationRepo struct {
	data *Data
	log  *log.Helper
}

func (a AuthorizationRepo) CheckAuthorization(ctx context.Context, authorization *biz.Authorization) (bool, error) {
	return casbinEnforcer.Enforce(authorization.Sub, authorization.Obj, authorization.Act)
}

func NewAuthorizationRepo(data *Data, conf *conf.Casbin, logger log.Logger) biz.AuthorizationRepo {
	// 初始化基础数据库 casbin权限控制策略,连接基础库
	db, _ := gormadapter.NewAdapterByDB(data.db)
	// 加载权限配置文件
	m, err := model.NewModelFromString(conf.RbacModel)
	if err != nil {
		log.Fatalf(err.Error())
	}
	enforcer, err := casbin.NewEnforcer(m, db)
	// 开启权限日志
	enforcer.EnableLog(true)
	// 从DB加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf(err.Error())
	}
	data.enforcer = enforcer
	// 赋值全局变量
	casbinEnforcer = data.enforcer
	// 数据化数据
	return &AuthorizationRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/authorization-service")),
	}
}

// GetCasbinEnforcer 获取已经初始化过的casbin对象
func GetCasbinEnforcer() *casbin.Enforcer {
	return casbinEnforcer
}
