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