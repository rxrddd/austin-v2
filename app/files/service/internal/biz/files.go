package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OssStsToken struct {
	AccessKey     string
	AccessSecret  string
	Expiration    string
	SecurityToken string
	EndPoint      string
	BucketName    string
	Region        string
	Url           string
}

// FilesRepo 模块接口
type FilesRepo interface {
	GetOssStsToken(ctx context.Context) (*OssStsToken, error)
}

type FilesUseCase struct {
	repo FilesRepo
	log  *log.Helper
}

func NewFilesUseCase(repo FilesRepo, logger log.Logger) *FilesUseCase {
	return &FilesUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *FilesUseCase) GetOssStsToken(ctx context.Context) (*OssStsToken, error) {
	return uc.repo.GetOssStsToken(ctx)
}
