/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : 127.0.0.1:3306
 Source Schema         : kratos-base-project-authorization

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 07/12/2022 17:13:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for authorization_api
-- ----------------------------
DROP TABLE IF EXISTS `authorization_api`;
CREATE TABLE `authorization_api`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分组',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `method` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路径',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_api
-- ----------------------------
INSERT INTO `authorization_api` VALUES (1, '基础分组', '登录', 'POST', '/api.admin.v1.Admin/Login', '2022-12-06 11:08:41', '2022-12-06 11:08:41');
INSERT INTO `authorization_api` VALUES (2, '基础分组', '登录成功', 'GET', '/api.admin.v1.Admin/LoginSuccess', '2022-12-06 11:09:08', '2022-12-06 11:09:08');
INSERT INTO `authorization_api` VALUES (3, '基础分组', '获取登录用户信息', 'GET', '/api.admin.v1.Admin/GetAdministratorInfo', '2022-12-06 11:09:37', '2022-12-06 11:09:37');
INSERT INTO `authorization_api` VALUES (4, '基础分组', '退出登录', 'POST', '/api.admin.v1.Admin/Logout', '2022-12-06 11:10:21', '2022-12-06 11:10:21');
INSERT INTO `authorization_api` VALUES (5, '基础分组', '获取菜单栏信息', 'GET', '/api.admin.v1.Admin/GetRoleMenuTree', '2022-12-06 11:13:11', '2022-12-06 11:13:47');
INSERT INTO `authorization_api` VALUES (6, '基础分组', '获取按钮列表信息', 'GET', '/api.admin.v1.Admin/GetRoleMenuBtn', '2022-12-06 11:14:45', '2022-12-06 11:14:45');
INSERT INTO `authorization_api` VALUES (7, '管理员管理', '新增', 'POST', '/api.admin.v1.Admin/CreateAdministrator', '2022-12-06 11:16:14', '2022-12-06 11:24:28');
INSERT INTO `authorization_api` VALUES (8, '管理员管理', '编辑', 'PUT', '/api.admin.v1.Admin/UpdateAdministrator', '2022-12-06 11:17:40', '2022-12-06 11:24:32');
INSERT INTO `authorization_api` VALUES (9, '管理员管理', '列表', 'GET', '/api.admin.v1.Admin/GetAdministratorList', '2022-12-06 11:18:42', '2022-12-06 11:24:36');
INSERT INTO `authorization_api` VALUES (10, '管理员管理', '查看', 'GET', '/api.admin.v1.Admin/GetAdministrator', '2022-12-06 11:19:14', '2022-12-06 11:24:40');
INSERT INTO `authorization_api` VALUES (11, '管理员管理', '删除', 'DELETE', '/api.admin.v1.Admin/DeleteAdministrator', '2022-12-06 11:19:59', '2022-12-06 11:24:45');
INSERT INTO `authorization_api` VALUES (12, '管理员管理', '恢复', 'PATCH', '/api.admin.v1.Admin/RecoverAdministrator', '2022-12-06 11:20:41', '2022-12-06 11:24:49');
INSERT INTO `authorization_api` VALUES (13, '管理员管理', '禁用', 'PATCH', '/api.admin.v1.Admin/forbidAdministrator', '2022-12-06 11:21:27', '2022-12-06 11:24:53');
INSERT INTO `authorization_api` VALUES (14, '管理员管理', '解禁', 'PATCH', '/api.admin.v1.Admin/approveAdministrator', '2022-12-06 11:22:26', '2022-12-06 11:24:57');
INSERT INTO `authorization_api` VALUES (15, '菜单管理', '新增', 'GET', '/api.admin.v1.Admin/CreateMenu', '2022-12-06 11:27:05', '2022-12-06 11:27:05');
INSERT INTO `authorization_api` VALUES (16, '菜单管理', '编辑', 'PUT', '/api.admin.v1.Admin/UpdateMenu', '2022-12-06 11:27:48', '2022-12-06 11:27:48');
INSERT INTO `authorization_api` VALUES (17, '菜单管理', '列表', 'GET', '/api.admin.v1.Admin/GetMenuTree', '2022-12-07 16:52:30', '2022-12-07 16:52:32');
INSERT INTO `authorization_api` VALUES (18, '菜单管理', '删除', 'DELETE', '/api.admin.v1.Admin/DeleteMenu', '2022-12-06 11:29:24', '2022-12-06 11:29:24');
INSERT INTO `authorization_api` VALUES (21, '角色管理', '新增', 'POST', '/api.admin.v1.Admin/CreateRole', '2022-12-06 12:03:52', '2022-12-06 12:03:52');
INSERT INTO `authorization_api` VALUES (22, '角色管理', '编辑', 'PUT', '/api.admin.v1.Admin/UpdateRole', '2022-12-06 12:04:40', '2022-12-06 12:04:40');
INSERT INTO `authorization_api` VALUES (23, '角色管理', '删除', 'DELETE', '/api.admin.v1.Admin/DeleteRole', '2022-12-06 12:05:11', '2022-12-06 12:05:11');
INSERT INTO `authorization_api` VALUES (24, '角色管理', '列表', 'GET', '/api.admin.v1.Admin/GetRoleList', '2022-12-06 12:05:42', '2022-12-06 12:05:42');
INSERT INTO `authorization_api` VALUES (25, '角色管理', '获取全部权限api', 'GET', '/api.admin.v1.Admin/GetApiAll', '2022-12-06 12:08:04', '2022-12-06 12:08:04');
INSERT INTO `authorization_api` VALUES (26, '角色管理', '获取已有权限api', 'GET', '/api.admin.v1.Admin/GetPolicies', '2022-12-06 12:09:22', '2022-12-06 12:09:22');
INSERT INTO `authorization_api` VALUES (27, '角色管理', '设置角色菜单', 'POST', '/api.admin.v1.Admin/SetRoleMenu', '2022-12-06 12:10:04', '2022-12-06 12:10:04');
INSERT INTO `authorization_api` VALUES (28, '角色管理', '设置角色API', 'POST', '/api.admin.v1.Admin/UpdatePolicies', '2022-12-06 12:12:10', '2022-12-06 12:12:10');
INSERT INTO `authorization_api` VALUES (30, '角色管理', '设置角色菜单按钮', 'POST', '/api.admin.v1.Admin/SetRoleMenuBtn', '2022-12-06 12:13:20', '2022-12-06 12:13:20');
INSERT INTO `authorization_api` VALUES (31, 'API管理', '新增', 'POST', '/api.admin.v1.Admin/CreateApi', '2022-12-06 12:14:07', '2022-12-06 12:14:07');
INSERT INTO `authorization_api` VALUES (32, 'API管理', '编辑', 'PUT', '/api.admin.v1.Admin/UpdateApi', '2022-12-06 12:14:35', '2022-12-06 12:14:35');
INSERT INTO `authorization_api` VALUES (33, 'API管理', '列表', 'GET', '/api.admin.v1.Admin/GetApiList', '2022-12-06 12:15:21', '2022-12-06 12:15:21');
INSERT INTO `authorization_api` VALUES (34, 'API管理', '删除', 'DELETE', '/api.admin.v1.Admin/DeleteApi', '2022-12-06 12:16:19', '2022-12-06 12:16:19');

-- ----------------------------
-- Table structure for authorization_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `authorization_menu_btns`;
CREATE TABLE `authorization_menu_btns`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `menu_id` int(0) NOT NULL COMMENT '菜单id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '按钮名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '按钮描述',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_menu_btns
-- ----------------------------
INSERT INTO `authorization_menu_btns` VALUES (1, 12, '新增', '新增管理员', '2022-12-06 11:05:36.421', '2022-12-06 11:05:36.421');
INSERT INTO `authorization_menu_btns` VALUES (2, 12, '编辑', '编辑管理员', '2022-12-06 11:05:36.421', '2022-12-06 11:05:36.421');
INSERT INTO `authorization_menu_btns` VALUES (3, 12, '删除/恢复', '删除/恢复管理员', '2022-12-06 11:05:36.421', '2022-12-06 11:05:36.421');
INSERT INTO `authorization_menu_btns` VALUES (4, 12, '禁用/解禁', '禁用/解禁管理员', '2022-12-06 11:05:36.421', '2022-12-06 11:05:36.421');
INSERT INTO `authorization_menu_btns` VALUES (20, 2, '新增', '新增管理员', '2022-12-06 14:12:38.442', '2022-12-06 14:12:38.442');
INSERT INTO `authorization_menu_btns` VALUES (21, 2, '编辑', '编辑管理员', '2022-12-06 14:12:38.444', '2022-12-06 14:12:38.444');
INSERT INTO `authorization_menu_btns` VALUES (22, 2, '禁用/解禁', '禁用/解禁管理员', '2022-12-06 14:12:38.446', '2022-12-06 14:12:38.446');
INSERT INTO `authorization_menu_btns` VALUES (23, 2, '删除/恢复', '删除/恢复管理员', '2022-12-06 14:12:38.448', '2022-12-06 14:12:38.448');
INSERT INTO `authorization_menu_btns` VALUES (24, 3, '新增', '新增菜单', '2022-12-06 14:18:02.694', '2022-12-06 14:18:02.694');
INSERT INTO `authorization_menu_btns` VALUES (25, 3, '编辑', '编辑菜单', '2022-12-06 14:18:02.695', '2022-12-06 14:18:02.695');
INSERT INTO `authorization_menu_btns` VALUES (26, 3, '删除', '删除菜单', '2022-12-06 14:18:02.697', '2022-12-06 14:18:02.697');
INSERT INTO `authorization_menu_btns` VALUES (27, 4, '新增', '新增角色', '2022-12-06 14:22:46.669', '2022-12-06 14:22:46.669');
INSERT INTO `authorization_menu_btns` VALUES (28, 4, '编辑', '编辑角色', '2022-12-06 14:22:46.671', '2022-12-06 14:22:46.671');
INSERT INTO `authorization_menu_btns` VALUES (29, 4, '删除', '删除角色', '2022-12-06 14:22:46.672', '2022-12-06 14:22:46.672');
INSERT INTO `authorization_menu_btns` VALUES (30, 4, '设置权限', '设置角色权限', '2022-12-06 14:22:46.673', '2022-12-06 14:22:46.673');
INSERT INTO `authorization_menu_btns` VALUES (31, 4, '设置菜单权限', '设置角色菜单权限', '2022-12-06 14:22:46.674', '2022-12-06 14:22:46.674');
INSERT INTO `authorization_menu_btns` VALUES (32, 4, '设置API权限', '设置角色API权限', '2022-12-06 14:22:46.675', '2022-12-06 14:22:46.675');
INSERT INTO `authorization_menu_btns` VALUES (33, 4, '设置按钮权限', '设置角色按钮权限', '2022-12-06 14:22:46.677', '2022-12-06 14:22:46.677');
INSERT INTO `authorization_menu_btns` VALUES (34, 5, '新增', '新增API', '2022-12-06 14:23:26.970', '2022-12-06 14:23:26.970');
INSERT INTO `authorization_menu_btns` VALUES (35, 5, '编辑', '编辑API', '2022-12-06 14:23:26.971', '2022-12-06 14:23:26.971');
INSERT INTO `authorization_menu_btns` VALUES (36, 5, '删除', '删除API', '2022-12-06 14:23:26.972', '2022-12-06 14:23:26.972');
INSERT INTO `authorization_menu_btns` VALUES (39, 36, '有权限', '有权限按钮', '2022-12-07 11:44:35.456', '2022-12-07 11:44:35.456');
INSERT INTO `authorization_menu_btns` VALUES (40, 36, '无权限', '无权限按钮', '2022-12-07 11:44:35.456', '2022-12-07 11:44:35.456');

-- ----------------------------
-- Table structure for authorization_menus
-- ----------------------------
DROP TABLE IF EXISTS `authorization_menus`;
CREATE TABLE `authorization_menus`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `parent_id` int(0) NOT NULL COMMENT '父级菜单id',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色父级id数组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由名称',
  `hidden` tinyint(1) NOT NULL COMMENT '是否隐藏 0否1是',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '前端文件路径',
  `sort` int(0) NOT NULL COMMENT '排序',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'icon图标',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '直接跳转',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_menus
-- ----------------------------
INSERT INTO `authorization_menus` VALUES (1, 0, '0', '/system', 'system', 0, '#', 1, '系统管理', 'el-icon-s-tools', '', '2022-11-15 14:56:16.000', '2022-11-15 14:56:18.000');
INSERT INTO `authorization_menus` VALUES (2, 1, '1', '/system/adminstrator', 'administrator', 0, '/system/administrator/index', 1, '管理员管理', 'el-icon-user-solid', '', '2022-11-15 14:58:41.000', '2022-12-06 14:12:38.439');
INSERT INTO `authorization_menus` VALUES (3, 1, '1', '/system/auth/menu', 'menu', 0, '/system/auth/menu/index', 2, '菜单管理', 'el-icon-notebook-2', '', '2022-12-01 15:48:57.000', '2022-12-06 14:18:02.691');
INSERT INTO `authorization_menus` VALUES (4, 1, '1', '/system/auth/role', 'role', 0, '/system/auth/role/index', 3, '角色管理', 'el-icon-user', '', '2022-11-24 15:33:30.000', '2022-12-06 14:22:46.665');
INSERT INTO `authorization_menus` VALUES (5, 1, '1', '/system/auth/api', 'api', 0, '/system/auth/api/index', 4, 'API管理', 'el-icon-setting', '', '2022-11-23 14:42:04.000', '2022-12-06 14:23:26.966');
INSERT INTO `authorization_menus` VALUES (35, 0, '0', '/example', 'example', 0, '#', 1, '测试菜单', 'el-icon-s-flag', '', '2022-12-07 10:36:58.676', '2022-12-07 11:43:33.458');
INSERT INTO `authorization_menus` VALUES (36, 35, '35', '/example/permission', 'permission_test', 0, '/example/index', 1, '按钮权限测试', 'el-icon-s-tools', '', '2022-12-07 11:44:35.452', '2022-12-07 11:44:35.452');

-- ----------------------------
-- Table structure for authorization_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `authorization_role_menu`;
CREATE TABLE `authorization_role_menu`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` int(0) NOT NULL COMMENT '角色id',
  `menu_id` int(0) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_role_menu
-- ----------------------------
INSERT INTO `authorization_role_menu` VALUES (74, 1, 1);
INSERT INTO `authorization_role_menu` VALUES (75, 1, 2);
INSERT INTO `authorization_role_menu` VALUES (76, 1, 3);
INSERT INTO `authorization_role_menu` VALUES (77, 1, 4);
INSERT INTO `authorization_role_menu` VALUES (78, 1, 5);
INSERT INTO `authorization_role_menu` VALUES (79, 1, 35);
INSERT INTO `authorization_role_menu` VALUES (80, 1, 36);
INSERT INTO `authorization_role_menu` VALUES (81, 39, 35);
INSERT INTO `authorization_role_menu` VALUES (82, 39, 36);

-- ----------------------------
-- Table structure for authorization_role_menu_btn
-- ----------------------------
DROP TABLE IF EXISTS `authorization_role_menu_btn`;
CREATE TABLE `authorization_role_menu_btn`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` int(0) NOT NULL COMMENT '角色id',
  `menu_id` int(0) NOT NULL COMMENT '菜单id',
  `btn_id` int(0) NOT NULL COMMENT '按钮id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色菜单按钮表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_role_menu_btn
-- ----------------------------
INSERT INTO `authorization_role_menu_btn` VALUES (2, 1, 2, 20);
INSERT INTO `authorization_role_menu_btn` VALUES (3, 1, 2, 21);
INSERT INTO `authorization_role_menu_btn` VALUES (4, 1, 2, 22);
INSERT INTO `authorization_role_menu_btn` VALUES (5, 1, 2, 23);
INSERT INTO `authorization_role_menu_btn` VALUES (6, 1, 3, 24);
INSERT INTO `authorization_role_menu_btn` VALUES (7, 1, 3, 25);
INSERT INTO `authorization_role_menu_btn` VALUES (8, 1, 3, 26);
INSERT INTO `authorization_role_menu_btn` VALUES (9, 1, 4, 27);
INSERT INTO `authorization_role_menu_btn` VALUES (10, 1, 4, 28);
INSERT INTO `authorization_role_menu_btn` VALUES (11, 1, 4, 29);
INSERT INTO `authorization_role_menu_btn` VALUES (12, 1, 4, 30);
INSERT INTO `authorization_role_menu_btn` VALUES (13, 1, 4, 31);
INSERT INTO `authorization_role_menu_btn` VALUES (14, 1, 4, 32);
INSERT INTO `authorization_role_menu_btn` VALUES (15, 1, 4, 33);
INSERT INTO `authorization_role_menu_btn` VALUES (16, 1, 5, 34);
INSERT INTO `authorization_role_menu_btn` VALUES (17, 1, 5, 35);
INSERT INTO `authorization_role_menu_btn` VALUES (18, 1, 5, 36);
INSERT INTO `authorization_role_menu_btn` VALUES (21, 1, 36, 39);
INSERT INTO `authorization_role_menu_btn` VALUES (22, 1, 36, 40);
INSERT INTO `authorization_role_menu_btn` VALUES (23, 39, 36, 39);

-- ----------------------------
-- Table structure for authorization_roles
-- ----------------------------
DROP TABLE IF EXISTS `authorization_roles`;
CREATE TABLE `authorization_roles`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(0) NOT NULL COMMENT '父级角色id',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色父级id数组',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_unique_idx`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authorization_roles
-- ----------------------------
INSERT INTO `authorization_roles` VALUES (1, 0, '0', '超级管理员', '2022-09-07 01:12:43', '2022-11-25 14:48:48');
INSERT INTO `authorization_roles` VALUES (39, 0, '0', '测试管理员', '2022-12-06 11:57:30', '2022-12-06 14:24:40');
INSERT INTO `authorization_roles` VALUES (40, 0, '0', '游客', '2022-12-07 11:25:42', '2022-12-07 11:25:42');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (50, 'g', '2', '测试管理员', '', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (278, 'g', 'guest', '游客', '', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (277, 'g', 'test', '测试管理员', '', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'g', 'zhouqi', '超级管理员', '', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (357, 'p', '测试管理员', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (355, 'p', '测试管理员', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (359, 'p', '测试管理员', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (356, 'p', '测试管理员', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (354, 'p', '测试管理员', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (358, 'p', '测试管理员', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (363, 'p', '游客', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (361, 'p', '游客', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (365, 'p', '游客', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (362, 'p', '游客', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (360, 'p', '游客', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (364, 'p', '游客', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (327, 'p', '超级管理员', '/api.admin.v1.Admin/approveAdministrator', 'PATCH', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (326, 'p', '超级管理员', '/api.admin.v1.Admin/CreateAdministrator', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (319, 'p', '超级管理员', '/api.admin.v1.Admin/CreateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (337, 'p', '超级管理员', '/api.admin.v1.Admin/CreateMenu', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (338, 'p', '超级管理员', '/api.admin.v1.Admin/CreateRole', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (332, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteAdministrator', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (316, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteApi', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (334, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteMenu', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (340, 'p', '超级管理员', '/api.admin.v1.Admin/DeleteRole', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (333, 'p', '超级管理员', '/api.admin.v1.Admin/forbidAdministrator', 'PATCH', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (331, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministrator', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (323, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministratorInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (330, 'p', '超级管理员', '/api.admin.v1.Admin/GetAdministratorList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (342, 'p', '超级管理员', '/api.admin.v1.Admin/GetApiAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (317, 'p', '超级管理员', '/api.admin.v1.Admin/GetApiList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (335, 'p', '超级管理员', '/api.admin.v1.Admin/GetMenuTree', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (343, 'p', '超级管理员', '/api.admin.v1.Admin/GetPolicies', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (341, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (321, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleMenuBtn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (325, 'p', '超级管理员', '/api.admin.v1.Admin/GetRoleMenuTree', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (322, 'p', '超级管理员', '/api.admin.v1.Admin/Login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (320, 'p', '超级管理员', '/api.admin.v1.Admin/LoginSuccess', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (324, 'p', '超级管理员', '/api.admin.v1.Admin/Logout', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (328, 'p', '超级管理员', '/api.admin.v1.Admin/RecoverAdministrator', 'PATCH', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (344, 'p', '超级管理员', '/api.admin.v1.Admin/SetRoleMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (346, 'p', '超级管理员', '/api.admin.v1.Admin/SetRoleMenuBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (329, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateAdministrator', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (318, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateApi', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (336, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateMenu', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (345, 'p', '超级管理员', '/api.admin.v1.Admin/UpdatePolicies', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (339, 'p', '超级管理员', '/api.admin.v1.Admin/UpdateRole', 'PUT', '', '', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
