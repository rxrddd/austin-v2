package service

import (
	pb "austin-v2/api/mgr"
	"austin-v2/app/mgr/internal/data"
	"austin-v2/common/enums/channelType"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GmpPlatformService struct {
	msgPusherManagerRepo *data.MsgPusherManagerRepo
}

func NewGmpPlatformService(
	msgPusherManagerRepo *data.MsgPusherManagerRepo,
) *GmpPlatformService {
	return &GmpPlatformService{msgPusherManagerRepo: msgPusherManagerRepo}
}

func (s *GmpPlatformService) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.SendAccountEdit(ctx, req)
}
func (s *GmpPlatformService) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.SendAccountChangeStatus(ctx, req)
}
func (s *GmpPlatformService) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	return s.msgPusherManagerRepo.SendAccountList(ctx, req)
}
func (s *GmpPlatformService) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	return s.msgPusherManagerRepo.SendAccountQuery(ctx, req)
}
func (s *GmpPlatformService) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateEdit(ctx, req)
}
func (s *GmpPlatformService) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateChangeStatus(ctx, req)
}
func (s *GmpPlatformService) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	return s.msgPusherManagerRepo.TemplateList(ctx, req)
}
func (s *GmpPlatformService) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateRemove(ctx, req)
}

func (s *GmpPlatformService) TemplateOne(ctx context.Context, req *pb.TemplateOneRequest) (*pb.TemplateOneResp, error) {
	return s.msgPusherManagerRepo.TemplateOne(ctx, req)
}

func (s *GmpPlatformService) GetAllChannel(ctx context.Context, req *emptypb.Empty) (*pb.GetAllChannelResp, error) {
	var res []*pb.Channel
	for key, val := range channelType.TypeEnCode {
		res = append(res, &pb.Channel{
			Id:      int32(val),
			Name:    channelType.TypeText[val],
			Channel: key,
		})
	}
	return &pb.GetAllChannelResp{
		Items: res,
	}, nil
}

func (s *GmpPlatformService) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {
	return s.msgPusherManagerRepo.GetSmsRecord(ctx, req)
}
func (s *GmpPlatformService) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	return s.msgPusherManagerRepo.GetMsgRecord(ctx, req)
}
func (s *GmpPlatformService) GetOfficialAccountTemplateList(ctx context.Context, req *pb.OfficialAccountTemplateRequest) (*pb.OfficialAccountTemplateResp, error) {
	return s.msgPusherManagerRepo.GetOfficialAccountTemplateList(ctx, req)
}
