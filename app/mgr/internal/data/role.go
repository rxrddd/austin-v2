package data

import (
	"austin-v2/app/mgr/internal/domain"
	"austin-v2/common/dal/model"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/errorx"
	"austin-v2/utils/gormHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"gorm.io/gen/field"
	"time"
)

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

type IRoleRepo interface {
	ListPage(ctx context.Context, req *domain.RoleListReq) (m []*model.LaSystemAuthRole, total int32, err error)
	ListAll(ctx context.Context) (m []*model.LaSystemAuthRole, err error)
	MemberCountMap(ctx context.Context, roleIds []int32) (m map[int32]int32, err error)
	Save(ctx context.Context, m *model.LaSystemAuthRole) (err error)
	Disable(ctx context.Context, id int32, status int32) (err error)
	Detail(ctx context.Context, id int32) (m *model.LaSystemAuthRole, err error)
}

func NewRoleRepo(data *Data, logger log.Logger) IRoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/UserRepo")),
	}
}

func (r RoleRepo) ListPage(ctx context.Context, req *domain.RoleListReq) (m []*model.LaSystemAuthRole, total int32, err error) {
	sar := r.data.Query(ctx).LaSystemAuthRole
	qx := sar.Where(sar.IsDisable.Eq(0)).
		Order(sar.Sort.Desc())

	if emptyHelper.IsNotEmpty(req.Keywords) {
		qx = qx.Where(sar.Name.Like("%" + req.Keywords + "%"))
	}
	m = make([]*model.LaSystemAuthRole, 0)
	count, err := qx.Count()
	if err != nil || count <= 0 {
		return m, 0, err
	}
	total = cast.ToInt32(count)
	res, err := qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Find()
	return res, total, err
}

func (r RoleRepo) ListAll(ctx context.Context) (m []*model.LaSystemAuthRole, err error) {
	q := r.data.Query(ctx).LaSystemAuthRole
	m, err = q.Where(q.IsDisable.Eq(0)).
		Order(q.Sort.Desc()).
		Find()
	if err != nil {
		return m, err
	}
	return m, err
}

func (r RoleRepo) Save(ctx context.Context, m *model.LaSystemAuthRole) (err error) {
	q := r.data.Query(ctx).LaSystemAuthRole
	if m.ID > 0 {
		if count, _ := q.
			Where(q.ID.Neq(m.ID), q.Name.Eq(m.Name)).
			Count(); count > 0 {
			return errorx.NewBizErr(fmt.Sprintf("[%s]已存在", m.Name))
		}

		var ulist = []field.AssignExpr{
			q.Name.Value(m.Name),
			q.Remark.Value(m.Remark),
			q.Sort.Value(m.Sort),
			q.IsDisable.Value(m.IsDisable),
			q.MenuIds.Value(m.MenuIds),
			q.UpdateTime.Value(time.Now().Unix()),
		}
		_, err = q.Where(q.ID.Eq(m.ID)).UpdateSimple(ulist...)
	} else {
		if count, _ := q.
			Where(q.Name.Eq(m.Name)).
			Count(); count > 0 {
			return errorx.NewBizErr(fmt.Sprintf("[%s]已存在", m.Name))
		}
		m.CreateTime = time.Now().Unix()
		err = q.
			Create(m)
	}
	return err
}

func (r RoleRepo) Disable(ctx context.Context, id int32, status int32) (err error) {
	q := r.data.Query(ctx).LaSystemAuthRole
	_, err = q.
		Where(q.ID.Eq(id)).
		UpdateSimple(q.IsDisable.Value(status))
	return err
}
func (r RoleRepo) MemberCountMap(ctx context.Context, ids []int32) (m map[int32]int32, err error) {
	q := r.data.Query(ctx).LaSystemAdminRole
	var res = make([]map[string]any, 0)
	m = make(map[int32]int32)
	err = q.
		Select(q.RoleID, q.AdminID.Count().As("count")).
		Where(q.RoleID.In(ids...)).
		Group(q.RoleID).
		Scan(&res)
	for _, re := range res {
		m[cast.ToInt32(re["role_id"])] = cast.ToInt32(re["count"])
	}
	return m, err
}
func (r RoleRepo) Detail(ctx context.Context, id int32) (m *model.LaSystemAuthRole, err error) {
	q := r.data.Query(ctx).LaSystemAuthRole
	m, err = q.
		Where(q.ID.Eq(id)).
		First()
	return m, err
}
