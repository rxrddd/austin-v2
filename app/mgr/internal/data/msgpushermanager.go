package data

import (
	pb "austin-v2/api/mgr"
	msgpushermanagerV1 "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/mgr/internal/conf"
	"austin-v2/app/mgr/internal/pkg/ctxdata"
	"austin-v2/utils/jsonHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	metadataMidd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/cast"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewMsgPusherManagerClient(_ *conf.Auth, sr *conf.Service, r registry.Discovery) msgpushermanagerV1.MsgPusherManagerClient {
	fmt.Println(`sr.Msgpushermanager.Endpoint`, sr.Msgpushermanager.Endpoint)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Msgpushermanager.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(sr.Msgpushermanager.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metaUserMiddleware(),
			metadataMidd.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := msgpushermanagerV1.NewMsgPusherManagerClient(conn)
	return c
}

const loginKey = "x-md-global-admin-login"

type LoginUser struct {
	UserId   int32  `json:"user_id"`
	UserName string `json:"user_name"`
}

func metaUserMiddleware() func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = metadata.AppendToClientContext(ctx, loginKey, jsonHelper.MustToString(LoginUser{
				UserId:   ctxdata.MgrID(ctx),
				UserName: ctxdata.MgrName(ctx),
			}))
			return handler(ctx, req)
		}
	}
}

func NewMsgPusherManagerRepo(
	data *Data,
	logger log.Logger,
) *MsgPusherManagerRepo {
	return &MsgPusherManagerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/msgPusherManager")),
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
		Id:          req.Id,
		Title:       req.Title,
		Config:      req.Config,
		SendChannel: req.SendChannel,
	})
}
func (s *MsgPusherManagerRepo) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.SendAccountChangeStatus(ctx, &msgpushermanagerV1.SendAccountChangeStatusRequest{
		Id:     req.Id,
		Status: req.Status,
	})
}
func (s *MsgPusherManagerRepo) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	list, err := s.data.msgPusherManagerClient.SendAccountList(ctx, &msgpushermanagerV1.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
		PageSize:    req.PageSize,
	})

	if err != nil {
		return nil, err
	}
	resp := make([]*pb.SendAccountRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.SendAccountRow{
			Id:          item.Id,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
			Status:      item.Status,
		})
	}
	return &pb.SendAccountListResp{
		Items: resp,
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
			Id:          item.Id,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}
	return &pb.SendAccountQueryResp{
		Items: resp,
	}, nil

}
func (s *MsgPusherManagerRepo) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.TemplateEdit(ctx, &msgpushermanagerV1.TemplateEditRequest{
		Id:                  req.Id,
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
		Id:     req.Id,
		Status: req.Status,
	})
}
func (s *MsgPusherManagerRepo) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	list, err := s.data.msgPusherManagerClient.TemplateList(ctx, &msgpushermanagerV1.TemplateListRequest{
		Name:        req.Name,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.TemplateListRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.TemplateListRow{
			Id:                  item.Id,
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
	return &pb.TemplateListResp{
		Items: resp,
		Total: list.Total,
	}, nil
}
func (s *MsgPusherManagerRepo) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return s.data.msgPusherManagerClient.TemplateRemove(ctx, &msgpushermanagerV1.TemplateRemoveRequest{
		Id: req.Id,
	})
}

func (s *MsgPusherManagerRepo) TemplateOne(ctx context.Context, req *pb.TemplateOneRequest) (*pb.TemplateOneResp, error) {
	one, err := s.data.msgPusherManagerClient.TemplateOne(ctx, &msgpushermanagerV1.TemplateOneRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.TemplateOneResp{
		Id:                  one.Id,
		Name:                one.Name,
		IdType:              one.IdType,
		SendChannel:         one.SendChannel,
		TemplateType:        one.TemplateType,
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
		Items: resp,
	}, nil
}

func (s *MsgPusherManagerRepo) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {

	list, err := s.data.msgPusherManagerClient.GetSmsRecord(ctx, &msgpushermanagerV1.SmsRecordRequest{
		TemplateId:  req.TemplateId,
		RequestId:   req.RequestId,
		SendChannel: req.SendChannel,
		PageNo:      req.PageNo,
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
			CreatedAt:         item.CreateAt,
			SendDate:          item.SendDate,
			Status:            item.Status,
			ReportContent:     item.ReportContent,
			ChargingNum:       item.ChargingNum,
			UpdatedAt:         item.UpdateAt,
		})
	}

	return &pb.SmsRecordResp{
		Items: resp,
		Total: list.Total,
	}, nil

}
func (s *MsgPusherManagerRepo) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {

	list, err := s.data.msgPusherManagerClient.GetMsgRecord(ctx, &msgpushermanagerV1.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.MsgRecordRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.MsgRecordRow{
			Id:                item.Id,
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
		Items: resp,
		Total: list.Total,
	}, nil

}

func (s *MsgPusherManagerRepo) GetOfficialAccountTemplateList(ctx context.Context, req *pb.OfficialAccountTemplateRequest) (*pb.OfficialAccountTemplateResp, error) {
	list, err := s.data.msgPusherManagerClient.GetOfficialAccountTemplateList(ctx, &msgpushermanagerV1.OfficialAccountTemplateRequest{
		SendAccount: cast.ToInt64(req.SendAccount),
	})
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.OfficialAccountTemplateRow, 0)
	for _, item := range list.Rows {
		resp = append(resp, &pb.OfficialAccountTemplateRow{
			TemplateId: item.TemplateId,
			Title:      item.Title,
			Content:    item.Content,
			Example:    item.Example,
		})
	}

	return &pb.OfficialAccountTemplateResp{
		Items: resp,
	}, nil

}
