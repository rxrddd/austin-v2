package data

import (
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	"net/http"

	goodsClientV1 "github.com/ZQCard/kratos-base-project/api/goods/v1"
)

type goodsRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func (g goodsRepo) CreateGoods(ctx context.Context, reqData *biz.Goods) (*biz.Goods, error) {
	reply, err := g.data.goodsClient.CreateGoods(ctx, &goodsClientV1.CreateGoodsRequest{
		Name:   reqData.Name,
		Style:  reqData.Style,
		Size:   reqData.Size,
		Code:   reqData.Code,
		Price:  reqData.Price,
		Number: reqData.Number,
	})

	if err != nil {
		return nil, err
	}

	return &biz.Goods{
		Id:        reply.Id,
		Name:      reply.Name,
		Style:     reply.Style,
		Size:      reply.Size,
		Code:      reply.Code,
		Price:     reply.Price,
		Number:    reply.Number,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
		DeletedAt: reply.DeletedAt,
	}, nil

}

func (g goodsRepo) UpdateGoods(ctx context.Context, reqData *biz.Goods) (*biz.Goods, error) {
	result, err, _ := g.sg.Do(fmt.Sprintf("update_goods_by_id_%s", reqData.Id), func() (interface{}, error) {
		reply, err := g.data.goodsClient.UpdateGoods(ctx, &goodsClientV1.UpdateGoodsRequest{
			Id:     reqData.Id,
			Name:   reqData.Name,
			Style:  reqData.Style,
			Size:   reqData.Size,
			Code:   reqData.Code,
			Price:  reqData.Price,
			Number: reqData.Number,
		})

		if err != nil {
			return nil, err
		}
		return &biz.Goods{
			Id:        reply.Id,
			Name:      reply.Name,
			Style:     reply.Style,
			Size:      reply.Size,
			Code:      reply.Code,
			Price:     reply.Price,
			Number:    reply.Number,
			CreatedAt: reply.CreatedAt,
			UpdatedAt: reply.UpdatedAt,
			DeletedAt: reply.DeletedAt,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Goods), nil
}

func (g goodsRepo) GetGoods(ctx context.Context, id int64) (*biz.Goods, error) {
	result, err, _ := g.sg.Do(fmt.Sprintf("get_goods_by_id_%s", id), func() (interface{}, error) {
		reply, err := g.data.goodsClient.GetGoods(ctx, &goodsClientV1.GetGoodsRequest{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		return &biz.Goods{
			Id:        reply.Id,
			Name:      reply.Name,
			Style:     reply.Style,
			Size:      reply.Size,
			Code:      reply.Code,
			Price:     reply.Price,
			Number:    reply.Number,
			CreatedAt: reply.CreatedAt,
			UpdatedAt: reply.UpdatedAt,
			DeletedAt: reply.DeletedAt,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.Goods), nil
}

func (g goodsRepo) ListGoods(ctx context.Context, pageNum, pageSize int64, goods *biz.Goods) ([]*biz.Goods, int64, error) {
	var res []*biz.Goods
	reply, err := g.data.goodsClient.ListGoods(ctx, &goodsClientV1.ListGoodsRequest{
		PageNum:   pageNum,
		PageSize:  pageSize,
		Id:        goods.Id,
		Name:      goods.Name,
		Style:     goods.Style,
		Size:      goods.Size,
		Code:      goods.Code,
		DeletedAt: goods.DeletedAt,
	})
	if err != nil {
		return res, 0, err
	}

	for _, v := range reply.List {
		tmp := &biz.Goods{
			Id:        v.Id,
			Name:      v.Name,
			Style:     v.Style,
			Size:      v.Size,
			Code:      v.Code,
			Price:     v.Price,
			Number:    v.Number,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
		res = append(res, tmp)
	}
	return res, reply.Total, nil
}

func (g goodsRepo) DeleteGoods(ctx context.Context, id int64) error {
	result, err, _ := g.sg.Do(fmt.Sprintf("delete_goods_by_id_%s", id), func() (interface{}, error) {
		reply, err := g.data.goodsClient.DeleteGoods(ctx, &goodsClientV1.DeleteGoodsRequest{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		return reply.IsSuccess, nil
	})
	if err != nil {
		return err
	}
	if result.(bool) != true {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "删除失败")
	}
	return nil
}

func (g goodsRepo) SaleGoods(ctx context.Context, id int64, number int64) error {
	result, err, _ := g.sg.Do(fmt.Sprintf("sale_goods_by_id_%s", id), func() (interface{}, error) {
		reply, err := g.data.goodsClient.SaleGoods(ctx, &goodsClientV1.SaleGoodsRequest{
			Id:     id,
			Number: number,
		})
		if err != nil {
			return nil, err
		}
		return reply.IsSuccess, nil
	})
	if err != nil {
		return err
	}
	if result.(bool) != true {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "出售失败")
	}
	return nil
}

func (g goodsRepo) RecoverGoods(ctx context.Context, id int64) error {
	result, err, _ := g.sg.Do(fmt.Sprintf("recover_goods_by_id_%s", id), func() (interface{}, error) {
		reply, err := g.data.goodsClient.RecoverGoods(ctx, &goodsClientV1.RecoverGoodsRequest{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		return reply.IsSuccess, nil
	})
	if err != nil {
		return err
	}
	if result.(bool) != true {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "恢复失败")
	}
	return nil
}

func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/goods")),
		sg:   &singleflight.Group{},
	}
}
