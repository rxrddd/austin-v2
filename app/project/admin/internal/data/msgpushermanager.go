package data

import (
	msgpushermanagerV1 "austin-v2/api/msgpusher-manager/v1"
	pb "austin-v2/api/project/admin/v1"
	"austin-v2/app/project/admin/internal/conf"
	"austin-v2/app/project/admin/pkg/ctxdata"
	"austin-v2/pkg/utils/metaHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	metadataMidd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewMsgPusherManagerClient(_ *conf.Auth, sr *conf.Service, r registry.Discovery) msgpushermanagerV1.MsgPusherManagerClient {
	// 初始化auth配置
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Msgpushermanager.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(sr.Msgpushermanager.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			func(handler middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					return handler(metaHelper.WithContext(ctx, metaHelper.LoginUser{
						UserId:   ctxdata.GetAdminId(ctx),
						UserName: ctxdata.GetAdminName(ctx),
					}), req)
				}
			},
			metadataMidd.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := msgpushermanagerV1.NewMsgPusherManagerClient(conn)
	return c
}

func NewMsgPusherManagerRepo(data *Data, logger log.Logger) *MsgPusherManagerRepo {
	return &MsgPusherManagerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}

type MsgPusherManagerRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (s *MsgPusherManagerRepo) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.SendAccountEdit(ctx, &msgpushermanagerV1.SendAccountEditRequest{
		ID:          req.Id,
		Title:       req.Title,
		Config:      req.Config,
		SendChannel: req.SendChannel,
	})
}
func (s *MsgPusherManagerRepo) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.SendAccountChangeStatus(ctx, &msgpushermanagerV1.SendAccountChangeStatusRequest{
		ID:     req.Id,
		Status: req.Status,
	})
}
func (s *MsgPusherManagerRepo) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	list, err := s.data.msgPusherManagerClient.SendAccountList(ctx, &msgpushermanagerV1.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})

	if err != nil {
		return nil, err
	}
	resp := make([]*pb.SendAccountRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.SendAccountRow{
			Id:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
			Status:      item.Status,
		})
	}
	return &pb.SendAccountListResp{
		Rows:  resp,
		Total: list.Total,
	}, nil

}
func (s *MsgPusherManagerRepo) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	list, err := s.data.msgPusherManagerClient.SendAccountQuery(ctx, &msgpushermanagerV1.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
	})

	if err != nil {
		return nil, err
	}
	resp := make([]*pb.SendAccountRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.SendAccountRow{
			Id:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}
	return &pb.SendAccountQueryResp{
		Rows: resp,
	}, nil

}
func (s *MsgPusherManagerRepo) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.TemplateEdit(ctx, &msgpushermanagerV1.TemplateEditRequest{
		ID:                  req.Id,
		Name:                req.Name,
		IdType:              req.IdType,
		SendChannel:         req.SendChannel,
		TemplateType:        req.TemplateType,
		MsgType:             req.MsgType,
		ShieldType:          req.ShieldType,
		MsgContent:          req.MsgContent,
		SendAccount:         req.SendAccount,
		TemplateSn:          req.TemplateSn,
		SmsChannel:          req.SmsChannel,
		DeduplicationConfig: req.DeduplicationConfig,
	})
}
func (s *MsgPusherManagerRepo) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.TemplateChangeStatus(ctx, &msgpushermanagerV1.TemplateChangeStatusRequest{
		ID:     req.Id,
		Status: req.Status,
	})
}
func (s *MsgPusherManagerRepo) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	list, err := s.data.msgPusherManagerClient.TemplateList(ctx, &msgpushermanagerV1.TemplateListRequest{
		Name:        req.Name,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.TemplateListRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.TemplateListRow{
			Id:              item.ID,
			Name:            item.Name,
			IdType:          item.IdType,
			SendChannel:     item.SendChannel,
			TemplateType:    item.TemplateType,
			MsgType:         item.MsgType,
			ShieldType:      item.ShieldType,
			MsgContent:      item.MsgContent,
			SendAccount:     item.SendAccount,
			SendAccountName: item.SendAccountName,
			TemplateSn:      item.TemplateSn,
			SmsChannel:      item.SmsChannel,
			CreateAt:        item.CreateAt,
		})
	}
	return &pb.TemplateListResp{
		Rows:  resp,
		Total: list.Total,
	}, nil
}
func (s *MsgPusherManagerRepo) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.TemplateRemove(ctx, &msgpushermanagerV1.TemplateRemoveRequest{
		ID: req.Id,
	})
}
func (s *MsgPusherManagerRepo) GetAllChannel(ctx context.Context, req *emptypb.Empty) (*pb.GetAllChannelResp, error) {
	list, err := s.data.msgPusherManagerClient.GetAllChannel(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.Channel, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.Channel{
			Id:      item.Id,
			Name:    item.Name,
			Channel: item.Channel,
		})
	}

	return &pb.GetAllChannelResp{
		Rows: resp,
	}, nil
}

func (s *MsgPusherManagerRepo) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {

	list, err := s.data.msgPusherManagerClient.GetSmsRecord(ctx, &msgpushermanagerV1.SmsRecordRequest{
		TemplateId:  req.TemplateId,
		RequestId:   req.RequestId,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.SmsRecordRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.SmsRecordRow{
			Id:                item.Id,
			SeriesId:          item.SeriesId,
			MsgContent:        item.MsgContent,
			SupplierName:      item.SupplierName,
			SupplierId:        item.SupplierId,
			Phone:             item.Phone,
			MessageTemplateId: item.MessageTemplateId,
			CreatedAt:         item.CreatedAt,
			SendDate:          item.SendDate,
			Status:            item.Status,
			ReportContent:     item.ReportContent,
			ChargingNum:       item.ChargingNum,
			UpdatedAt:         item.UpdatedAt,
		})
	}

	return &pb.SmsRecordResp{
		Rows:  resp,
		Total: list.Total,
	}, nil

}
func (s *MsgPusherManagerRepo) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {

	list, err := s.data.msgPusherManagerClient.GetMsgRecord(ctx, &msgpushermanagerV1.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.MsgRecordRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.MsgRecordRow{
			Id:                item.ID,
			MessageTemplateId: item.MessageTemplateId,
			RequestId:         item.RequestId,
			Receiver:          item.Receiver,
			MsgId:             item.MsgId,
			Channel:           item.Channel,
			Msg:               item.Msg,
			SendAt:            item.SendAt,
			CreateAt:          item.CreateAt,
			SendSinceTime:     item.SendSinceTime,
		})
	}

	return &pb.MsgRecordResp{
		Rows:  resp,
		Total: list.Total,
	}, nil

}
