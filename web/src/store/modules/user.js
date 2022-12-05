import { login, logout, getInfo, loginSuccess } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { getRoleMenuBtn } from '@/api/auth/role'
import router, { resetRouter } from '@/router'

const state = {
  token: getToken(),
  name: '',
  avatar: '',
  introduction: '',
  roles: [],
  buttons: [],
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_BUTTONS: (state, buttons) => {
    state.buttons = buttons
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        const { data } = response
        commit('SET_TOKEN', data.token)
        setToken(data.token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        const { data } = response

        if (!data) {
          reject('Verification failed, please Login again.')
        }

        const { nickname, avatar } = data
        const roles = [data.role]
        // roles must be a non-empty array
        if (!roles || roles.length <= 0) {
          reject('getInfo: roles must be a non-null array!')
        }
        // 获取数据成功请求后台保存登录数据
        loginSuccess()

        // 获取角色按钮权限id集合
        getRoleMenuBtn({ role_name: roles[0] }).then(response => {
          let data = response
          if (response.code !== 0) {
            throw new Error('按钮权限数据加载异常')
          } else {
            const menuBtnIds = response.data.menuBtnIds
            // 设置按钮
            commit('SET_BUTTONS', menuBtnIds)
          }
        }).catch(error => {
          console.log(error)
        })

        commit('SET_ROLES', roles)
        commit('SET_NAME', nickname)
        commit('SET_AVATAR', avatar)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        removeToken()
        resetRouter()

        // reset visited views and cached views
        // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
        dispatch('tagsView/delAllViews', null, { root: true })

        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },

  // dynamically modify permissions
  async changeRoles({ commit, dispatch }, role) {
    const token = role + '-token'

    commit('SET_TOKEN', token)
    setToken(token)

    const { roles } = await dispatch('getInfo')

    resetRouter()

    // generate accessible routes map based on roles
    const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })
    // dynamically add accessible routes
    router.addRoutes(accessRoutes)

    // reset visited views and cached views
    dispatch('tagsView/delAllViews', null, { root: true })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
