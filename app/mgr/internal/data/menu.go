package data

import (
	"austin-v2/common/dal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen/field"
	"time"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper
}
type IMenuRepo interface {
	MenuListByIds(ctx context.Context, ids []int32) (m []*model.LaSystemAuthMenu, err error)
	MenuListAll(ctx context.Context) (m []*model.LaSystemAuthMenu, err error)
	ListAll(ctx context.Context, ids ...int32) (m []*model.LaSystemAuthMenu, err error)
	FormatMenu(list []*model.LaSystemAuthMenu, pid int32) []*model.LaSystemAuthMenu
	SaveMenu(ctx context.Context, m *model.LaSystemAuthMenu) (err error)
	DeleteMenu(ctx context.Context, id int32) (err error)
	CountChildrenMenu(ctx context.Context, id int32) (count int64, err error)
}

func NewMenuRepo(data *Data, logger log.Logger) IMenuRepo {
	return &MenuRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/UserRepo")),
	}
}

func (r MenuRepo) MenuListByIds(ctx context.Context, ids []int32) (m []*model.LaSystemAuthMenu, err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu

	res, err := q.Where(
		q.ID.In(ids...),
		q.IsShow.Eq(1),
		q.MenuType.Neq("A"),
	).
		Order(q.MenuSort.Desc()).
		Find()
	return res, err
}

func (r MenuRepo) MenuListAll(ctx context.Context) (m []*model.LaSystemAuthMenu, err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu
	res, err := q.Where(
		q.IsShow.Eq(1),
		q.MenuType.Neq("A"),
	).
		Order(q.MenuSort.Desc()).
		Find()
	return res, err
}

func (r MenuRepo) ListAll(ctx context.Context, ids ...int32) (m []*model.LaSystemAuthMenu, err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu
	qx := q.Where(q.IsDisable.Eq(0)).Order(q.MenuSort.Desc())
	if len(ids) > 0 {
		qx = qx.Where(q.ID.In(ids...))
	}
	res, err := qx.
		Find()
	return res, err
}

// FormatMenu 无限极分类
func (r MenuRepo) FormatMenu(list []*model.LaSystemAuthMenu, pid int32) []*model.LaSystemAuthMenu {
	var arr []*model.LaSystemAuthMenu
	for _, v := range list {
		if v.Pid == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := r.FormatMenu(list, v.ID)
			node := &model.LaSystemAuthMenu{
				ID:         v.ID,
				Pid:        v.Pid,
				MenuType:   v.MenuType,
				MenuName:   v.MenuName,
				MenuIcon:   v.MenuIcon,
				MenuSort:   v.MenuSort,
				Perms:      v.Perms,
				Paths:      v.Paths,
				Component:  v.Component,
				Selected:   v.Selected,
				Params:     v.Params,
				IsCache:    v.IsCache,
				IsShow:     v.IsShow,
				IsDisable:  v.IsDisable,
				CreateTime: v.CreateTime,
				UpdateTime: v.UpdateTime,
				Children:   child,
			}
			arr = append(arr, node)
		}
	}
	return arr
}

func (r MenuRepo) SaveMenu(ctx context.Context, m *model.LaSystemAuthMenu) (err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu
	if m.ID > 0 {
		var ulist = []field.AssignExpr{
			q.Pid.Value(m.Pid),
			q.MenuType.Value(m.MenuType),
			q.MenuName.Value(m.MenuName),
			q.MenuIcon.Value(m.MenuIcon),
			q.MenuSort.Value(m.MenuSort),
			q.Perms.Value(m.Perms),
			q.Paths.Value(m.Paths),
			q.Component.Value(m.Component),
			q.Selected.Value(m.Selected),
			q.Params.Value(m.Params),
			q.IsCache.Value(m.IsCache),
			q.IsShow.Value(m.IsShow),
			q.IsDisable.Value(m.IsDisable),
			q.UpdateTime.Value(time.Now().Unix()),
		}
		_, err = q.Where(q.ID.Eq(m.ID)).UpdateSimple(ulist...)
	} else {
		m.CreateTime = time.Now().Unix()
		err = q.
			Create(m)
	}
	return err
}
func (r MenuRepo) DeleteMenu(ctx context.Context, id int32) (err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu
	_, err = q.
		Where(q.ID.Eq(id)).
		Delete()
	return err
}

func (r MenuRepo) CountChildrenMenu(ctx context.Context, id int32) (count int64, err error) {
	q := r.data.Query(ctx).LaSystemAuthMenu
	count, err = q.
		Where(q.Pid.Eq(id)).
		Count()
	return count, err
}
