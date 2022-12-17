package data

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/ZQCard/kratos-base-project/app/files/service/internal/biz"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
)

var OssSessionName = "kratos-base-project"

type FilesRepo struct {
	data *Data
	log  *log.Helper
}

func (f FilesRepo) GetOssStsToken(ctx context.Context) (*biz.OssStsToken, error) {

	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &f.data.config.Oss.AccessKey,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &f.data.config.Oss.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("sts.cn-hangzhou.aliyuncs.com")
	client, err := sts.NewClient(config)
	if err != nil {
		return &biz.OssStsToken{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	assumeRoleRequest := &sts.AssumeRoleRequest{
		RoleArn:         &f.data.config.Oss.StsRoleArn,
		RoleSessionName: &OssSessionName,
	}
	resp, err := client.AssumeRole(assumeRoleRequest)
	if err != nil {
		return &biz.OssStsToken{}, kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}

	response := &biz.OssStsToken{}
	response.AccessKey = *resp.Body.Credentials.AccessKeyId
	response.AccessSecret = *resp.Body.Credentials.AccessKeySecret
	response.Expiration = *resp.Body.Credentials.Expiration
	response.SecurityToken = *resp.Body.Credentials.SecurityToken
	response.EndPoint = f.data.config.Oss.EndPoint
	response.BucketName = f.data.config.Oss.BucketName
	response.Region = f.data.config.Oss.Region
	response.Url = "https://" + f.data.config.Oss.BucketName + "." + f.data.config.Oss.EndPoint + "/"
	return response, nil
}

func NewFilesRepo(data *Data, logger log.Logger) biz.FilesRepo {
	return &FilesRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/file-service")),
	}
}
