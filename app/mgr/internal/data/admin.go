package data

import (
	"austin-v2/app/mgr/internal/domain"
	"austin-v2/common/dal/model"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/encryption"
	"austin-v2/utils/errorx"
	"austin-v2/utils/gormHelper"
	"austin-v2/utils/stringHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"gorm.io/gen/field"
	"strings"
	"time"
)

type AdminRepo struct {
	data *Data
	log  *log.Helper
}
type IAdminRepo interface {
	FindUserByPhone(ctx context.Context, username string, extId ...int32) (m *model.LaSystemAuthAdmin, err error)
	FindUserByID(ctx context.Context, id int32) (m *model.LaSystemAuthAdmin, err error)
	ListPage(ctx context.Context, req *domain.AdminListReq) (m []*model.LaSystemAuthAdmin, total int32, err error)
	Save(ctx context.Context, m *model.LaSystemAuthAdmin, roles []string) (err error)
	ChangeStatus(ctx context.Context, id int32, status int32) (err error)
	UpdateInfo(ctx context.Context, id int32, req *domain.UpdateInfoReq) error
}

func NewAdminRepo(data *Data, logger log.Logger) IAdminRepo {
	return &AdminRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/AdminRepo")),
	}
}

func (r AdminRepo) FindUserByPhone(ctx context.Context, username string, extId ...int32) (m *model.LaSystemAuthAdmin, err error) {
	q := r.data.Query(ctx).LaSystemAuthAdmin
	if len(extId) > 0 {
		q.Where(q.ID.Neq(extId[0]))
	}
	res, err := q.Where(field.Or(q.Username.Eq(username), q.Username.Eq(username))).First()
	return res, err
}

func (r AdminRepo) FindUserByID(ctx context.Context, id int32) (m *model.LaSystemAuthAdmin, err error) {
	q := r.data.Query(ctx).LaSystemAuthAdmin
	res, err := q.Where(q.ID.Eq(id)).Preload(q.AuthRoles.Role).First()
	return res, err
}

func (r AdminRepo) ListPage(ctx context.Context, req *domain.AdminListReq) (m []*model.LaSystemAuthAdmin, total int32, err error) {
	u := r.data.Query(ctx).LaSystemAuthAdmin
	qx := u.Where(u.IsDelete.Eq(0)).
		Order(u.IsDisable)
	if emptyHelper.IsNotEmpty(req.Username) {
		qx = qx.Where(u.Username.Like("%" + req.Username + "%"))
	}
	if emptyHelper.IsNotEmpty(req.Nickname) {
		qx = qx.Where(u.Nickname.Like("%" + req.Nickname + "%"))
	}
	if emptyHelper.IsNotEmpty(req.Role) {
		rq := r.data.Query(ctx).LaSystemAdminRole
		split := strings.Split(req.Role, ",")
		var roleIds []int32
		for _, s := range split {
			roleIds = append(roleIds, cast.ToInt32(s))
		}
		qx = qx.Where(
			u.Columns(u.ID).
				In(
					rq.Select(rq.AdminID).
						Where(rq.RoleID.In(roleIds...)),
				),
		)
	}
	m = make([]*model.LaSystemAuthAdmin, 0)
	count, err := qx.Count()
	if err != nil || count <= 0 {
		return m, 0, err
	}
	total = cast.ToInt32(count)
	res, err := qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Preload(u.AuthRoles.Role).
		Find()
	return res, total, err
}

const defaultPwd = "123456"

func (r AdminRepo) Save(ctx context.Context, m *model.LaSystemAuthAdmin, roles []string) (err error) {
	q := r.data.Query(ctx).LaSystemAuthAdmin
	sar := r.data.Query(ctx).LaSystemAdminRole
	if m.ID > 0 {
		if count, _ := q.
			Where(q.Username.Eq(m.Username), q.ID.Neq(m.ID)).
			Count(); count > 0 {
			return errorx.NewBizErr(fmt.Sprintf("用户名[%s]已存在", m.Username))
		}
		var ulist = []field.AssignExpr{
			q.Avatar.Value(m.Avatar),
			q.Username.Value(m.Username),
			q.Nickname.Value(m.Nickname),
			q.UpdateTime.Value(time.Now().Unix()),
		}
		_, err = q.Where(q.ID.Eq(m.ID)).UpdateSimple(ulist...)
		sar.Where(sar.AdminID.Eq(m.ID)).Delete()
	} else {
		if count, _ := q.
			Where(q.Username.Eq(m.Username)).
			Count(); count > 0 {
			return errorx.NewBizErr(fmt.Sprintf("用户名[%s]已存在", m.Username))
		}
		m.Salt = stringHelper.RandString(10)
		m.Password = r.getPwd(defaultPwd, m.Salt)
		m.CreateTime = time.Now().Unix()
		err = q.
			Create(m)
	}
	if err != nil {
		return err
	}
	var arr []*model.LaSystemAdminRole
	for _, v := range roles {
		arr = append(arr, &model.LaSystemAdminRole{
			RoleID:  cast.ToInt32(v),
			AdminID: m.ID,
		})
	}
	return sar.CreateInBatches(arr, 10)
}
func (r AdminRepo) getPwd(pwd, salt string) string {
	return encryption.EncodeMD5(encryption.EncodeMD5(pwd) + salt)
}
func (r AdminRepo) ChangeStatus(ctx context.Context, id int32, status int32) (err error) {
	q := r.data.Query(ctx).LaSystemAuthAdmin
	_, err = q.Where(q.ID.Eq(id)).
		UpdateSimple(q.IsDisable.Value(status))
	return err
}
func (r AdminRepo) UpdateInfo(ctx context.Context, id int32, req *domain.UpdateInfoReq) (err error) {
	q := r.data.Query(ctx).LaSystemAuthAdmin
	var upList = []field.AssignExpr{
		q.Avatar.Value(req.Avatar),
		q.Username.Value(req.Username),
		q.Nickname.Value(req.Nickname),
	}
	if req.Password != "" {
		if req.PasswordConfirm != req.Password {
			return errorx.NewBizErr("两次输入密码不一致!")
		}
		salt := stringHelper.RandString(10)
		upList = append(upList, q.Salt.Value(salt))
		upList = append(upList, q.Password.Value(r.getPwd(req.Password, salt)))
	}
	_, err = q.Where(q.ID.Eq(id)).
		UpdateSimple(upList...)
	return err
}
