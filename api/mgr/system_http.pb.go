// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.19.4
// source: api/mgr/system.proto

package mgr

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSystemAdminDetail = "/api.gvs_mgr.System/AdminDetail"
const OperationSystemAdminDisable = "/api.gvs_mgr.System/AdminDisable"
const OperationSystemAdminList = "/api.gvs_mgr.System/AdminList"
const OperationSystemAdminLogin = "/api.gvs_mgr.System/AdminLogin"
const OperationSystemAdminLogout = "/api.gvs_mgr.System/AdminLogout"
const OperationSystemAdminSave = "/api.gvs_mgr.System/AdminSave"
const OperationSystemChangeRoleStatus = "/api.gvs_mgr.System/ChangeRoleStatus"
const OperationSystemDeleteMenu = "/api.gvs_mgr.System/DeleteMenu"
const OperationSystemGetMenuAllList = "/api.gvs_mgr.System/GetMenuAllList"
const OperationSystemGetMenuRoute = "/api.gvs_mgr.System/GetMenuRoute"
const OperationSystemGetSelfInfo = "/api.gvs_mgr.System/GetSelfInfo"
const OperationSystemRoleAll = "/api.gvs_mgr.System/RoleAll"
const OperationSystemRoleDetail = "/api.gvs_mgr.System/RoleDetail"
const OperationSystemRoleList = "/api.gvs_mgr.System/RoleList"
const OperationSystemSaveMenu = "/api.gvs_mgr.System/SaveMenu"
const OperationSystemSaveRole = "/api.gvs_mgr.System/SaveRole"
const OperationSystemUpdateInfo = "/api.gvs_mgr.System/UpdateInfo"

type SystemHTTPServer interface {
	AdminDetail(context.Context, *AdminDetailReq) (*AdminDetailReply, error)
	AdminDisable(context.Context, *AdminDisableReq) (*emptypb.Empty, error)
	AdminList(context.Context, *AdminListReq) (*AdminListReply, error)
	AdminLogin(context.Context, *LoginReq) (*LoginResp, error)
	AdminLogout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	AdminSave(context.Context, *AdminSaveReq) (*AdminSaveReply, error)
	ChangeRoleStatus(context.Context, *ChangeRoleStatusReq) (*emptypb.Empty, error)
	DeleteMenu(context.Context, *DeleteMenuReq) (*emptypb.Empty, error)
	GetMenuAllList(context.Context, *emptypb.Empty) (*GetMenuRouteReply, error)
	GetMenuRoute(context.Context, *emptypb.Empty) (*GetMenuRouteReply, error)
	GetSelfInfo(context.Context, *emptypb.Empty) (*SelfReply, error)
	RoleAll(context.Context, *emptypb.Empty) (*RoleAllReply, error)
	RoleDetail(context.Context, *RoleDetailReq) (*RoleDetailReply, error)
	RoleList(context.Context, *RoleListReq) (*RoleListReply, error)
	SaveMenu(context.Context, *SaveMenuReq) (*emptypb.Empty, error)
	SaveRole(context.Context, *SaveRoleReq) (*emptypb.Empty, error)
	UpdateInfo(context.Context, *UpdateInfoReq) (*emptypb.Empty, error)
}

func RegisterSystemHTTPServer(s *http.Server, srv SystemHTTPServer) {
	r := s.Route("/")
	r.POST("/system/login", _System_AdminLogin0_HTTP_Handler(srv))
	r.POST("/system/logout", _System_AdminLogout0_HTTP_Handler(srv))
	r.GET("/system/admin/self", _System_GetSelfInfo0_HTTP_Handler(srv))
	r.GET("/system/menu/route", _System_GetMenuRoute0_HTTP_Handler(srv))
	r.GET("/system/menu/list", _System_GetMenuAllList0_HTTP_Handler(srv))
	r.POST("/system/menu/save", _System_SaveMenu0_HTTP_Handler(srv))
	r.POST("/system/menu/del", _System_DeleteMenu0_HTTP_Handler(srv))
	r.GET("/system/role/list", _System_RoleList0_HTTP_Handler(srv))
	r.GET("/system/role/all", _System_RoleAll0_HTTP_Handler(srv))
	r.POST("/system/role/save", _System_SaveRole0_HTTP_Handler(srv))
	r.POST("/system/role/change", _System_ChangeRoleStatus0_HTTP_Handler(srv))
	r.GET("/system/role/detail", _System_RoleDetail0_HTTP_Handler(srv))
	r.GET("/system/admin/list", _System_AdminList0_HTTP_Handler(srv))
	r.GET("/system/admin/detail", _System_AdminDetail0_HTTP_Handler(srv))
	r.POST("/system/admin/save", _System_AdminSave0_HTTP_Handler(srv))
	r.POST("/system/admin/disable", _System_AdminDisable0_HTTP_Handler(srv))
	r.POST("/system/admin/upInfo", _System_UpdateInfo0_HTTP_Handler(srv))
}

func _System_AdminLogin0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminLogin(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _System_AdminLogout0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminLogout(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_GetSelfInfo0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetSelfInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSelfInfo(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SelfReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetMenuRoute0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetMenuRoute)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenuRoute(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenuRouteReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetMenuAllList0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetMenuAllList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenuAllList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenuRouteReply)
		return ctx.Result(200, reply)
	}
}

func _System_SaveMenu0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveMenuReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemSaveMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveMenu(ctx, req.(*SaveMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_DeleteMenu0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenuReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*DeleteMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_RoleList0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemRoleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleList(ctx, req.(*RoleListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleListReply)
		return ctx.Result(200, reply)
	}
}

func _System_RoleAll0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemRoleAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleAll(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleAllReply)
		return ctx.Result(200, reply)
	}
}

func _System_SaveRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveRoleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemSaveRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveRole(ctx, req.(*SaveRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_ChangeRoleStatus0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangeRoleStatusReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemChangeRoleStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeRoleStatus(ctx, req.(*ChangeRoleStatusReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_RoleDetail0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleDetailReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemRoleDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleDetail(ctx, req.(*RoleDetailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleDetailReply)
		return ctx.Result(200, reply)
	}
}

func _System_AdminList0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminList(ctx, req.(*AdminListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminListReply)
		return ctx.Result(200, reply)
	}
}

func _System_AdminDetail0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminDetailReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminDetail(ctx, req.(*AdminDetailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminDetailReply)
		return ctx.Result(200, reply)
	}
}

func _System_AdminSave0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminSaveReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminSave(ctx, req.(*AdminSaveReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminSaveReply)
		return ctx.Result(200, reply)
	}
}

func _System_AdminDisable0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminDisableReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAdminDisable)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminDisable(ctx, req.(*AdminDisableReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _System_UpdateInfo0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemUpdateInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateInfo(ctx, req.(*UpdateInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type SystemHTTPClient interface {
	AdminDetail(ctx context.Context, req *AdminDetailReq, opts ...http.CallOption) (rsp *AdminDetailReply, err error)
	AdminDisable(ctx context.Context, req *AdminDisableReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	AdminList(ctx context.Context, req *AdminListReq, opts ...http.CallOption) (rsp *AdminListReply, err error)
	AdminLogin(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	AdminLogout(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	AdminSave(ctx context.Context, req *AdminSaveReq, opts ...http.CallOption) (rsp *AdminSaveReply, err error)
	ChangeRoleStatus(ctx context.Context, req *ChangeRoleStatusReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetMenuAllList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetMenuRouteReply, err error)
	GetMenuRoute(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetMenuRouteReply, err error)
	GetSelfInfo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SelfReply, err error)
	RoleAll(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *RoleAllReply, err error)
	RoleDetail(ctx context.Context, req *RoleDetailReq, opts ...http.CallOption) (rsp *RoleDetailReply, err error)
	RoleList(ctx context.Context, req *RoleListReq, opts ...http.CallOption) (rsp *RoleListReply, err error)
	SaveMenu(ctx context.Context, req *SaveMenuReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	SaveRole(ctx context.Context, req *SaveRoleReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateInfo(ctx context.Context, req *UpdateInfoReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type SystemHTTPClientImpl struct {
	cc *http.Client
}

func NewSystemHTTPClient(client *http.Client) SystemHTTPClient {
	return &SystemHTTPClientImpl{client}
}

func (c *SystemHTTPClientImpl) AdminDetail(ctx context.Context, in *AdminDetailReq, opts ...http.CallOption) (*AdminDetailReply, error) {
	var out AdminDetailReply
	pattern := "/system/admin/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemAdminDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AdminDisable(ctx context.Context, in *AdminDisableReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/admin/disable"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAdminDisable))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AdminList(ctx context.Context, in *AdminListReq, opts ...http.CallOption) (*AdminListReply, error) {
	var out AdminListReply
	pattern := "/system/admin/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemAdminList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AdminLogin(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/system/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAdminLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AdminLogout(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAdminLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AdminSave(ctx context.Context, in *AdminSaveReq, opts ...http.CallOption) (*AdminSaveReply, error) {
	var out AdminSaveReply
	pattern := "/system/admin/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAdminSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) ChangeRoleStatus(ctx context.Context, in *ChangeRoleStatusReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/role/change"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemChangeRoleStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/menu/del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetMenuAllList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetMenuRouteReply, error) {
	var out GetMenuRouteReply
	pattern := "/system/menu/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetMenuAllList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetMenuRoute(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetMenuRouteReply, error) {
	var out GetMenuRouteReply
	pattern := "/system/menu/route"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetMenuRoute))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetSelfInfo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SelfReply, error) {
	var out SelfReply
	pattern := "/system/admin/self"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetSelfInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) RoleAll(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*RoleAllReply, error) {
	var out RoleAllReply
	pattern := "/system/role/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemRoleAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) RoleDetail(ctx context.Context, in *RoleDetailReq, opts ...http.CallOption) (*RoleDetailReply, error) {
	var out RoleDetailReply
	pattern := "/system/role/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemRoleDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) RoleList(ctx context.Context, in *RoleListReq, opts ...http.CallOption) (*RoleListReply, error) {
	var out RoleListReply
	pattern := "/system/role/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemRoleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) SaveMenu(ctx context.Context, in *SaveMenuReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/menu/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemSaveMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) SaveRole(ctx context.Context, in *SaveRoleReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/role/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemSaveRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) UpdateInfo(ctx context.Context, in *UpdateInfoReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/admin/upInfo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemUpdateInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}