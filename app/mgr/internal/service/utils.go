package service

import (
	pb "austin-v2/api/mgr"
	"austin-v2/common/dal/model"
)

func modelMenuToApi(m *model.LaSystemAuthMenu) *pb.GetMenuRouteReply_MenuRoute {
	items := make([]*pb.GetMenuRouteReply_MenuRoute, 0)
	for _, child := range m.Children {
		items = append(items, modelMenuToApi(child))
	}
	return &pb.GetMenuRouteReply_MenuRoute{
		Children:   items,
		Component:  m.Component,
		CreateTime: m.CreateTime,
		Id:         m.ID,
		IsCache:    m.IsCache,
		IsDisable:  m.IsDisable,
		IsShow:     m.IsShow,
		MenuIcon:   m.MenuIcon,
		MenuName:   m.MenuName,
		MenuSort:   m.MenuSort,
		MenuType:   m.MenuType,
		Params:     m.Params,
		Paths:      m.Paths,
		Perms:      m.Perms,
		Pid:        m.Pid,
		Selected:   m.Selected,
		UpdateTime: m.UpdateTime,
	}
}
