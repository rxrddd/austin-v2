package service

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-manager/internal/biz"
	"context"
	"github.com/spf13/cast"

	pb "austin-v2/api/msgpusher-manager/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MsgPusherManagerService struct {
	pb.UnimplementedMsgPusherManagerServer
	mr  *biz.MsgRecordUseCase
	mt  *biz.MessageTemplateUseCase
	suc *biz.SmsRecordUseCase
	sa  *biz.SendAccountUseCase
}

func NewMsgPusherManagerService(
	mr *biz.MsgRecordUseCase,
	mt *biz.MessageTemplateUseCase,
	suc *biz.SmsRecordUseCase,
	sa *biz.SendAccountUseCase,
) *MsgPusherManagerService {
	return &MsgPusherManagerService{
		mr:  mr,
		mt:  mt,
		suc: suc,
		sa:  sa,
	}
}

func (s *MsgPusherManagerService) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	return s.sa.SendAccountEdit(ctx, req)
}
func (s *MsgPusherManagerService) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.sa.SendAccountChangeStatus(ctx, req)
}
func (s *MsgPusherManagerService) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	return s.sa.SendAccountList(ctx, req)
}
func (s *MsgPusherManagerService) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	return s.sa.SendAccountQuery(ctx, req)
}
func (s *MsgPusherManagerService) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateEdit(ctx, req)
}
func (s *MsgPusherManagerService) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateChangeStatus(ctx, req)
}
func (s *MsgPusherManagerService) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	return s.mt.TemplateList(ctx, req)
}
func (s *MsgPusherManagerService) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateRemove(ctx, req)
}
func (s *MsgPusherManagerService) GetAllChannel(ctx context.Context, req *emptypb.Empty) (*pb.GetAllChannelResp, error) {
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
func (s *MsgPusherManagerService) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {
	return s.suc.GetSmsRecord(ctx, req)
}
func (s *MsgPusherManagerService) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	return s.mr.GetMsgRecord(ctx, req)
}
