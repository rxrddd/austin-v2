package initDB

import (
	"context"
	"encoding/json"
	"github.com/ZQCard/kratos-base-project/app/jobs/service/conf"
	"github.com/ZQCard/kratos-base-project/pkg/utils/redisHelper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const TypeInitDB = "initDB"

func NewInitDbTask(data *conf.Data) (*asynq.Task, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeInitDB, payload), nil
}

func ListenClearSignal(conf *conf.Data) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
		},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeInitDB, initDb)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func initDb(ctx context.Context, t *asynq.Task) error {
	var config conf.Data
	if err := json.Unmarshal(t.Payload(), &config); err != nil {
		return err
	}
	// 清除以 "kratos-base-project" 开头的缓存 redis缓存
	client := redis.NewClient(&redis.Options{
		Addr:         config.Redis.Addr,
		Password:     config.Redis.Password,
		ReadTimeout:  config.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: config.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	redisHelper.BatchDeleteRedisCache(client, "kratos-base-project")

	recordAdministrator := []string{
		"INSERT INTO `sys_administrator` VALUES (18, 'admin', 'b9819e53ed8ea2b2b422ff1d2f1317ca', 'e701bf4e804804773099b4b20130d418', '18158445331', '卡牌', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '超级管理员', '2022-12-07 17:05:08', '127.0.0.1:62201', '2022-08-17 16:15:17', '2022-11-22 17:41:38', '')",
		"INSERT INTO `sys_administrator` VALUES (27, 'test', '8ab138656a71b1e001aa12cc7298f901', 'cb1461e3e59ec7a2237bf5f5fa105ab5', '18158445332', '测试', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '测试管理员', '2022-12-07 17:05:50', '127.0.0.1:59530', '2022-12-07 11:28:48', '2022-12-07 11:28:48', '')",
		"INSERT INTO `sys_administrator` VALUES (28, 'guest', 'a5ac55c657800544fa68377d4bb64505', 'dabba3032e7d0c51c48e7ca794e589b8', '18158445333', '游客', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '游客', '2022-12-07 17:06:15', '127.0.0.1:59572', '2022-12-07 11:29:26', '2022-12-07 11:29:26', '')",
	}
	administratorDB, err := gorm.Open(mysql.Open(config.Administrator.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 清空数据表
	administratorDB.Exec("TRUNCATE sys_administrator")
	// 插入数据
	for _, sql := range recordAdministrator {
		if sql == "" {
			continue
		}
		administratorDB.Exec(sql)
	}

	recordAuthorization := []string{
		"INSERT INTO `authorization_menu_btns` VALUES (20, 2, '新增', '新增管理员', 'createAdministrator', '2022-12-06 14:12:38.442', '2022-12-06 14:12:38.442')",
		"INSERT INTO `authorization_menu_btns` VALUES (21, 2, '编辑', '编辑管理员', 'updateAdministrator', '2022-12-06 14:12:38.444', '2022-12-06 14:12:38.444')",
		"INSERT INTO `authorization_menu_btns` VALUES (22, 2, '禁用/解禁', '禁用/解禁管理员', 'deleteRecoverAdministrator', '2022-12-06 14:12:38.446', '2022-12-06 14:12:38.446')",
		"INSERT INTO `authorization_menu_btns` VALUES (23, 2, '删除/恢复', '删除/恢复管理员', 'forbidAApproveAdministrator', '2022-12-06 14:12:38.448', '2022-12-06 14:12:38.448')",
		"INSERT INTO `authorization_menu_btns` VALUES (24, 3, '新增', '新增菜单', 'createMenu', '2022-12-06 14:18:02.694', '2022-12-06 14:18:02.694')",
		"INSERT INTO `authorization_menu_btns` VALUES (25, 3, '编辑', '编辑菜单', 'updateMenu', '2022-12-06 14:18:02.695', '2022-12-06 14:18:02.695')",
		"INSERT INTO `authorization_menu_btns` VALUES (26, 3, '删除', '删除菜单', 'deleteMenu', '2022-12-06 14:18:02.697', '2022-12-06 14:18:02.697')",
		"INSERT INTO `authorization_menu_btns` VALUES (27, 4, '新增', '新增角色', 'createRole', '2022-12-07 18:23:24.319', '2022-12-07 18:23:24.319')",
		"INSERT INTO `authorization_menu_btns` VALUES (28, 4, '编辑', '编辑角色', 'updateRole', '2022-12-07 18:23:24.321', '2022-12-07 18:23:24.321')",
		"INSERT INTO `authorization_menu_btns` VALUES (29, 4, '删除', '删除角色', 'deleteRole', '2022-12-07 18:23:24.322', '2022-12-07 18:23:24.322')",
		"INSERT INTO `authorization_menu_btns` VALUES (30, 4, '设置权限', '设置角色权限', 'setRolePermission', '2022-12-07 18:23:24.324', '2022-12-07 18:23:24.324')",
		"INSERT INTO `authorization_menu_btns` VALUES (31, 4, '设置菜单权限', '设置角色菜单权限', 'setRoleMenuPermission', '2022-12-07 18:23:24.325', '2022-12-07 18:23:24.325')",
		"INSERT INTO `authorization_menu_btns` VALUES (32, 4, '设置API权限', '设置角色API权限', 'setRoleAPIPermission', '2022-12-07 18:23:24.327', '2022-12-07 18:23:24.327')",
		"INSERT INTO `authorization_menu_btns` VALUES (33, 4, '设置按钮权限', '设置角色按钮权限', 'setRoleMenuButtonPermission', '2022-12-07 18:23:24.328', '2022-12-07 18:23:24.328')",
		"INSERT INTO `authorization_menu_btns` VALUES (34, 4, '查看成员', '查看角色所有成员', 'getRoleMembers', '2022-12-07 18:23:24.330', '2022-12-07 18:23:24.330')",
		"INSERT INTO `authorization_menu_btns` VALUES (35, 4, '移除角色成员', '移除角色成员', 'removeRoleMember', '2022-12-07 18:23:24.331', '2022-12-07 18:23:24.331')",
		"INSERT INTO `authorization_menu_btns` VALUES (36, 5, '删除', '删除API', 'deleteAPI', '2022-12-06 14:23:26.972', '2022-12-06 14:23:26.972')",
		"INSERT INTO `authorization_menu_btns` VALUES (37, 5, '新增', '新增API', 'createAPI', '2022-12-06 14:23:26.970', '2022-12-06 14:23:26.970')",
		"INSERT INTO `authorization_menu_btns` VALUES (38, 5, '编辑', '编辑API', 'updateAPI', '2022-12-06 14:23:26.971', '2022-12-06 14:23:26.971')",
		"INSERT INTO `authorization_menu_btns` VALUES (39, 36, '有权限', '有权限按钮', 'hasPermissionButton', '2022-12-07 11:44:35.456', '2022-12-07 11:44:35.456')",
		"INSERT INTO `authorization_menu_btns` VALUES (40, 36, '无权限', '无权限按钮', 'noPermissionButton', '2022-12-07 11:44:35.456', '2022-12-07 11:44:35.456')",
		"INSERT INTO `authorization_menus` VALUES (1, 0, '0', '/system', 'system', 0, '#', 1, '系统管理', 'el-icon-s-tools', '', '2022-11-15 14:56:16.000', '2022-11-15 14:56:18.000')",
		"INSERT INTO `authorization_menus` VALUES (2, 1, '1', '/system/adminstrator', 'administrator', 0, '/system/administrator/index', 1, '管理员管理', 'el-icon-user-solid', '', '2022-11-15 14:58:41.000', '2022-12-06 14:12:38.439')",
		"INSERT INTO `authorization_menus` VALUES (3, 1, '1', '/system/auth/menu', 'menu', 0, '/system/auth/menu/index', 2, '菜单管理', 'el-icon-notebook-2', '', '2022-12-01 15:48:57.000', '2022-12-06 14:18:02.691')",
		"INSERT INTO `authorization_menus` VALUES (4, 1, '1', '/system/auth/role', 'role', 0, '/system/auth/role/index', 3, '角色管理', 'el-icon-user', '', '2022-11-24 15:33:30.000', '2022-12-07 18:23:24.315')",
		"INSERT INTO `authorization_menus` VALUES (5, 1, '1', '/system/auth/api', 'api', 0, '/system/auth/api/index', 4, 'API管理', 'el-icon-setting', '', '2022-11-23 14:42:04.000', '2022-12-06 14:23:26.966')",
		"INSERT INTO `authorization_menus` VALUES (35, 0, '0', '/example', 'example', 0, '#', 1, '测试菜单', 'el-icon-s-flag', '', '2022-12-07 10:36:58.676', '2022-12-07 11:43:33.458')",
		"INSERT INTO `authorization_menus` VALUES (36, 35, '35', '/example/permission', 'permission_test', 0, '/example/index', 1, '按钮权限测试', 'el-icon-s-tools', '', '2022-12-07 11:44:35.452', '2022-12-07 11:44:35.452')",
		"INSERT INTO `authorization_role_menu` VALUES (74, 1, 1)",
		"INSERT INTO `authorization_role_menu` VALUES (75, 1, 2)",
		"INSERT INTO `authorization_role_menu` VALUES (76, 1, 3)",
		"INSERT INTO `authorization_role_menu` VALUES (77, 1, 4)",
		"INSERT INTO `authorization_role_menu` VALUES (78, 1, 5)",
		"INSERT INTO `authorization_role_menu` VALUES (79, 1, 35)",
		"INSERT INTO `authorization_role_menu` VALUES (80, 1, 36)",
		"INSERT INTO `authorization_role_menu` VALUES (81, 39, 35)",
		"INSERT INTO `authorization_role_menu` VALUES (82, 39, 36)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (2, 1, 2, 20)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (3, 1, 2, 21)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (4, 1, 2, 22)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (5, 1, 2, 23)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (6, 1, 3, 24)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (7, 1, 3, 25)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (8, 1, 3, 26)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (9, 1, 4, 27)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (10, 1, 4, 28)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (11, 1, 4, 29)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (12, 1, 4, 30)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (13, 1, 4, 31)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (14, 1, 4, 32)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (15, 1, 4, 33)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (16, 1, 5, 34)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (17, 1, 5, 35)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (18, 1, 5, 36)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (21, 1, 36, 39)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (22, 1, 36, 40)",
		"INSERT INTO `authorization_role_menu_btn` VALUES (23, 39, 36, 39)",
		"INSERT INTO `authorization_roles` VALUES (1, 0, '0', '超级管理员', '2022-09-07 01:12:43', '2022-11-25 14:48:48')",
		"INSERT INTO `authorization_roles` VALUES (39, 0, '0', '测试管理员', '2022-12-06 11:57:30', '2022-12-06 14:24:40')",
		"INSERT INTO `authorization_roles` VALUES (40, 0, '0', '游客', '2022-12-07 11:25:42', '2022-12-07 11:25:42')",
		"INSERT INTO `casbin_rule` VALUES (367, 'g', 'admin', '测试管理员', '', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (366, 'g', 'admin', '超级管理员', '', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (278, 'g', 'guest', '游客', '', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (277, 'g', 'test', '测试管理员', '', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (357, 'p', '测试管理员', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (355, 'p', '测试管理员', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (359, 'p', '测试管理员', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (356, 'p', '测试管理员', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (354, 'p', '测试管理员', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (358, 'p', '测试管理员', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (363, 'p', '游客', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (361, 'p', '游客', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (365, 'p', '游客', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (362, 'p', '游客', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (360, 'p', '游客', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (364, 'p', '游客', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (385, 'p', '超级管理员', '/api.admin.v1.Admin/approveAdministrator', 'PATCH', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (378, 'p', '超级管理员', '/api.admin.v1.Admin/CreateAdministrator', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (368, 'p', '超级管理员', '/api.admin.v1.Admin/CreateApi', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (389, 'p', '超级管理员', '/api.admin.v1.Admin/CreateMenu', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (398, 'p', '超级管理员', '/api.admin.v1.Admin/CreateRole', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (382, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteAdministrator', 'DELETE', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (371, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteApi', 'DELETE', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (387, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteMenu', 'DELETE', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (391, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteRole', 'DELETE', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (400, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteRoleForUser', 'DELETE', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (384, 'p', '超级管理员', '/api.admin.v1.Admin/forbidAdministrator', 'PATCH', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (381, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministrator', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (377, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (380, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministratorList', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (393, 'p', '超级管理员', '/api.admin.v1.Admin/GetApiAll', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (370, 'p', '超级管理员', '/api.admin.v1.Admin/GetApiList', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (386, 'p', '超级管理员', '/api.admin.v1.Admin/GetMenuTree', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (394, 'p', '超级管理员', '/api.admin.v1.Admin/GetPolicies', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (392, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleList', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (374, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (375, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (399, 'p', '超级管理员', '/api.admin.v1.Admin/GetUsersForRole', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (373, 'p', '超级管理员', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (372, 'p', '超级管理员', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (376, 'p', '超级管理员', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (383, 'p', '超级管理员', '/api.admin.v1.Admin/RecoverAdministrator', 'PATCH', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (395, 'p', '超级管理员', '/api.admin.v1.Admin/SetRoleMenu', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (397, 'p', '超级管理员', '/api.admin.v1.Admin/SetRoleMenuBtn', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (379, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateAdministrator', 'PUT', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (369, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateApi', 'PUT', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (388, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateMenu', 'PUT', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (396, 'p', '超级管理员', '/api.admin.v1.Admin/UpdatePolicies', 'POST', '', '', '', '', '')",
		"INSERT INTO `casbin_rule` VALUES (390, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateRole', 'PUT', '', '', '', '', '')",
	}
	authorizationDB, err := gorm.Open(mysql.Open(config.Authorization.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 清空数据表
	authorizationDB.Exec("TRUNCATE authorization_api")
	authorizationDB.Exec("TRUNCATE authorization_menu_btns")
	authorizationDB.Exec("TRUNCATE authorization_menus")
	authorizationDB.Exec("TRUNCATE authorization_role_menu")
	authorizationDB.Exec("TRUNCATE authorization_role_menu_btn")
	authorizationDB.Exec("TRUNCATE authorization_roles")
	authorizationDB.Exec("TRUNCATE casbin_rule")
	// 插入数据
	for _, sql := range recordAuthorization {
		if sql == "" {
			continue
		}
		authorizationDB.Exec(sql)
	}
	return nil
}
