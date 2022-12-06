import request from '@/utils/request'

export function listRoleAll(params) {
  return request({
    url: '/authorization/v1/roleAll',
    method: 'get',
    params
  })
}

export function listRole(params) {
  return request({
    url: '/authorization/v1/role',
    method: 'get',
    params
  })
}

export function createRole(data) {
  return request({
    url: '/authorization/v1/role',
    method: 'post',
    data
  })
}

export function updateRole(data) {
  return request({
    url: '/authorization/v1/role',
    method: 'put',
    data
  })
}

export function deleteRole(params) {
  return request({
    url: '/authorization/v1/role',
    method: 'delete',
    params
  })
}

export function saveRoleMenu(data) {
  return request({
    url: '/authorization/v1/roleMenu',
    method: 'post',
    data
  })
}

export function getRolePolicies(params) {
  return request({
    url: '/authorization/v1/getPolicies',
    method: 'get',
    params
  })
}

export function saveRolePolicies(data) {
  return request({
    url: '/authorization/v1/updatePolicies',
    method: 'post',
    data
  })
}

export function getRoleMenuBtn(params) {
  return request({
    url: '/authorization/v1/roleMenuBtn',
    method: 'get',
    params
  })
}

export function setRoleMenuBtn(data) {
  return request({
    url: '/authorization/v1/roleMenuBtn',
    method: 'post',
    data
  })
}

export function saveAdministratorRole(data) {
  return request({
    url: '/authorization/v1/setRolesForUser',
    method: 'post',
    data
  })
}

export function getAdministratorRole(params) {
  return request({
    url: '/authorization/v1/getRolesForUser',
    method: 'get',
    params
  })
}

