package service

import (
	"austin-v2/api/authorization/v1"
	"context"
)

func (s *AuthorizationService) GetRoleMenuTree(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {

	menu, err := s.authorizationUsecase.GetRoleMenuTree(ctx, req.Role)
	if err != nil {
		return nil, err
	}
	list := []*v1.MenuInfo{}
	for k := range menu {
		children := findChildrenMenu(menu[k])
		res := &v1.MenuInfo{
			Id:        menu[k].Id,
			ParentId:  menu[k].ParentId,
			Path:      menu[k].Path,
			Name:      menu[k].Name,
			Hidden:    menu[k].Hidden,
			Component: menu[k].Component,
			Sort:      menu[k].Sort,
			Title:     menu[k].Title,
			Icon:      menu[k].Icon,
			CreatedAt: menu[k].CreatedAt,
			UpdatedAt: menu[k].UpdatedAt,
			Children:  children,
		}
		list = append(list, res)
	}

	return &v1.GetMenuTreeReply{
		List: list,
	}, nil

}

func (s *AuthorizationService) GetRoleMenu(ctx context.Context, req *v1.GetRoleMenuRequest) (*v1.GetMenuTreeReply, error) {
	menu, err := s.authorizationUsecase.GetRoleMenu(ctx, req.Role)
	if err != nil {
		return nil, err
	}
	list := []*v1.MenuInfo{}
	for k, v := range menu {
		btns := []*v1.MenuBtn{}
		for _, btn := range v.MenuBtns {
			btns = append(btns, &v1.MenuBtn{
				Id:          btn.Id,
				MenuId:      btn.MenuId,
				Name:        btn.Name,
				Description: btn.Description,
				CreatedAt:   btn.CreatedAt,
				UpdatedAt:   btn.UpdatedAt,
			})
		}

		res := &v1.MenuInfo{
			Id:        menu[k].Id,
			ParentId:  menu[k].ParentId,
			Path:      menu[k].Path,
			Name:      menu[k].Name,
			Hidden:    menu[k].Hidden,
			Component: menu[k].Component,
			Sort:      menu[k].Sort,
			Title:     menu[k].Title,
			Icon:      menu[k].Icon,
			CreatedAt: menu[k].CreatedAt,
			UpdatedAt: menu[k].UpdatedAt,
			MenuBtns:  btns,
		}
		list = append(list, res)
	}

	return &v1.GetMenuTreeReply{
		List: list,
	}, nil
}
func (s *AuthorizationService) SetRoleMenu(ctx context.Context, req *v1.SetRoleMenuRequest) (*v1.CheckReply, error) {
	err := s.authorizationUsecase.SaveRoleMenu(ctx, req.RoleId, req.MenuIds)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: true,
	}, nil
}

func (s *AuthorizationService) GetRoleMenuBtn(ctx context.Context, req *v1.GetRoleMenuBtnRequest) (*v1.GetRoleMenuBtnReply, error) {
	list, err := s.authorizationUsecase.GetRoleMenuBtn(ctx, req.RoleId, req.RoleName, req.MenuId)
	if err != nil {
		return nil, err
	}
	resList := []*v1.MenuBtn{}
	for _, v := range list {
		resList = append(resList, &v1.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return &v1.GetRoleMenuBtnReply{
		List: resList,
	}, nil

}

func (s *AuthorizationService) SetRoleMenuBtn(ctx context.Context, req *v1.SetRoleMenuBtnRequest) (*v1.CheckReply, error) {
	err := s.authorizationUsecase.SetRoleMenuBtn(ctx, req.RoleId, req.MenuId, req.MenuBtnIds)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: true,
	}, nil
}
