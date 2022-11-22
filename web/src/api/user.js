import request from '@/utils/request'
export function login(data) {
  return request({
    url: '/admin/v1/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/admin/v1/getAdministratorInfo',
    method: 'get'
  })
}

export function loginSuccess() {
  return request({
    url: '/admin/v1/loginSuccess',
    method: 'get'
  })
}

export function logout(data) {
  data = {}
  return request({
    url: '/admin/v1/logout',
    method: 'post',
    data
  })
}
