import request from '@/utils/request'

export function listApiAll(params) {
  return request({
    url: '/authorization/v1/apiAll',
    method: 'get',
    params
  })
}

export function listApi(params) {
  return request({
    url: '/authorization/v1/api',
    method: 'get',
    params
  })
}

export function createApi(data) {
  return request({
    url: '/authorization/v1/api',
    method: 'post',
    data
  })
}

export function updateApi(data) {
  return request({
    url: '/authorization/v1/api',
    method: 'put',
    data
  })
}

export function deleteApi(params) {
  return request({
    url: '/authorization/v1/api',
    method: 'delete',
    params
  })
}
