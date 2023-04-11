package service

import (
	"austin-v2/app/msgpusher-manager/internal/biz"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/enums/channelType"
	"context"
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
		ID:          req.Id,
		SendChannel: req.SendChannel,
		Config:      req.Config,
		Title:       req.Title,
	})
}
func (s *MsgPusherManagerService) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.sa.SendAccountChangeStatus(ctx, req.Id, req.Status)
}
func (s *MsgPusherManagerService) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	result, err := s.sa.SendAccountList(ctx, &domain.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
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
			Id:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
			Status:      item.Status,
		})
	}

	return response, err
}
func (s *MsgPusherManagerService) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	result, err := s.sa.SendAccountList(ctx, &domain.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.SendAccountQueryResp{}
	for _, item := range result.Rows {
		response.Rows = append(response.Rows, &pb.SendAccountRow{
			Id:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}

	return response, err
}
func (s *MsgPusherManagerService) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateEdit(ctx, &domain.TemplateEditRequest{
		ID:                  req.Id,
		Name:                req.Name,
		IDType:              req.IdType,
		SendChannel:         req.SendChannel,
		TemplateType:        req.TemplateType,
		TemplateSn:          req.TemplateSn,
		MsgType:             req.MsgType,
		ShieldType:          req.ShieldType,
		MsgContent:          req.MsgContent,
		SendAccount:         req.SendAccount,
		SmsChannel:          req.SmsChannel,
		UpdateAt:            time.Now().Unix(),
		DeduplicationConfig: req.DeduplicationConfig,
	})
}
func (s *MsgPusherManagerService) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.mt.TemplateChangeStatus(ctx, req.Id, req.Status)
}
func (s *MsgPusherManagerService) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	result, err := s.mt.TemplateList(ctx, &domain.TemplateListRequest{
		Name:        req.Name,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
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
			Id:                  item.ID,
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
	return s.mt.TemplateRemove(ctx, req.Id)
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
		MsgType:             one.MsgType,
		ShieldType:          one.ShieldType,
		MsgContent:          one.MsgContent,
		SendAccount:         one.SendAccount,
		CreateBy:            one.CreateBy,
		UpdateBy:            one.UpdateBy,
		SmsChannel:          one.SmsChannel,
		DeduplicationConfig: one.DeduplicationConfig,
	}, nil
}

func (s *MsgPusherManagerService) GetAllChannel(ctx context.Context, req *emptypb.Empty) (*pb.GetAllChannelResp, error) {
	var res []*pb.Channel
	for key, val := range channelType.TypeEnCode {
		res = append(res, &pb.Channel{
			Id:      int32(val),
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
		Page:        req.PageNo,
		PageSize:    req.PageSize,
	})
}
func (s *MsgPusherManagerService) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	result, err := s.mr.GetMsgRecord(ctx, &domain.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		PageNo:     req.PageNo,
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
			Id:                item.ID,
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
			TemplateId: item.TemplateID,
			Title:      item.Title,
			Content:    item.Content,
			Example:    item.Example,
		})
	}

	return response, nil
}
