package data

import (
	"context"
	"fmt"
	administratorServiceV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	v1 "github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/utils/encryption"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang-jwt/jwt/v4"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/sync/singleflight"
	"time"
)

func NewAdministratorServiceClient(ac *conf.Auth, sr *conf.Service, r registry.Discovery, tp *tracesdk.TracerProvider) administratorServiceV1.AdministratorClient {
	// 初始化auth配置
	auth = ac

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Administrator.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			//jwt.Client(func(token *jwt2.Token) (interface{}, error) {
			//	return []byte(ac.ServiceKey), nil
			//}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := administratorServiceV1.NewAdministratorClient(conn)
	return c
}

type AdministratorRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (rp AdministratorRepo) CreateAdministrator(ctx context.Context, req *v1.CreateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {
	reply, err := rp.data.administratorClient.CreateAdministrator(ctx, &administratorServiceV1.CreateAdministratorRequest{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.AdministratorInfoResponse{
		Id:        reply.Id,
		Username:  reply.Username,
		Mobile:    reply.Mobile,
		Nickname:  reply.Nickname,
		Avatar:    reply.Avatar,
		Status:    reply.Status,
		Role:      reply.Role,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
	}
	return res, err
}

func (rp AdministratorRepo) UpdateAdministrator(ctx context.Context, req *v1.UpdateAdministratorRequest) (*v1.AdministratorInfoResponse, error) {

	reply, err := rp.data.administratorClient.UpdateAdministrator(ctx, &administratorServiceV1.UpdateAdministratorRequest{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Status:   req.Status,
	})

	if err != nil {
		return nil, err
	}
	res := &v1.AdministratorInfoResponse{
		Id:        reply.Id,
		Username:  reply.Username,
		Mobile:    reply.Mobile,
		Nickname:  reply.Nickname,
		Avatar:    reply.Avatar,
		Status:    reply.Status,
		Role:      reply.Role,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
	}
	return res, err
}

func (rp AdministratorRepo) ListAdministrator(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	fmt.Println("CreatedAtStart")
	fmt.Println(req.CreatedAtStart)
	fmt.Println(req.CreatedAtEnd)
	list := []*v1.AdministratorInfoResponse{}
	reply, err := rp.data.administratorClient.ListAdministrator(ctx, &administratorServiceV1.ListAdministratorRequest{
		PageNum:        req.PageNum,
		PageSize:       req.PageSize,
		Mobile:         req.Mobile,
		Username:       req.Username,
		Nickname:       req.Nickname,
		Status:         req.Status,
		CreatedAtStart: req.CreatedAtStart,
		CreatedAtEnd:   req.CreatedAtEnd,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range reply.List {
		tmp := administratorServiceToApi(v)
		list = append(list, tmp)
	}

	response := &v1.ListAdministratorReply{}
	response.Total = reply.Total
	response.List = list
	return response, nil
}

func (rp AdministratorRepo) DeleteAdministrator(ctx context.Context, id int64) (*v1.CheckReply, error) {
	administratorInfoReply, err := rp.data.administratorClient.GetAdministrator(ctx, &administratorServiceV1.GetAdministratorRequest{Id: id})
	if err != nil {
		return nil, err
	}
	if administratorInfoReply.Id != id {
		return nil, errResponse.SetErrByReason(errResponse.ReasonAdministratorNotFound)
	}
	reply, err := rp.data.administratorClient.DeleteAdministrator(ctx, &administratorServiceV1.DeleteAdministratorRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	// 删除成功 将管理员token清除
	if reply.IsSuccess {
		administratorInfo := administratorServiceToApi(administratorInfoReply)
		_ = rp.DestroyAdministratorToken(ctx, administratorInfo)
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}

	return res, nil
}

func (rp AdministratorRepo) RecoverAdministrator(ctx context.Context, id int64) (*v1.CheckReply, error) {
	reply, err := rp.data.administratorClient.RecoverAdministrator(ctx, &administratorServiceV1.RecoverAdministratorRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AdministratorRepo) ForbidAdministrator(ctx context.Context, id int64) (*v1.CheckReply, error) {
	reply, err := rp.data.administratorClient.AdministratorStatusChange(ctx, &administratorServiceV1.AdministratorStatusChangeRequest{
		Id:     id,
		Status: 2,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AdministratorRepo) ApproveAdministrator(ctx context.Context, id int64) (*v1.CheckReply, error) {
	reply, err := rp.data.administratorClient.AdministratorStatusChange(ctx, &administratorServiceV1.AdministratorStatusChangeRequest{
		Id:     id,
		Status: 1,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func NewAdministratorRepo(data *Data, logger log.Logger) *AdministratorRepo {
	return &AdministratorRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}

func (rp AdministratorRepo) GetAdministrator(ctx context.Context, id int64) (*v1.AdministratorInfoResponse, error) {
	reply, err := rp.data.administratorClient.GetAdministrator(ctx, &administratorServiceV1.GetAdministratorRequest{
		Id: id,
	})
	res := &v1.AdministratorInfoResponse{}
	if err != nil {
		if reply == nil || reply.Id == 0 {
			return res, errResponse.SetErrByReason(errResponse.ReasonAdministratorNotFound)
		}
		return res, err
	}
	res = administratorServiceToApi(reply)
	return res, err
}

func (rp AdministratorRepo) FindLoginAdministratorByUsername(ctx context.Context, username string) (*v1.AdministratorInfoResponse, error) {
	reply, err := rp.data.administratorClient.GetAdministrator(ctx, &administratorServiceV1.GetAdministratorRequest{
		Username: username,
	})
	res := &v1.AdministratorInfoResponse{}
	if err != nil {
		return res, err
	}
	// 如果管理员被删除   无法登陆
	if reply.DeletedAt != "" {
		return res, errResponse.SetErrByReason(errResponse.ReasonAdministratorDeleted)
	}
	res = administratorServiceToApi(reply)
	return res, err
}

func (rp AdministratorRepo) AdministratorLoginSuccess(ctx context.Context, administrator *v1.AdministratorInfoResponse) error {
	reply, err := rp.data.administratorClient.AdministratorLoginSuccess(ctx, &administratorServiceV1.AdministratorLoginSuccessRequest{
		Id:            administrator.Id,
		LastLoginTime: administrator.LastLoginTime,
		LastLoginIp:   administrator.LastLoginIp,
	})
	if err != nil {
		return err
	}

	if reply.IsSuccess == false {
		return errResponse.SetErrByReason(errResponse.ReasonAdministratorNotFound)
	}
	return nil
}

func (rp AdministratorRepo) VerifyPassword(ctx context.Context, id int64, password string) error {
	reply, err := rp.data.administratorClient.VerifyAdministratorPassword(ctx, &administratorServiceV1.VerifyAdministratorPasswordRequest{
		Id:       id,
		Password: password,
	})
	if err != nil {
		return err
	}

	if reply.IsSuccess == false {
		return errResponse.SetErrByReason(errResponse.ReasonAdministratorPasswordError)
	}
	return nil
}

func (rp AdministratorRepo) GenerateAdministratorToken(ctx context.Context, administrator *v1.AdministratorInfoResponse) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"AdministratorId":            administrator.Id,
			"AdministratorUsername":      administrator.Username,
			"AdministratorRole":          administrator.Role,
			"AdministratorLastLoginTime": administrator.LastLoginTime,
			"AdministratorLastLoginIp":   administrator.LastLoginIp,
		})
	signedString, _ := claims.SignedString([]byte(GetAuthApiKey()))
	key := encryption.EncodeMD5(signedString)
	// 生成redis
	rp.data.redisCli.Set(key, signedString, time.Second*time.Duration(auth.ApiKeyExpire))
	return key, nil
}

func (rp AdministratorRepo) DestroyAdministratorToken(ctx context.Context, administrator *v1.AdministratorInfoResponse) error {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"AdministratorId":            administrator.Id,
			"AdministratorUsername":      administrator.Username,
			"AdministratorRole":          administrator.Role,
			"AdministratorLastLoginTime": administrator.LastLoginTime,
			"AdministratorLastLoginIp":   administrator.LastLoginIp,
		})
	signedString, _ := claims.SignedString([]byte(GetAuthApiKey()))
	key := encryption.EncodeMD5(signedString)
	fmt.Println("signedString")
	fmt.Println(key)
	// 删除redis
	rp.data.redisCli.Del(key)
	return nil
}

func administratorServiceToApi(info *administratorServiceV1.AdministratorInfoResponse) *v1.AdministratorInfoResponse {
	return &v1.AdministratorInfoResponse{
		Id:            info.Id,
		Username:      info.Username,
		Nickname:      info.Nickname,
		Mobile:        info.Mobile,
		Status:        info.Status,
		Avatar:        info.Avatar,
		Role:          info.Role,
		LastLoginTime: info.LastLoginTime,
		LastLoginIp:   info.LastLoginIp,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.CreatedAt,
		DeletedAt:     info.DeletedAt,
	}
}
