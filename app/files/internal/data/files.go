package data

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/app/files/internal/biz"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

var OssSessionName = "kratos-base-project"

type FilesRepo struct {
	data *Data
	log  *log.Helper
}

func (f FilesRepo) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(f.data.config.Oss.EndPoint, f.data.config.Oss.AccessKey, f.data.config.Oss.AccessSecret)
	if err != nil {
		return "", errResponse.SetErrByReason(errResponse.ReasonOssConfigWrong)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(f.data.config.Oss.BucketName)
	if err != nil {
		return "", errResponse.SetErrByReason(errResponse.ReasonOssConfigWrong)
	}
	path := "uploadFile/" + strings.Trim(fileType, ".") + "/" + fileName + fileType
	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject(path, bytes.NewReader(content))
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return "", errResponse.SetErrByReason(errResponse.ReasonOssPutObjectFail)
	}
	url := "https://" + f.data.config.Oss.BucketName + "." + f.data.config.Oss.EndPoint + "/" + path
	return url, nil
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
