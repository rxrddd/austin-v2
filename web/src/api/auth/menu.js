import request from '@/utils/request'

export function getRoleMenu(params) {
  return request({
    url: '/authorization/v1/roleMenuTree',
    method: 'get',
    params
  })
}

export function getBaseMenuTree(params) {
  return request({
    url: '/authorization/v1/menuTree',
    method: 'get',
    params
  })
}


export function createMenu(data) {
  return request({
    url: '/authorization/v1/menu',
    method: 'post',
    data
  })
}

export function updateMenu(data) {
  return request({
    url: '/authorization/v1/menu',
    method: 'put',
    data
  })
}

export function deleteMenu(params) {
  return request({
    url: '/authorization/v1/menu',
    method: 'delete',
    params
  })
}