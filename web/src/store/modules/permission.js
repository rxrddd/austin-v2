import { asyncRoutes, constantRoutes } from '@/router'
import { getRoleMenu } from '@/api/auth/menu'
import Layout from '@/layout'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  }
}

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise(resolve => {
      const loadMenuData = []
      // 先查询后台并返回左侧菜单数据并把数据添加到路由
      getRoleMenu({ role: roles[0] }).then(response => {
        let data = response
        if (response.code !== 0) {
          throw new Error('菜单数据加载异常')
        } else {
          data = response.data.list
          Object.assign(loadMenuData, data)
          const tempAsyncRoutes = Object.assign([], asyncRoutes)
          generateMenu(tempAsyncRoutes, loadMenuData)
          let accessedRoutes
          if (roles.includes('超级管理员')) {
            accessedRoutes = tempAsyncRoutes || []
          } else {
            accessedRoutes = filterAsyncRoutes(tempAsyncRoutes, roles)
          }
          // 设置路由
          commit('SET_ROUTES', accessedRoutes)
          resolve(accessedRoutes)
        }
      }).catch(error => {
        console.log(error)
      })

    })
  }
}

/**
 * 后台查询的菜单数据拼装成路由格式的数据
 * @param routes (resolve: any) => require([`@${view}`], resolve)
 */
export function generateMenu(routes, data) {
  data.forEach(item => {
    // 设置meta
    const menu = {
      id: item.id,
      label: item.title,
      name: item.name,
      path: item.path === '#' ? item.id + '_key' : item.path,
      component: item.component === '#' ? Layout : (resolve) => require([`@/views${item.component}`], resolve),
      hidden: item.hidden !== '0',
      meta: { 'icon': item.icon, 'title': item.title, 'keepAlive': true },
      children: []
    }
    if (item.children) {
      generateMenu(menu.children, item.children)
    }
    routes.push(menu)
  })
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
