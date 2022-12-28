package service

import (
	v1 "austin-v2/api/project/admin/v1"
	"austin-v2/pkg/errResponse"
	"bytes"
	"context"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"path"
)

func (s *AdminInterface) GetOssStsToken(ctx context.Context, pb *emptypb.Empty) (*v1.OssStsTokenResponse, error) {
	reply, err := s.filesRepo.GetOssStsToken(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *AdminInterface) UploadFile(ctx khttp.Context) (err error) {
	fileName := ctx.Request().FormValue("file_name")
	file, fileHeader, _ := ctx.Request().FormFile("file")
	if fileName == "" {
		return
	}
	if file == nil {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonFileMissing)
	}
	defer file.Close()
	if err != nil {
		return err
	}
	// 文件最大为5M
	if fileHeader.Size > 1024*1024*5 {
		return errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonFileOverLimitSize)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil
	}
	url, err := s.filesRepo.UploadFile(ctx, fileName, path.Ext(fileHeader.Filename), buf.Bytes())
	if err != nil {
		return kerrors.InternalServer(errResponse.ReasonSystemError, err.Error())
	}
	return ctx.Result(200, url)
}
