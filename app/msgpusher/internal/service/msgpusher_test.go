package service

import (
	msgpusherV1 "austin-v2/api/msgpusher/v1"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	etcdclient "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestMsgPusherService_Send(t *testing.T) {
	r := NewDiscovery()
	client := NewMsgPusherServiceClient(r)
	newStruct, _ := structpb.NewStruct(map[string]interface{}{
		"content": "恭喜你:{$content}",
	})
	send, err := client.Send(context.Background(), &msgpusherV1.SendRequest{
		Code:              "send",
		MessageTemplateId: 1,
		MessageParam: &msgpusherV1.MessageParam{
			Receiver:  "13541514612",
			Variables: newStruct,
			Extra:     newStruct,
		},
	})
	fmt.Println(`send, err`, send, err)
}

func NewMsgPusherServiceClient(r registry.Discovery) msgpusherV1.MsgPusherClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///austin-v2.msgpusher.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := msgpusherV1.NewMsgPusherClient(conn)
	return c
}
func NewDiscovery() registry.Discovery {
	point := "127.0.0.1:2379"
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}
