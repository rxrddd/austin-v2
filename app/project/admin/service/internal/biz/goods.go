package biz

import (
	"context"
	"net/http"

	"github.com/ZQCard/kratos-base-project/pkg/validate"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Id        int64
	Name      string  `validate:"required,max=20" label:"商品名称"`
	Style     string  `validate:"required,max=20" label:"商品款式"`
	Size      string  `validate:"required,max=20" label:"商品尺码"`
	Code      string  `validate:"required,max=20" label:"商品序列号"`
	Price     float32 `validate:"required,numeric,gt=0" label:"商品价格"`
	Number    int64   `validate:"required,numeric,gte=1" label:"商品库存"`
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type GoodsLog struct {
	Id        int64
	Name      string
	Style     string
	Size      string
	Code      string
	Price     float32
	Number    int64
	CreatedAt string
}

// GoodsRepo 模块接口
type GoodsRepo interface {
	CreateGoods(ctx context.Context, reqData *Goods) (*Goods, error)
	UpdateGoods(ctx context.Context, reqData *Goods) (*Goods, error)
	GetGoods(ctx context.Context, id int64) (*Goods, error)
	ListGoods(ctx context.Context, pageNum, pageSize int64, goods *Goods) ([]*Goods, int64, error)
	DeleteGoods(ctx context.Context, id int64) error
	RecoverGoods(ctx context.Context, id int64) error
	SaleGoods(ctx context.Context, id int64, number int64) error
	SaleGoodsLogList(ctx context.Context, pageNum, pageSize, goodsId int64) ([]*GoodsLog, int64, error)
}

type GoodsUseCase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUseCase(repo GoodsRepo, logger log.Logger) *GoodsUseCase {
	return &GoodsUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/goods"))}
}

func (uc *GoodsUseCase) Create(ctx context.Context, data *Goods) (*Goods, error) {
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Goods{}, errors.New(http.StatusBadRequest, "PARAMS_ERROR", err.Error())
	}
	return uc.repo.CreateGoods(ctx, data)
}

func (uc *GoodsUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteGoods(ctx, id)
}

func (uc *GoodsUseCase) Update(ctx context.Context, data *Goods) (*Goods, error) {
	if data.Id == 0 {
		return &Goods{}, errors.New(http.StatusBadRequest, "PARAMS_ERROR", "id不得为空")
	}
	err := validate.ValidateStructCN(data)
	if err != nil {
		return &Goods{}, errors.New(http.StatusBadRequest, "PARAMS_ERROR", err.Error())
	}
	return uc.repo.UpdateGoods(ctx, data)
}

func (uc *GoodsUseCase) Get(ctx context.Context, id int64) (*Goods, error) {
	if id == 0 {
		return &Goods{}, errors.New(http.StatusBadRequest, "PARAMS_ERROR", "id不得为空")
	}
	return uc.repo.GetGoods(ctx, id)
}

func (uc *GoodsUseCase) List(ctx context.Context, pageNum, pageSize int64, goods *Goods) ([]*Goods, int64, error) {
	return uc.repo.ListGoods(ctx, pageNum, pageSize, goods)
}

func (uc *GoodsUseCase) Sale(ctx context.Context, id, number int64) error {
	if id == 0 {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "id不得为空")
	}
	if number == 0 {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "数量不得为空")
	}
	return uc.repo.SaleGoods(ctx, id, number)
}

func (uc *GoodsUseCase) SaleGoodsLogList(ctx context.Context, pageNum, pageSize, goodsId int64) ([]*GoodsLog, int64, error) {
	return uc.repo.SaleGoodsLogList(ctx, pageNum, pageSize, goodsId)
}

func (uc *GoodsUseCase) Recover(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New(http.StatusBadRequest, "PARAMS_ERROR", "id不得为空")
	}
	return uc.repo.RecoverGoods(ctx, id)
}
