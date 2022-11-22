import request from '@/utils/request'

export function listAdministrator(params) {
  return request({
    url: '/admin/v1/getAdministrators',
    method: 'get',
    params
  })
}

export function createAdministrator(data) {
  return request({
    url: '/admin/v1/administrator',
    method: 'post',
    data
  })
}

export function updateAdministrator(data) {
  return request({
    url: '/admin/v1/administrator',
    method: 'put',
    data
  })
}

export function deleteAdministrator(params) {
  return request({
    url: '/admin/v1/administrator',
    method: 'delete',
    params
  })
}

export function recoverAdministrator(data) {
  return request({
    url: '/admin/v1/administrator',
    method: 'patch',
    data
  })
}

export function forbidAdministrator(data) {
  return request({
    url: '/admin/v1/administrator/forbid',
    method: 'patch',
    data
  })
}

export function approveAdministrator(data) {
  return request({
    url: '/admin/v1/administrator/approve',
    method: 'patch',
    data
  })
}
