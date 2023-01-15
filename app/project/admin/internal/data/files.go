package data

import (
	filesServiceV1 "austin-v2/api/files/v1"
	v1 "austin-v2/api/project/admin/v1"
	"austin-v2/app/project/admin/internal/conf"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewFilesServiceClient(_ *conf.Auth, sr *conf.Service, r registry.Discovery) filesServiceV1.FilesClient {
	// 初始化auth配置
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Files.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := filesServiceV1.NewFilesClient(conn)
	return c
}

func NewFilesRepo(data *Data, logger log.Logger) *FilesRepo {
	return &FilesRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}

type FilesRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (rp FilesRepo) GetOssStsToken(ctx context.Context) (*v1.OssStsTokenResponse, error) {
	reply, err := rp.data.filesClient.GetOssStsToken(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	fmt.Println("reply")
	fmt.Println(reply)
	response := &v1.OssStsTokenResponse{}
	response.AccessKey = reply.AccessKey
	response.AccessSecret = reply.AccessSecret
	response.Expiration = reply.Expiration
	response.SecurityToken = reply.SecurityToken
	response.EndPoint = reply.EndPoint
	response.BucketName = reply.BucketName
	response.Region = reply.Region
	response.Url = reply.Url
	return response, err
}

func (rp FilesRepo) UploadFile(ctx context.Context, fileName string, fileType string, context []byte) (string, error) {
	reply, err := rp.data.filesClient.UploadFile(ctx, &filesServiceV1.UploadFileRequest{
		FileName: fileName,
		FileType: fileType,
		Content:  context,
	})
	if err != nil {
		return "", err
	}
	return reply.Url, err
}
