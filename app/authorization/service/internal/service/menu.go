package service

import (
	"context"
	"github.com/ZQCard/kratos-base-project/api/authorization/v1"
	"github.com/ZQCard/kratos-base-project/app/authorization/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthorizationService) GetMenuAll(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {

	menu, err := s.authorizationUsecase.GetMenuAll(ctx)
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
			ParentIds: menu[k].ParentIds,
			Path:      menu[k].Path,
			Name:      menu[k].Name,
			Hidden:    menu[k].Hidden,
			Component: menu[k].Component,
			Sort:      menu[k].Sort,
			Title:     menu[k].Title,
			Icon:      menu[k].Icon,
			CreatedAt: menu[k].CreatedAt,
			UpdatedAt: menu[k].UpdatedAt,
			Children:  nil,
			MenuBtns:  btns,
		}
		list = append(list, res)
	}

	return &v1.GetMenuTreeReply{
		List: list,
	}, nil
}

func (s *AuthorizationService) GetMenuTree(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {

	menu, err := s.authorizationUsecase.GetMenuTree(ctx)
	if err != nil {
		return nil, err
	}
	list := []*v1.MenuInfo{}
	for k, v := range menu {
		children := findChildrenMenu(menu[k])
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
			ParentIds: menu[k].ParentIds,
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
			MenuBtns:  btns,
		}
		list = append(list, res)
	}

	return &v1.GetMenuTreeReply{
		List: list,
	}, nil

}

func findChildrenMenu(menu *biz.Menu) []*v1.MenuInfo {
	children := []*v1.MenuInfo{}
	if len(menu.Children) != 0 {
		for k := range menu.Children {
			btns := []*v1.MenuBtn{}
			for _, btn := range menu.Children[k].MenuBtns {
				btns = append(btns, &v1.MenuBtn{
					Id:          btn.Id,
					MenuId:      btn.MenuId,
					Name:        btn.Name,
					Description: btn.Description,
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
				Children:  findChildrenMenu(&menu.Children[k]),
			})
		}
	}
	return children
}

func (s *AuthorizationService) CreateMenu(ctx context.Context, req *v1.CreateMenuRequest) (*v1.MenuInfo, error) {
	btns := []biz.MenuBtn{}
	for _, v := range req.MenuBtns {
		btns = append(btns, biz.MenuBtn{
			Name:        v.Name,
			Description: v.Description,
		})
	}
	bc := &biz.Menu{
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
	}

	menu, err := s.authorizationUsecase.CreateMenu(ctx, bc)
	if err != nil {
		return nil, err
	}
	btns2 := []*v1.MenuBtn{}
	for _, v := range menu.MenuBtns {
		btns2 = append(btns2, &v1.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return &v1.MenuInfo{
		Id:        menu.Id,
		ParentId:  menu.ParentId,
		ParentIds: menu.ParentIds,
		Path:      menu.Path,
		Name:      menu.Name,
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
		MenuBtns:  btns2,
	}, nil
}

func (s *AuthorizationService) UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest) (*v1.MenuInfo, error) {
	btns := []biz.MenuBtn{}
	for _, v := range req.MenuBtns {
		btns = append(btns, biz.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	bc := &biz.Menu{
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
	}

	menu, err := s.authorizationUsecase.UpdateMenu(ctx, bc)
	if err != nil {
		return nil, err
	}
	btns2 := []*v1.MenuBtn{}
	for _, v := range menu.MenuBtns {
		btns2 = append(btns2, &v1.MenuBtn{
			Id:          v.Id,
			MenuId:      v.MenuId,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return &v1.MenuInfo{
		Id:        menu.Id,
		ParentId:  menu.ParentId,
		ParentIds: menu.ParentIds,
		Path:      menu.Path,
		Name:      menu.Name,
		Hidden:    menu.Hidden,
		Component: menu.Component,
		Sort:      menu.Sort,
		Title:     menu.Title,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
		MenuBtns:  btns2,
	}, nil
}

func (s *AuthorizationService) DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest) (*v1.CheckReply, error) {
	err := s.authorizationUsecase.DeleteMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.CheckReply{
		IsSuccess: true,
	}, nil
}
