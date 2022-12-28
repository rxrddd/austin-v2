package data

import (
	"context"
	v1 "github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/emptypb"

	authorizationServiceV1 "github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type AuthorizationRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewAuthorizationRepo(data *Data, logger log.Logger) *AuthorizationRepo {
	return &AuthorizationRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/administrator")),
		sg:   &singleflight.Group{},
	}
}

//, tp *tracesdk.TracerProvider
func NewAuthorizationServiceClient(ac *conf.Auth, sr *conf.Service, r registry.Discovery) authorizationServiceV1.AuthorizationClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Authorization.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			//tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			//jwt.Client(func(token *jwt2.Token) (interface{}, error) {
			//	return []byte(ac.ServiceKey), nil
			//}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := authorizationServiceV1.NewAuthorizationClient(conn)
	return c
}

func (rp AuthorizationRepo) GetRoleList(ctx context.Context) (*v1.GetRoleListReply, error) {
	reply, err := rp.data.authorizationClient.GetRoleList(ctx, &emptypb.Empty{})

	if err != nil {
		return nil, err
	}
	roles := []*v1.RoleInfo{}
	for _, v := range reply.List {
		children := findChildrenRole(v)
		roles = append(roles, &v1.RoleInfo{
			Id:        v.Id,
			ParentId:  v.ParentId,
			Name:      v.Name,
			ParentIds: v.ParentIds,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  children,
		})
	}
	res := &v1.GetRoleListReply{
		List: roles,
	}
	return res, err
}

func findChildrenRole(role *authorizationServiceV1.RoleInfo) []*v1.RoleInfo {
	children := []*v1.RoleInfo{}
	if len(role.Children) != 0 {
		for k := range role.Children {
			children = append(children, &v1.RoleInfo{
				Id:        role.Children[k].Id,
				Name:      role.Children[k].Name,
				ParentId:  role.Children[k].ParentId,
				ParentIds: role.Children[k].ParentIds,
				CreatedAt: role.Children[k].CreatedAt,
				UpdatedAt: role.Children[k].UpdatedAt,
				Children:  findChildrenRole(role.Children[k]),
			})
		}
	}
	return children
}

func (rp AuthorizationRepo) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.RoleInfo, error) {
	reply, err := rp.data.authorizationClient.CreateRole(ctx, &authorizationServiceV1.CreateRoleRequest{
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.RoleInfo{
		Id:        reply.Id,
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
	return res, nil
}

func (rp AuthorizationRepo) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.RoleInfo, error) {
	reply, err := rp.data.authorizationClient.UpdateRole(ctx, &authorizationServiceV1.UpdateRoleRequest{
		Id:        req.Id,
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.RoleInfo{
		Id:        reply.Id,
		Name:      req.Name,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
	return res, nil
}

func (rp AuthorizationRepo) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.DeleteRole(ctx, &authorizationServiceV1.DeleteRoleRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) SetRolesForUser(ctx context.Context, req *v1.SetRolesForUserRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.SetRolesForUser(ctx, &authorizationServiceV1.SetRolesForUserRequest{
		Username: req.Username,
		Roles:    req.Roles,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetRolesForUser(ctx context.Context, req *v1.GetRolesForUserRequest) (*v1.GetRolesForUserReply, error) {
	reply, err := rp.data.authorizationClient.GetRolesForUser(ctx, &authorizationServiceV1.GetRolesForUserRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.GetRolesForUserReply{
		Roles: reply.Roles,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetUsersForRole(ctx context.Context, req *v1.GetUsersForRoleRequest) (*v1.GetUsersForRoleReply, error) {
	reply, err := rp.data.authorizationClient.GetUsersForRole(ctx, &authorizationServiceV1.GetUsersForRoleRequest{
		Role: req.Role,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.GetUsersForRoleReply{
		Users: reply.Users,
	}
	return res, nil
}

func (rp AuthorizationRepo) DeleteRoleForUser(ctx context.Context, req *v1.DeleteRoleForUserRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.DeleteRoleForUser(ctx, &authorizationServiceV1.DeleteRoleForUserRequest{
		Username: req.Username,
		Role:     req.Role,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetPolicies(ctx context.Context, req *v1.GetPoliciesRequest) (*v1.GetPoliciesReply, error) {
	reply, err := rp.data.authorizationClient.GetPolicies(ctx, &authorizationServiceV1.GetPoliciesRequest{
		Role: req.Role,
	})
	if err != nil {
		return nil, err
	}
	var policyRules []*v1.PolicyRules
	for _, v := range reply.PolicyRules {
		policyRules = append(policyRules, &v1.PolicyRules{
			Path:   v.Path,
			Method: v.Method,
		})
	}
	res := &v1.GetPoliciesReply{
		PolicyRules: policyRules,
	}
	return res, nil
}

func (rp AuthorizationRepo) UpdatePolicies(ctx context.Context, req *v1.UpdatePoliciesRequest) (*v1.CheckReply, error) {
	var policyRules []*authorizationServiceV1.PolicyRules
	for _, v := range req.PolicyRules {
		policyRules = append(policyRules, &authorizationServiceV1.PolicyRules{
			Path:   v.Path,
			Method: v.Method,
		})
	}

	reply, err := rp.data.authorizationClient.UpdatePolicies(ctx, &authorizationServiceV1.UpdatePoliciesRequest{
		Role:        req.Role,
		PolicyRules: policyRules,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetApiAll(ctx context.Context) (*v1.GetApiAllReply, error) {
	reply, err := rp.data.authorizationClient.GetApiAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var list []*v1.ApiInfo
	for _, v := range reply.List {
		list = append(list, &v1.ApiInfo{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	res := &v1.GetApiAllReply{
		List: list,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetApiList(ctx context.Context, req *v1.GetApiListRequest) (*v1.GetApiListReply, error) {
	reply, err := rp.data.authorizationClient.GetApiList(ctx, &authorizationServiceV1.GetApiListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Group:    req.Group,
		Name:     req.Name,
		Path:     req.Path,
		Method:   req.Method,
	})
	if err != nil {
		return nil, err
	}
	var list []*v1.ApiInfo
	for _, v := range reply.List {
		list = append(list, &v1.ApiInfo{
			Id:        v.Id,
			Group:     v.Group,
			Name:      v.Name,
			Path:      v.Path,
			Method:    v.Method,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	res := &v1.GetApiListReply{
		List:  list,
		Total: reply.Total,
	}
	return res, nil
}

func (rp AuthorizationRepo) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.ApiInfo, error) {
	reply, err := rp.data.authorizationClient.CreateApi(ctx, &authorizationServiceV1.CreateApiRequest{
		Group:  req.Group,
		Name:   req.Name,
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.ApiInfo{
		Id:        reply.Id,
		Group:     reply.Group,
		Name:      reply.Name,
		Path:      reply.Path,
		Method:    reply.Method,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
	return res, nil
}

func (rp AuthorizationRepo) UpdateApi(ctx context.Context, req *v1.UpdateApiRequest) (*v1.ApiInfo, error) {
	reply, err := rp.data.authorizationClient.UpdateApi(ctx, &authorizationServiceV1.UpdateApiRequest{
		Id:     req.Id,
		Group:  req.Group,
		Name:   req.Name,
		Path:   req.Path,
		Method: req.Method,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.ApiInfo{
		Id:        reply.Id,
		Group:     reply.Group,
		Name:      reply.Name,
		Path:      reply.Path,
		Method:    reply.Method,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: reply.UpdatedAt,
	}
	return res, nil
}

func (rp AuthorizationRepo) DeleteApi(ctx context.Context, req *v1.DeleteApiRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.DeleteApi(ctx, &authorizationServiceV1.DeleteApiRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetMenuAll(ctx context.Context) (*v1.GetMenuTreeReply, error) {
	reply, err := rp.data.authorizationClient.GetMenuAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var list []*v1.MenuInfo
	for _, v := range reply.List {
		var btns []*v1.MenuBtn
		for _, btn := range v.MenuBtns {
			btns = append(btns, &v1.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt,
				UpdatedAt:   btn.UpdatedAt,
			})
		}

		list = append(list, &v1.MenuInfo{
			Id:        v.Id,
			ParentId:  v.ParentId,
			ParentIds: v.ParentIds,
			Path:      v.Path,
			Name:      v.Name,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			MenuBtns:  btns,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	res := &v1.GetMenuTreeReply{
		List: list,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetMenuTree(ctx context.Context) (*v1.GetMenuTreeReply, error) {
	reply, err := rp.data.authorizationClient.GetMenuTree(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var list []*v1.MenuInfo
	menu := reply.List
	for k, v := range menu {
		children := findChildrenMenu(menu[k])

		var btns []*v1.MenuBtn
		for _, btn := range v.MenuBtns {
			btns = append(btns, &v1.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt,
				UpdatedAt:   btn.UpdatedAt,
			})
		}

		list = append(list, &v1.MenuInfo{
			Id:        v.Id,
			ParentId:  v.ParentId,
			ParentIds: v.ParentIds,
			Path:      v.Path,
			Name:      v.Name,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			MenuBtns:  btns,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  children,
		})

	}

	res := &v1.GetMenuTreeReply{
		List: list,
	}
	return res, nil
}

func findChildrenMenu(menu *authorizationServiceV1.MenuInfo) []*v1.MenuInfo {
	children := []*v1.MenuInfo{}
	if len(menu.Children) != 0 {
		for k := range menu.Children {
			var btns []*v1.MenuBtn
			for _, btn := range menu.Children[k].MenuBtns {
				btns = append(btns, &v1.MenuBtn{
					Id:          btn.Id,
					MenuId:      btn.MenuId,
					Name:        btn.Name,
					Description: btn.Description,
					Identifier:  btn.Identifier,
					CreatedAt:   btn.CreatedAt,
					UpdatedAt:   btn.UpdatedAt,
				})
			}

			children = append(children, &v1.MenuInfo{
				Id:        menu.Children[k].Id,
				Name:      menu.Children[k].Name,
				Path:      menu.Children[k].Path,
				ParentId:  menu.Children[k].ParentId,
				ParentIds: menu.Children[k].ParentIds,
				Hidden:    menu.Children[k].Hidden,
				Component: menu.Children[k].Component,
				Sort:      menu.Children[k].Sort,
				Title:     menu.Children[k].Title,
				Icon:      menu.Children[k].Icon,
				CreatedAt: menu.Children[k].CreatedAt,
				UpdatedAt: menu.Children[k].UpdatedAt,
				MenuBtns:  btns,
				Children:  findChildrenMenu(menu.Children[k]),
			})
		}
	}
	return children
}

func (rp AuthorizationRepo) CreateMenu(ctx context.Context, req *v1.CreateMenuRequest) (*v1.MenuInfo, error) {
	var btns []*authorizationServiceV1.MenuBtn
	for _, btn := range req.MenuBtns {
		btns = append(btns, &authorizationServiceV1.MenuBtn{
			Id:          btn.Id,
			MenuId:      btn.MenuId,
			Name:        btn.Name,
			Description: btn.Description,
			Identifier:  btn.Identifier,
		})
	}

	reply, err := rp.data.authorizationClient.CreateMenu(ctx, &authorizationServiceV1.CreateMenuRequest{
		Name:      req.Name,
		Path:      req.Path,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		Hidden:    req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		Title:     req.Title,
		Icon:      req.Icon,
		MenuBtns:  btns,
	})
	if err != nil {
		return nil, err
	}
	var btns2 []*v1.MenuBtn
	for _, btn := range reply.MenuBtns {
		btns2 = append(btns2, &v1.MenuBtn{
			Id:          btn.Id,
			MenuId:      btn.MenuId,
			Name:        btn.Name,
			Description: btn.Description,
			Identifier:  btn.Identifier,
			CreatedAt:   btn.CreatedAt,
			UpdatedAt:   btn.UpdatedAt,
		})
	}

	res := &v1.MenuInfo{
		Id:        reply.Id,
		Name:      req.Name,
		Path:      req.Path,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		Hidden:    req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		Title:     req.Title,
		Icon:      req.Icon,
		MenuBtns:  btns2,
	}
	return res, nil
}

func (rp AuthorizationRepo) UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest) (*v1.MenuInfo, error) {
	var btns []*authorizationServiceV1.MenuBtn
	for _, btn := range req.MenuBtns {
		btns = append(btns, &authorizationServiceV1.MenuBtn{
			Id:          btn.Id,
			MenuId:      btn.MenuId,
			Name:        btn.Name,
			Description: btn.Description,
			Identifier:  btn.Identifier,
		})
	}
	reply, err := rp.data.authorizationClient.UpdateMenu(ctx, &authorizationServiceV1.UpdateMenuRequest{
		Id:        req.Id,
		Name:      req.Name,
		Path:      req.Path,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		Hidden:    req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		Title:     req.Title,
		Icon:      req.Icon,
		MenuBtns:  btns,
	})
	if err != nil {
		return nil, err
	}
	var btns2 []*v1.MenuBtn
	for _, btn := range reply.MenuBtns {
		btns2 = append(btns2, &v1.MenuBtn{
			Id:          btn.Id,
			MenuId:      btn.MenuId,
			Name:        btn.Name,
			Description: btn.Description,
			Identifier:  btn.Identifier,
			CreatedAt:   btn.CreatedAt,
			UpdatedAt:   btn.UpdatedAt,
		})
	}

	res := &v1.MenuInfo{
		Id:        reply.Id,
		Name:      req.Name,
		Path:      req.Path,
		ParentId:  req.ParentId,
		ParentIds: req.ParentIds,
		Hidden:    req.Hidden,
		Component: req.Component,
		Sort:      req.Sort,
		Title:     req.Title,
		Icon:      req.Icon,
		MenuBtns:  btns2,
	}
	return res, nil
}

func (rp AuthorizationRepo) DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.DeleteMenu(ctx, &authorizationServiceV1.DeleteMenuRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetRoleMenuTree(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {
	reply, err := rp.data.authorizationClient.GetRoleMenuTree(ctx, &authorizationServiceV1.GetRoleMenuRequest{
		Role: req.Role,
	})
	if err != nil {
		return nil, err
	}
	var list []*v1.MenuInfo
	menu := reply.List
	for k, v := range menu {
		children := findChildrenMenu(menu[k])

		var btns []*v1.MenuBtn
		for _, btn := range v.MenuBtns {
			btns = append(btns, &v1.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt,
				UpdatedAt:   btn.UpdatedAt,
			})
		}

		list = append(list, &v1.MenuInfo{
			Id:        v.Id,
			ParentId:  v.ParentId,
			ParentIds: v.ParentIds,
			Path:      v.Path,
			Name:      v.Name,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			MenuBtns:  btns,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  children,
		})

	}

	res := &v1.GetMenuTreeReply{
		List: list,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {
	reply, err := rp.data.authorizationClient.GetRoleMenu(ctx, &authorizationServiceV1.GetRoleMenuRequest{
		Role: req.Role,
	})
	if err != nil {
		return nil, err
	}
	var list []*v1.MenuInfo
	menu := reply.List
	for _, v := range menu {
		var btns []*v1.MenuBtn
		for _, btn := range v.MenuBtns {
			btns = append(btns, &v1.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				Identifier:  btn.Identifier,
				CreatedAt:   btn.CreatedAt,
				UpdatedAt:   btn.UpdatedAt,
			})
		}

		list = append(list, &v1.MenuInfo{
			Id:        v.Id,
			ParentId:  v.ParentId,
			ParentIds: v.ParentIds,
			Path:      v.Path,
			Name:      v.Name,
			Hidden:    v.Hidden,
			Component: v.Component,
			Sort:      v.Sort,
			Title:     v.Title,
			Icon:      v.Icon,
			MenuBtns:  btns,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})

	}

	res := &v1.GetMenuTreeReply{
		List: list,
	}
	return res, nil
}

func (rp AuthorizationRepo) SetRoleMenu(ctx context.Context, req *v1.SetRoleMenuRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.SetRoleMenu(ctx, &authorizationServiceV1.SetRoleMenuRequest{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}

func (rp AuthorizationRepo) GetRoleMenuBtn(ctx context.Context, req *v1.GetRoleMenuBtnRequest) (*v1.GetRoleMenuBtnReply, error) {
	reply, err := rp.data.authorizationClient.GetRoleMenuBtn(ctx, &authorizationServiceV1.GetRoleMenuBtnRequest{
		RoleId:   req.RoleId,
		RoleName: req.RoleName,
		MenuId:   req.MenuId,
	})
	if err != nil {
		return nil, err
	}
	list := []*v1.MenuBtn{}
	for _, v := range reply.List {
		list = append(list, &v1.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	res := &v1.GetRoleMenuBtnReply{
		List: list,
	}
	return res, nil
}

func (rp AuthorizationRepo) SetRoleMenuBtn(ctx context.Context, req *v1.SetRoleMenuBtnRequest) (*v1.CheckReply, error) {
	reply, err := rp.data.authorizationClient.SetRoleMenuBtn(ctx, &authorizationServiceV1.SetRoleMenuBtnRequest{
		RoleId:     req.RoleId,
		MenuId:     req.MenuId,
		MenuBtnIds: req.MenuBtnIds,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.CheckReply{
		IsSuccess: reply.IsSuccess,
	}
	return res, nil
}
