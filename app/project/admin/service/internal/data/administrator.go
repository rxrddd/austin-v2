package data

import (
	"context"
	administratorServiceV1 "github.com/ZQCard/kratos-base-project/api/administrator/v1"
	v1 "github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/sync/singleflight"
)

var auth *conf.Auth

func GetAuthApiKey() string {
	return auth.ApiKey
}

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
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
	}
	return res, err
}

func (rp AdministratorRepo) ListAdministrator(ctx context.Context, req *v1.ListAdministratorRequest) (*v1.ListAdministratorReply, error) {
	list := []*v1.AdministratorInfoResponse{}
	reply, err := rp.data.administratorClient.ListAdministrator(ctx, &administratorServiceV1.ListAdministratorRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Mobile:   req.Mobile,
		Username: req.Username,
		Nickname: req.Nickname,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range reply.List {
		tmp := &v1.AdministratorInfoResponse{
			Id:        v.Id,
			Username:  v.Username,
			Mobile:    v.Mobile,
			Nickname:  v.Nickname,
			Avatar:    v.Avatar,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
		list = append(list, tmp)
	}

	response := &v1.ListAdministratorReply{}
	response.Total = reply.Total
	response.List = list
	return response, nil
}

func (rp AdministratorRepo) DeleteAdministrator(ctx context.Context, id int64) (*v1.CheckReply, error) {
	reply, err := rp.data.administratorClient.DeleteAdministrator(ctx, &administratorServiceV1.DeleteAdministratorRequest{
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
	res := &v1.AdministratorInfoResponse{
		Id:        reply.Id,
		Username:  reply.Username,
		Mobile:    reply.Mobile,
		Nickname:  reply.Nickname,
		Avatar:    reply.Avatar,
		Status:    reply.Status,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
		Role:      reply.Role,
	}
	return res, err
}

func (rp AdministratorRepo) FindLoginAdministratorByUsername(ctx context.Context, username string) (*v1.AdministratorInfoResponse, error) {
	reply, err := rp.data.administratorClient.GetAdministrator(ctx, &administratorServiceV1.GetAdministratorRequest{
		Username: username,
	})
	res := &v1.AdministratorInfoResponse{
		Id:        reply.Id,
		Username:  reply.Username,
		Mobile:    reply.Mobile,
		Nickname:  reply.Nickname,
		Avatar:    reply.Avatar,
		Status:    reply.Status,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
		Role:      reply.Role,
	}
	return res, err
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
		return err
	}
	return nil
}
