package service

import (
	"austin-v2/utils/timeHelper"
	"context"

	pb "austin-v2/api/mgr"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommonService struct {
	pb.UnimplementedCommonServer
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (s *CommonService) GetWorkbench(ctx context.Context, req *emptypb.Empty) (*pb.ConsoleReply, error) {
	return &pb.ConsoleReply{
		Today: &pb.ConsoleToday{
			Time:        timeHelper.CurrentTimeYMDHIS(),
			TodayOrder:  12,
			TodaySales:  30,
			TodayUsers:  120,
			TodayVisits: 10,
			TotalOrder:  255,
			TotalSales:  65,
			TotalUsers:  360,
			TotalVisits: 100,
		},
		Version: &pb.ConsoleVersion{
			Based:   "Vue3.x、ElementUI、MySQL",
			Name:    "LikeAdmin开源后台",
			Version: "v1.0.1",
			Website: "www.likeadmin.cn",
		},
		Visitor: &pb.ConsoleVisitor{
			Date: []string{
				"2023-03-01",
				"2023-03-02",
				"2023-03-03",
				"2023-03-04",
				"2023-03-05",
				"2023-03-06",
				"2023-03-07",
				"2023-03-08",
				"2023-03-09",
				"2023-03-10",
				"2023-03-11",
				"2023-03-12",
				"2023-03-13",
				"2023-03-14",
				"2023-03-15",
				"2023-03-16",
				"2023-03-17",
				"2023-03-18",
				"2023-03-19",
				"2023-03-20",
				"2023-03-21",
				"2023-03-22",
				"2023-03-23",
				"2023-03-24",
				"2023-03-25",
				"2023-03-26",
				"2023-03-27",
				"2023-03-28",
				"2023-03-29",
				"2023-03-30",
			},
			List: []int32{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
				11,
				12,
				13,
				14,
				15,
				16,
				17,
				18,
				19,
				20,
				11,
				22,
				33,
				44,
				55,
				66,
				77,
				88,
				89,
				90,
			},
		},
	}, nil
}
func (s *CommonService) GetIndexConfig(ctx context.Context, req *emptypb.Empty) (*pb.IndexConfigReply, error) {
	return &pb.IndexConfigReply{
		Copyright: []*pb.IndexConfigReply_Copyright{
			{
				Link: "http://www.beian.gov.cn",
				Name: "LikeAdmin开源系统",
			},
		},
		OssDomain:   "https://go-admin.likeadmin.cn",
		WebBackdrop: "https://go-admin.likeadmin.cn/api/static/backend_backdrop.png",
		WebFavicon:  "https://go-admin.likeadmin.cn/api/static/backend_favicon.ico",
		WebLogo:     "https://go-admin.likeadmin.cn/api/static/backend_logo.png",
		WebName:     "LikeAdmin开源后台1",
	}, nil
}
