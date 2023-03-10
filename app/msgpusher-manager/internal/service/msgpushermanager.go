package service

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-manager/internal/biz"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"context"
	"github.com/spf13/cast"
	"time"

	pb "austin-v2/api/msgpusher-manager/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MsgPusherManagerService struct {
	pb.UnimplementedMsgPusherManagerServer
	mr       *biz.MsgRecordUseCase
	mt       *biz.MessageTemplateUseCase
	suc      *biz.SmsRecordUseCase
	sa       *biz.SendAccountUseCase
	wxTempUc *biz.WxTemplateUseCase
}

func NewMsgPusherManagerService(
	mr *biz.MsgRecordUseCase,
	mt *biz.MessageTemplateUseCase,
	suc *biz.SmsRecordUseCase,
	sa *biz.SendAccountUseCase,
	wxTempUc *biz.WxTemplateUseCase,
) *MsgPusherManagerService {
	return &MsgPusherManagerService{
		mr:       mr,
		mt:       mt,
		suc:      suc,
		sa:       sa,
		wxTempUc: wxTempUc,
	}
}

func (s *MsgPusherManagerService) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	return s.sa.SendAccountEdit(ctx, &domain.SendAccountEditRequest{
		ID:          req.ID,
		SendChannel: req.SendChannel,
		Config:      req.Config,
		Title:       req.Title,
	})
}
func (s *MsgPusherManagerService) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.sa.SendAccountChangeStatus(ctx, req.ID, int(req.Status))
}
func (s *MsgPusherManagerService) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	result, err := s.sa.SendAccountList(ctx, &domain.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.SendAccountListResp{
		Total: result.Total,
	}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
			Status:      int64(item.Status),
		})
	}

	return response, err
}
func (s *MsgPusherManagerService) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	result, err := s.sa.SendAccountList(ctx, &domain.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.SendAccountQueryResp{}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}

	return response, err
}
func (s *MsgPusherManagerService) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateEdit(ctx, &domain.TemplateEditRequest{
		ID:                  req.ID,
		Name:                req.Name,
		IDType:              int(req.IdType),
		SendChannel:         int(req.SendChannel),
		TemplateType:        int(req.TemplateType),
		TemplateSn:          req.TemplateSn,
		MsgType:             int(req.MsgType),
		ShieldType:          int(req.ShieldType),
		MsgContent:          req.MsgContent,
		SendAccount:         req.SendAccount,
		SmsChannel:          req.SmsChannel,
		Updated:             time.Now().Unix(),
		DeduplicationConfig: req.DeduplicationConfig,
	})
}
func (s *MsgPusherManagerService) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateChangeStatus(ctx, req.ID, int(req.Status))
}
func (s *MsgPusherManagerService) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	result, err := s.mt.TemplateList(ctx, &domain.TemplateListRequest{
		Name:        req.Name,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.TemplateListResp{
		Total: result.Total,
	}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.TemplateListRow{
			ID:                  item.ID,
			Name:                item.Name,
			IdType:              item.IdType,
			SendChannel:         item.SendChannel,
			TemplateType:        item.TemplateType,
			MsgType:             item.MsgType,
			ShieldType:          item.ShieldType,
			MsgContent:          item.MsgContent,
			SendAccount:         item.SendAccount,
			SendAccountName:     item.SendAccountName,
			TemplateSn:          item.TemplateSn,
			SmsChannel:          item.SmsChannel,
			CreateAt:            item.CreateAt,
			DeduplicationConfig: item.DeduplicationConfig,
		})
	}

	return response, nil
}
func (s *MsgPusherManagerService) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateRemove(ctx, req.ID)
}
func (s *MsgPusherManagerService) TemplateOne(ctx context.Context, req *pb.TemplateOneRequest) (*pb.TemplateOneResp, error) {
	one, err := s.mt.TemplateOne(ctx, &domain.TemplateOneRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TemplateOneResp{
		Id:                  one.ID,
		Name:                one.Name,
		IdType:              int32(one.IDType),
		SendChannel:         int32(one.SendChannel),
		TemplateType:        int32(one.TemplateType),
		TemplateSn:          one.TemplateSn,
		MsgType:             int32(one.MsgType),
		ShieldType:          int32(one.ShieldType),
		MsgContent:          one.MsgContent,
		SendAccount:         one.SendAccount,
		Creator:             one.Creator,
		Updator:             one.Updator,
		Auditor:             one.Auditor,
		Team:                one.Team,
		Proposer:            one.Proposer,
		SmsChannel:          one.SmsChannel,
		DeduplicationConfig: one.DeduplicationConfig,
	}, nil
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
	return s.suc.GetSmsRecord(ctx, &domain.SmsRecordRequest{
		TemplateId:  req.TemplateId,
		RequestId:   req.RequestId,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
}
func (s *MsgPusherManagerService) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	result, err := s.mr.GetMsgRecord(ctx, &domain.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.MsgRecordResp{
		Total: result.Total,
	}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.MsgRecordRow{
			MessageTemplateId: item.MessageTemplateId,
			RequestId:         item.RequestId,
			Receiver:          item.Receiver,
			MsgId:             item.MsgId,
			Channel:           item.Channel,
			Msg:               item.Msg,
			SendAt:            item.SendAt,
			CreateAt:          item.CreateAt,
			SendSinceTime:     item.SendSinceTime,
			ID:                item.ID,
		})
	}

	return response, nil
}

func (s *MsgPusherManagerService) GetOfficialAccountTemplateList(ctx context.Context, req *pb.OfficialAccountTemplateRequest) (*pb.OfficialAccountTemplateResp, error) {
	result, err := s.wxTempUc.GetOfficialAccountTemplateList(ctx, &domain.OfficialAccountTemplateRequest{
		SendAccount: req.SendAccount,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.OfficialAccountTemplateResp{}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.OfficialAccountTemplateRow{
			TemplateID: item.TemplateID,
			Title:      item.Title,
			Content:    item.Content,
			Example:    item.Example,
		})
	}

	return response, nil
}
