package service

import (
	"austin-v2/app/files/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "austin-v2/api/files/v1"
)

type FilesService struct {
	v1.UnimplementedFilesServer
	filesUseCase *biz.FilesUseCase
	log          *log.Helper
}

func NewFilesService(FilesUseCase *biz.FilesUseCase,
	logger log.Logger) *FilesService {
	return &FilesService{
		log:          log.NewHelper(log.With(logger, "module", "service/interface")),
		filesUseCase: FilesUseCase,
	}
}

func (s *FilesService) GetOssStsToken(ctx context.Context, req *emptypb.Empty) (*v1.OssStsTokenResponse, error) {
	stsResponse, err := s.filesUseCase.GetOssStsToken(ctx)
	if err != nil {
		return nil, err
	}
	response := &v1.OssStsTokenResponse{}
	response.AccessKey = stsResponse.AccessKey
	response.AccessSecret = stsResponse.AccessSecret
	response.Expiration = stsResponse.Expiration
	response.SecurityToken = stsResponse.SecurityToken
	response.EndPoint = stsResponse.EndPoint
	response.BucketName = stsResponse.BucketName
	response.Region = stsResponse.Region
	response.Url = stsResponse.Url
	return response, nil
}

func (s *FilesService) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileResponse, error) {
	url, err := s.filesUseCase.UploadFile(ctx, req.FileName, req.FileType, req.Content)
	if err != nil {
		return nil, err
	}
	response := &v1.UploadFileResponse{}
	response.Url = url
	return response, err
}
