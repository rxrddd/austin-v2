import request from '@/utils/request'

export function getAdministratorMenu(params) {
  return request({
    url: '/authorization/v1/menuTree',
    method: 'get',
    params
  })
}
