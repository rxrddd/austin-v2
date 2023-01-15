package service

import (
	pb "austin-v2/api/project/admin/v1"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"context"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminInterface) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.SendAccountEdit(ctx, req)
}
func (s *AdminInterface) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.SendAccountChangeStatus(ctx, req)
}
func (s *AdminInterface) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	return s.msgPusherManagerRepo.SendAccountList(ctx, req)
}
func (s *AdminInterface) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	return s.msgPusherManagerRepo.SendAccountQuery(ctx, req)
}
func (s *AdminInterface) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateEdit(ctx, req)
}
func (s *AdminInterface) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateChangeStatus(ctx, req)
}
func (s *AdminInterface) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	return s.msgPusherManagerRepo.TemplateList(ctx, req)
}
func (s *AdminInterface) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.msgPusherManagerRepo.TemplateRemove(ctx, req)
}
func (s *AdminInterface) GetAllChannel(ctx context.Context, req *emptypb.Empty) (*pb.GetAllChannelResp, error) {
	var res []*pb.Channel
	for key, val := range channelType.TypeEnCode {
		res = append(res, &pb.Channel{
			Id:      cast.ToInt64(val),
			Name:    channelType.TypeText[val],
			Channel: key,
		})
	}
	return &pb.GetAllChannelResp{
		Rows: res,
	}, nil
}

func (s *AdminInterface) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {
	return s.msgPusherManagerRepo.GetSmsRecord(ctx, req)
}
func (s *AdminInterface) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	return s.msgPusherManagerRepo.GetMsgRecord(ctx, req)
}
