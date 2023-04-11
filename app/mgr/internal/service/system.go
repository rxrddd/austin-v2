package service

import (
	"austin-v2/app/mgr/internal/conf"
	"austin-v2/app/mgr/internal/data"
	"austin-v2/app/mgr/internal/domain"
	"austin-v2/app/mgr/internal/pkg/ctxdata"
	"austin-v2/common/dal/model"
	"austin-v2/pkg/transaction"
	"austin-v2/utils/encryption"
	"austin-v2/utils/errorx"
	"austin-v2/utils/gormHelper"
	"austin-v2/utils/sliceHelper"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"strings"
	"time"

	pb "austin-v2/api/mgr"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SystemService struct {
	pb.UnimplementedSystemServer

	cfg       *conf.Bootstrap
	adminRepo data.IAdminRepo
	menuRepo  data.IMenuRepo
	roleRepo  data.IRoleRepo
	tranMgr   transaction.ITranMgr
}

func NewSystemService(
	cfg *conf.Bootstrap,
	adminRepo data.IAdminRepo,
	menuRepo data.IMenuRepo,
	roleRepo data.IRoleRepo,
) *SystemService {
	return &SystemService{
		cfg:       cfg,
		adminRepo: adminRepo,
		menuRepo:  menuRepo,
		roleRepo:  roleRepo,
	}
}

func (s *SystemService) AdminLogin(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {

	m, err := s.adminRepo.FindUserByPhone(ctx, req.Username)

	if gormHelper.IsErrRecordNotFound(err) {
		return nil, errorx.NewBizErr("账号、密码错误")
	}

	if err != nil {
		return nil, err
	}

	if m.Password != encryption.EncodeMD5(encryption.EncodeMD5(req.Password)+m.Salt) {
		return nil, errorx.NewBizErr("账号、密码错误")
	}
	return &pb.LoginResp{
		Token: s.genToken(ctx, m.ID, m.Username),
	}, nil
}

func (s *SystemService) AdminLogout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		data.RedisCli.Del(ctx, tr.RequestHeader().Get("X-Token"))
	}
	return &emptypb.Empty{}, nil
}

func (s *SystemService) GetSelfInfo(ctx context.Context, req *emptypb.Empty) (*pb.SelfReply, error) {
	m, err := s.adminRepo.FindUserByID(ctx, ctxdata.MgrID(ctx))
	if gormHelper.IsErrRecordNotFound(err) {
		return nil, errors.New(401, "UNAUTHORIZED_INFO_MISSING", "授权已过期或授权异常,请重新授权！")
	}
	var roleList []string
	var menuIds []string
	if len(m.AuthRoles) > 0 {
		for _, v := range m.AuthRoles {
			if v.Role != nil {
				roleList = append(roleList, v.Role.Name)
				menuIds = append(menuIds, v.Role.MenuIds...)
			}
		}
	}

	var auths []string
	if sliceHelper.InArr("*", menuIds) {
		auths = []string{"*"}
	} else {
		var permIds []int32
		for _, v := range menuIds {
			permIds = append(permIds, cast.ToInt32(v))
		}
		if len(permIds) > 0 {
			all, _ := s.menuRepo.ListAll(ctx, permIds...)
			for _, v := range all {
				if v.Perms != "" {
					auths = append(auths, v.Perms)
				}
			}
		}
	}
	return &pb.SelfReply{
		User: &pb.UserInfo{
			Id:            m.ID,
			Username:      m.Username,
			Nickname:      m.Nickname,
			Avatar:        m.Avatar,
			Role:          strings.Join(roleList, ","),
			Dept:          "",
			IsMultipoint:  1,
			IsDisable:     m.IsDisable,
			LastLoginIp:   m.LastLoginIP,
			LastLoginTime: m.LastLoginTime,
			CreateTime:    m.CreateTime,
			UpdateTime:    m.UpdateTime,
		},
		Permissions: sliceHelper.Unique(auths),
	}, nil
}

func (s *SystemService) genToken(ctx context.Context, uid int32, uname string) string {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uid":   cast.ToString(uid),
			"uname": uname,
		})
	signedString, _ := claims.SignedString([]byte(s.cfg.Auth.ApiKey))
	token := encryption.EncodeMD5(signedString)
	// 生成redis
	data.RedisCli.Set(ctx, token, signedString, time.Second*time.Duration(s.cfg.Auth.ApiKeyExpire))

	return token
}

func (s *SystemService) GetMenuRoute(ctx context.Context, req *emptypb.Empty) (*pb.GetMenuRouteReply, error) {
	m, err := s.adminRepo.FindUserByID(ctx, ctxdata.MgrID(ctx))
	if gormHelper.IsErrRecordNotFound(err) {
		return nil, errors.New(401, "UNAUTHORIZED_INFO_MISSING", "授权已过期或授权异常,请重新授权！")
	}
	var permissions []string
	for _, v := range m.AuthRoles {
		if v.Role != nil {
			permissions = append(permissions, v.Role.MenuIds...)
		}
	}
	var list []*model.LaSystemAuthMenu
	if sliceHelper.InArr("*", permissions) {
		list, err = s.menuRepo.MenuListAll(ctx)
	} else {
		var ids []int32
		for _, permission := range permissions {
			ids = append(ids, cast.ToInt32(permission))
		}
		list, err = s.menuRepo.MenuListByIds(ctx, ids)
	}
	if err != nil {
		return nil, err
	}
	menu := s.menuRepo.FormatMenu(list, 0)
	var items []*pb.GetMenuRouteReply_MenuRoute
	for _, authMenu := range menu {
		items = append(items, modelMenuToApi(authMenu))
	}
	return &pb.GetMenuRouteReply{
		Items: items,
	}, nil
}

func (s *SystemService) GetMenuAllList(ctx context.Context, req *emptypb.Empty) (*pb.GetMenuRouteReply, error) {
	var list []*model.LaSystemAuthMenu
	list, err := s.menuRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	menu := s.menuRepo.FormatMenu(list, 0)
	var items []*pb.GetMenuRouteReply_MenuRoute
	for _, authMenu := range menu {
		items = append(items, modelMenuToApi(authMenu))
	}
	return &pb.GetMenuRouteReply{
		Items: items,
	}, nil
}

func (s *SystemService) SaveMenu(ctx context.Context, req *pb.SaveMenuReq) (*emptypb.Empty, error) {
	err := s.menuRepo.SaveMenu(ctx, &model.LaSystemAuthMenu{
		ID:        req.Id,
		Pid:       req.Pid,
		MenuType:  req.MenuType,
		MenuName:  req.MenuName,
		MenuIcon:  req.MenuIcon,
		MenuSort:  req.MenuSort,
		Perms:     req.Perms,
		Paths:     req.Paths,
		Component: req.Component,
		Selected:  req.Selected,
		Params:    req.Params,
		IsCache:   req.IsCache,
		IsShow:    req.IsShow,
		IsDisable: req.IsDisable,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SystemService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*emptypb.Empty, error) {
	count, err := s.menuRepo.CountChildrenMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewBizErr("该菜单有子菜单不可删除")
	}
	err = s.menuRepo.DeleteMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SystemService) RoleList(ctx context.Context, req *pb.RoleListReq) (*pb.RoleListReply, error) {
	m, total, err := s.roleRepo.ListPage(ctx, &domain.RoleListReq{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		Keywords: req.Keywords,
	})

	if err != nil {
		return nil, err
	}
	var roleIds []int32
	for _, v := range m {
		roleIds = append(roleIds, v.ID)
	}
	mMap, err := s.roleRepo.MemberCountMap(ctx, roleIds)
	if err != nil {
		return nil, err
	}

	var items []*pb.RoleListReply_Lists
	for _, role := range m {
		items = append(items, &pb.RoleListReply_Lists{
			Id:         role.ID,
			Name:       role.Name,
			Remark:     role.Remark,
			Menus:      role.MenuIds,
			Member:     mMap[role.ID],
			Sort:       role.Sort,
			IsDisable:  role.IsDisable,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
		})
	}
	return &pb.RoleListReply{
		Total: total,
		Items: items,
	}, nil
}

func (s *SystemService) SaveRole(ctx context.Context, req *pb.SaveRoleReq) (*emptypb.Empty, error) {
	err := s.roleRepo.Save(ctx, &model.LaSystemAuthRole{
		ID:        req.Id,
		Name:      req.Name,
		Remark:    req.Remark,
		IsDisable: req.IsDisable,
		Sort:      req.Sort,
		MenuIds:   strings.Split(req.MenuIds, ","),
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SystemService) ChangeRoleStatus(ctx context.Context, req *pb.ChangeRoleStatusReq) (*emptypb.Empty, error) {
	if err := s.roleRepo.Disable(ctx, req.Id, req.Status); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SystemService) RoleDetail(ctx context.Context, req *pb.RoleDetailReq) (*pb.RoleDetailReply, error) {
	role, err := s.roleRepo.Detail(ctx, req.Id)
	if gormHelper.IsErrRecordNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pb.RoleDetailReply{
		Id:         role.ID,
		Name:       role.Name,
		Remark:     role.Remark,
		Menus:      role.MenuIds,
		Sort:       role.Sort,
		IsDisable:  role.IsDisable,
		CreateTime: role.CreateTime,
		UpdateTime: role.UpdateTime,
	}, nil
}

func (s *SystemService) RoleAll(ctx context.Context, empty *emptypb.Empty) (*pb.RoleAllReply, error) {
	m, err := s.roleRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var items []*pb.RoleAllReply_Lists
	for _, role := range m {
		items = append(items, &pb.RoleAllReply_Lists{
			Id:         role.ID,
			Name:       role.Name,
			Remark:     role.Remark,
			Menus:      role.MenuIds,
			Sort:       role.Sort,
			IsDisable:  role.IsDisable,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
		})
	}
	return &pb.RoleAllReply{
		Items: items,
	}, nil
}
func (s *SystemService) AdminList(ctx context.Context, req *pb.AdminListReq) (*pb.AdminListReply, error) {
	m, total, err := s.adminRepo.ListPage(ctx, &domain.AdminListReq{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		Username: req.Username,
		Nickname: req.Nickname,
		Role:     req.Role,
	})
	if err != nil {
		return nil, err
	}
	var items []*pb.AdminListReply_AdminItems
	for _, user := range m {
		var role []string
		for _, v := range user.AuthRoles {
			if v.Role != nil {
				role = append(role, v.Role.Name)
			}
		}
		items = append(items, &pb.AdminListReply_AdminItems{
			Id:            user.ID,
			Username:      user.Username,
			Nickname:      user.Nickname,
			Avatar:        user.Avatar,
			Role:          strings.Join(role, ","),
			IsDisable:     user.IsDisable,
			LastLoginIp:   user.LastLoginIP,
			LastLoginTime: user.LastLoginTime,
			CreateTime:    user.CreateTime,
			UpdateTime:    user.UpdateTime,
		})
	}

	return &pb.AdminListReply{
		Items: items,
		Total: total,
	}, nil
}

func (s *SystemService) AdminDetail(ctx context.Context, req *pb.AdminDetailReq) (*pb.AdminDetailReply, error) {
	m, err := s.adminRepo.FindUserByID(ctx, ctxdata.MgrID(ctx))
	var role []string
	for _, v := range m.AuthRoles {
		role = append(role, cast.ToString(v.RoleID))
	}

	return &pb.AdminDetailReply{
		Id:       m.ID,
		Nickname: m.Nickname,
		Avatar:   m.Avatar,
		Username: m.Username,
		RoleIds:  role,
	}, err
}
func (s *SystemService) AdminSave(ctx context.Context, req *pb.AdminSaveReq) (*pb.AdminSaveReply, error) {
	m := &model.LaSystemAuthAdmin{
		ID:       req.Id,
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
		Username: req.Username,
	}

	if err := s.adminRepo.Save(ctx, m, req.RoleIds); err != nil {
		return nil, err
	}

	return &pb.AdminSaveReply{Id: m.ID}, nil
}
func (s *SystemService) AdminDisable(ctx context.Context, req *pb.AdminDisableReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.adminRepo.ChangeStatus(ctx, req.Id, req.Status)
}
func (s *SystemService) UpdateInfo(ctx context.Context, req *pb.UpdateInfoReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.adminRepo.UpdateInfo(ctx, ctxdata.MgrID(ctx), &domain.UpdateInfoReq{
		Avatar:          req.Avatar,
		Username:        req.Username,
		Nickname:        req.Nickname,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
}
