package service

import (
	"context"
	v1 "github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminInterface) GetOssStsToken(ctx context.Context, pb *emptypb.Empty) (*v1.OssStsTokenResponse, error) {
	reply, err := s.filesRepo.GetOssStsToken(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
