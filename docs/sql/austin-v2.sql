/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.127.128_mysql
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : 192.168.127.128:3306
 Source Schema         : austin-v2

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 11/04/2023 23:10:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for la_system_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `la_system_admin_role`;
CREATE TABLE `la_system_admin_role`  (
  `role_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `admin_id` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of la_system_admin_role
-- ----------------------------
INSERT INTO `la_system_admin_role` VALUES (1, 1);

-- ----------------------------
-- Table structure for la_system_auth_admin
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_admin`;
CREATE TABLE `la_system_auth_admin`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色主键',
  `salt` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `is_multipoint` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '多端登录: 0=否, 1=是',
  `is_disable` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of la_system_auth_admin
-- ----------------------------
INSERT INTO `la_system_auth_admin` VALUES (1, 'admin', 'zoujiangyong@qq.com', 'de3a9c46bca749cbc9c0aca39d6493b7', 'https://go-admin.likeadmin.cn/api/static/backend_avatar.png', '0', 'EAXKUFKFYS', 0, 1, 0, 0, '127.0.0.1', 1660641347, 1642321599, 1660287325, 0);
INSERT INTO `la_system_auth_admin` VALUES (2, 'lis', 'lis', '73af314e8fb14e3284d6c1625ee663a5', 'https://go-admin.likeadmin.cn/api/static/backend_avatar.png', '', 'MJVXCNFSPR', 0, 0, 0, 0, '', 0, 1680856353, 0, 0);

-- ----------------------------
-- Table structure for la_system_auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_menu`;
CREATE TABLE `la_system_auth_menu`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单',
  `menu_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限类型: M=目录，C=菜单，A=按钮',
  `menu_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `menu_sort` smallint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单排序',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `paths` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端组件',
  `selected` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选中路径',
  `params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_cache` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存: 0=否, 1=是',
  `is_show` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
  `is_disable` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 779 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of la_system_auth_menu
-- ----------------------------
INSERT INTO `la_system_auth_menu` VALUES (1, 0, 'C', '工作台', 'el-icon-Monitor', 50, 'index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, 1650341765, 1668672757);
INSERT INTO `la_system_auth_menu` VALUES (100, 0, 'M', '权限管理', 'el-icon-Lock', 44, '', 'permission', '', '', '', 0, 1, 0, 1650341765, 1662626201);
INSERT INTO `la_system_auth_menu` VALUES (101, 100, 'C', '管理员', 'local-icon-wode', 0, 'system:admin:list', 'admin', 'permission/admin/index', '', '', 1, 1, 0, 1650341765, 1663301404);
INSERT INTO `la_system_auth_menu` VALUES (102, 101, 'A', '管理员详情', '', 0, 'system:admin:detail', '', '', '', '', 0, 1, 0, 1650341765, 1660201785);
INSERT INTO `la_system_auth_menu` VALUES (103, 101, 'A', '管理员新增', '', 0, 'system:admin:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (104, 101, 'A', '管理员编辑', '', 0, 'system:admin:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (105, 101, 'A', '管理员删除', '', 0, 'system:admin:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (106, 101, 'A', '管理员状态', '', 0, 'system:admin:disable', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (110, 100, 'C', '角色管理', 'el-icon-Female', 0, 'system:role:list', 'role', 'permission/role/index', '', '', 1, 1, 0, 1650341765, 1663301451);
INSERT INTO `la_system_auth_menu` VALUES (111, 110, 'A', '角色详情', '', 0, 'system:role:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (112, 110, 'A', '角色新增', '', 0, 'system:role:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (113, 110, 'A', '角色编辑', '', 0, 'system:role:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (114, 110, 'A', '角色删除', '', 0, 'system:role:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (120, 100, 'C', '菜单管理', 'el-icon-Operation', 0, 'system:menu:list', 'menu', 'permission/menu/index', '', '', 1, 1, 0, 1650341765, 1663301388);
INSERT INTO `la_system_auth_menu` VALUES (121, 120, 'A', '菜单详情', '', 0, 'system:menu:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (122, 120, 'A', '菜单新增', '', 0, 'system:menu:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (123, 120, 'A', '菜单编辑', '', 0, 'system:menu:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (124, 120, 'A', '菜单删除', '', 0, 'system:menu:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1660636870);
INSERT INTO `la_system_auth_menu` VALUES (201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (202, 201, 'A', '文件列表', '', 0, 'albums:albumList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (203, 201, 'A', '文件命名', '', 0, 'albums:albumRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (204, 201, 'A', '文件移动', '', 0, 'albums:albumMove', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (205, 201, 'A', '文件删除', '', 0, 'albums:albumDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (206, 201, 'A', '分类列表', '', 0, 'albums:cateList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (207, 201, 'A', '分类新增', '', 0, 'albums:cateAdd', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (208, 201, 'A', '分类命名', '', 0, 'albums:cateRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (209, 201, 'A', '分类删除', '', 0, 'albums:cateDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (216, 215, 'A', '上传图片', '', 0, 'upload:image', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (217, 215, 'A', '上传视频', '', 0, 'upload:video', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'setting', '', '', '', 0, 1, 0, 1650341765, 1662626322);
INSERT INTO `la_system_auth_menu` VALUES (501, 500, 'M', '网站设置', 'el-icon-Basketball', 10, '', 'website', '', '', '', 0, 1, 0, 1650341765, 1663233572);
INSERT INTO `la_system_auth_menu` VALUES (502, 501, 'C', '网站信息', '', 0, 'setting:website:detail', 'information', 'setting/website/information', '', '', 0, 1, 0, 1650341765, 1660202218);
INSERT INTO `la_system_auth_menu` VALUES (503, 502, 'A', '保存配置', '', 0, 'setting:website:save', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (505, 501, 'C', '网站备案', '', 0, 'setting:copyright:detail', 'filing', 'setting/website/filing', '', '', 0, 1, 0, 1650341765, 1660202294);
INSERT INTO `la_system_auth_menu` VALUES (506, 505, 'A', '备案保存', '', 0, 'setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (510, 501, 'C', '政策协议', '', 0, 'setting:protocol:detail', 'protocol', 'setting/website/protocol', '', '', 0, 1, 0, 1660027606, 1660202312);
INSERT INTO `la_system_auth_menu` VALUES (511, 510, 'A', '协议保存', '', 0, 'setting:protocol:save', '', '', '', '', 0, 0, 0, 1660027606, 1663670865);
INSERT INTO `la_system_auth_menu` VALUES (550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, 1650341765, 1660202466);
INSERT INTO `la_system_auth_menu` VALUES (551, 550, 'C', '系统环境', '', 0, 'monitor:server', 'environment', 'setting/system/environment', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (552, 550, 'C', '系统缓存', '', 0, 'monitor:cache', 'cache', 'setting/system/cache', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (553, 550, 'C', '系统日志', '', 0, 'system:log:operate', 'journal', 'setting/system/journal', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (555, 500, 'C', '存储设置', 'el-icon-FolderOpened', 6, 'setting:storage:list', 'storage', 'setting/storage/index', '', '', 0, 1, 0, 1650341765, 1663312996);
INSERT INTO `la_system_auth_menu` VALUES (556, 555, 'A', '保存配置', '', 0, 'setting:storage:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu` VALUES (700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, 1660203293, 1663300847);
INSERT INTO `la_system_auth_menu` VALUES (701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0, 1660203402, 1663301493);
INSERT INTO `la_system_auth_menu` VALUES (778, 0, 'M', '222', '', 0, '', '', '', '', '', 0, 0, 0, 1680675750, 1680676198);

-- ----------------------------
-- Table structure for la_system_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_role`;
CREATE TABLE `la_system_auth_role`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` smallint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色排序',
  `is_disable` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `create_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `menu_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of la_system_auth_role
-- ----------------------------
INSERT INTO `la_system_auth_role` VALUES (1, '超級管理员', '超級管理员', 100, 0, 1668679451, 1680859349, '1,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,700,701,778,200,201,202,203,204,205,206,207,208,209,215,216,217,500,501,502,503,505,506,510,511,555,556,550,551,552,553');
INSERT INTO `la_system_auth_role` VALUES (2, '审核员', '超管2', 0, 0, 0, 0, '*');
INSERT INTO `la_system_auth_role` VALUES (3, '审核员22', '审核数据', 100, 0, 1680680188, 1680680662, '*');

-- ----------------------------
-- Table structure for message_template
-- ----------------------------
DROP TABLE IF EXISTS `message_template`;
CREATE TABLE `message_template`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `audit_status` int(0) NOT NULL DEFAULT 0 COMMENT '当前消息审核状态： 10.待审核 20.审核成功 30.被拒绝',
  `id_type` int(0) NOT NULL DEFAULT 0 COMMENT '消息的发送ID类型：10. userId 20.did 30.手机号 40.openId 50.email 60.企业微信userId',
  `send_channel` int(0) NOT NULL DEFAULT 0 COMMENT '消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信',
  `template_type` int(0) NOT NULL DEFAULT 0 COMMENT '10.运营类 20.技术类接口调用',
  `msg_type` int(0) NOT NULL DEFAULT 0 COMMENT '10.通知类消息 20.营销类消息 30.验证码类消息',
  `shield_type` int(0) NOT NULL DEFAULT 0 COMMENT '10.夜间不屏蔽 20.夜间屏蔽 30.夜间屏蔽(次日早上9点发送)',
  `msg_content` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '消息内容 占位符用{$var}表示',
  `send_account` int(0) NOT NULL DEFAULT 0 COMMENT '发送账号 一个渠道下可存在多个账号',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '更新者',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否删除：0.删除 1.正常 2.禁用',
  `create_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deduplication_config` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '数据去重配置',
  `template_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送消息的模版ID',
  `sms_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信渠道 send_channel=30的时候有用  tencent腾讯云  aliyun阿里云 yunpian云片',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_channel`(`send_channel`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '消息模板信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message_template
-- ----------------------------
INSERT INTO `message_template` VALUES (1, '买一送十活动', 10, 30, 30, 20, 20, 30, '{\"content\":\"恭喜你:{$content}\",\"url\":\"\",\"title\":\"\"}', 10, 'Java3y', 'Java3y', 0, 1646274112, 1646275242, '', '', '');
INSERT INTO `message_template` VALUES (2, '校招信息', 10, 50, 40, 20, 10, 0, '{\"content\":\"你已成功获取到offer 内容:{$content}\",\"url\":\"\",\"title\":\"招聘通知\"}', 1, 'Java3y', 'Java3y', 0, 1646274195, 1646274195, '', '', '');
INSERT INTO `message_template` VALUES (3, '验证码通知', 10, 30, 30, 20, 30, 0, '{\"content\":\"{$content}\",\"url\":\"\",\"title\":\"\"}', 10, 'Java3y', 'Java3y', 0, 1646275213, 1646275213, '', '', 'yunpian');
INSERT INTO `message_template` VALUES (4, '微信测试通知', 10, 40, 50, 20, 10, 0, '', 2, '', '', 0, 1646275213, 1646275213, '', 'OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM', '');
INSERT INTO `message_template` VALUES (5, '钉钉测试通知', 10, 40, 70, 20, 10, 0, '{\"content\":\"钉钉测试消息:\\n内容:{$content}\",\"url\":\"\",\"title\":\"\"}', 3, '', '', 0, 1646275213, 1646275213, '', '', '');

-- ----------------------------
-- Table structure for msg_record
-- ----------------------------
DROP TABLE IF EXISTS `msg_record`;
CREATE TABLE `msg_record`  (
  `id` bigint(0) NOT NULL,
  `message_template_id` bigint(0) NOT NULL DEFAULT 0 COMMENT '消息模板ID',
  `request_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '唯一请求 ID',
  `receiver` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接收人',
  `msg_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公众号消息id',
  `channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推送结果信息',
  `send_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息http 发送时间',
  `create_at` datetime(0) NOT NULL,
  `start_consume_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '开始消费时间',
  `end_consume_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '结束消费时间',
  `consume_since_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消费间距时间',
  `send_since_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'http->mq消费结束间距时间',
  `task_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `request_id`(`request_id`) USING BTREE,
  INDEX `message_template_id`(`message_template_id`) USING BTREE,
  INDEX `channel`(`channel`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of msg_record
-- ----------------------------
INSERT INTO `msg_record` VALUES (1625459874345259008, 4, '1625459871027564544', 'okEEF6WB92HO14qdy0Nosq62OVyY', '2798146648328945668', 'official_accounts', '推送成功', '2023-02-14 19:40:07', '2023-02-14 19:40:08', '2023-02-14 19:40:07', '2023-02-14 19:40:08', '616.983ms', '790.1012ms', '{\"request_id\":\"1625459871027564544\",\"message_template_id\":4,\"business_id\":2000000420230214,\"receiver\":[\"okEEF6WB92HO14qdy0Nosq62OVyY\"],\"id_type\":40,\"send_channel\":50,\"template_type\":20,\"msg_type\":10,\"shield_type\":0,\"content_model\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"},\"mini_program\":{\"appid\":\"\",\"pagepath\":\"\"},\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"url\":\"\"},\"send_account\":2,\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"sms_channel\":\"\",\"start_consume_at\":\"2023-02-14T19:40:07.668096+08:00\",\"send_at\":\"2023-02-14T19:40:07.4949778+08:00\",\"message_param\":{\"receiver\":\"okEEF6WB92HO14qdy0Nosq62OVyY\",\"variables\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"}},\"extra\":{}}}');
INSERT INTO `msg_record` VALUES (1625459878262738944, 4, '1625459874982793216', 'okEEF6WB92HO14qdy0Nosq62OVyY', '2798146664082751493', 'official_accounts', '推送成功', '2023-02-14 19:40:08', '2023-02-14 19:40:09', '2023-02-14 19:40:08', '2023-02-14 19:40:09', '543.0679ms', '782.3122ms', '{\"request_id\":\"1625459874982793216\",\"message_template_id\":4,\"business_id\":2000000420230214,\"receiver\":[\"okEEF6WB92HO14qdy0Nosq62OVyY\"],\"id_type\":40,\"send_channel\":50,\"template_type\":20,\"msg_type\":10,\"shield_type\":0,\"content_model\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"},\"mini_program\":{\"appid\":\"\",\"pagepath\":\"\"},\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"url\":\"\"},\"send_account\":2,\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"sms_channel\":\"\",\"start_consume_at\":\"2023-02-14T19:40:08.6768027+08:00\",\"send_at\":\"2023-02-14T19:40:08.4375584+08:00\",\"message_param\":{\"receiver\":\"okEEF6WB92HO14qdy0Nosq62OVyY\",\"variables\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"}},\"extra\":{}}}');
INSERT INTO `msg_record` VALUES (1625459959212806144, 4, '1625459954225778688', 'okEEF6WB92HO14qdy0Nosq62OVyY', '2798146987883020289', 'official_accounts', '推送成功', '2023-02-14 19:40:27', '2023-02-14 19:40:29', '2023-02-14 19:40:27', '2023-02-14 19:40:28', '649.7771ms', '1.1893555s', '{\"request_id\":\"1625459954225778688\",\"message_template_id\":4,\"business_id\":2000000420230214,\"receiver\":[\"okEEF6WB92HO14qdy0Nosq62OVyY\"],\"id_type\":40,\"send_channel\":50,\"template_type\":20,\"msg_type\":10,\"shield_type\":0,\"content_model\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"},\"mini_program\":{\"appid\":\"\",\"pagepath\":\"\"},\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"url\":\"\"},\"send_account\":2,\"template_sn\":\"OwqP1h3N8QNBvdTim7MTg9fo40RARsiplsvj_d7FtXM\",\"sms_channel\":\"\",\"start_consume_at\":\"2023-02-14T19:40:27.870003+08:00\",\"send_at\":\"2023-02-14T19:40:27.3304246+08:00\",\"message_param\":{\"receiver\":\"okEEF6WB92HO14qdy0Nosq62OVyY\",\"variables\":{\"data\":{\"order_no\":\"DD12345678\",\"time\":\"2022-01-11 10:00:00\"}},\"extra\":{}}}');

-- ----------------------------
-- Table structure for send_account
-- ----------------------------
DROP TABLE IF EXISTS `send_account`;
CREATE TABLE `send_account`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `send_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送渠道',
  `config` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账户配置',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号名称',
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of send_account
-- ----------------------------
INSERT INTO `send_account` VALUES (1, '40', '{\"host\":\"smtp.qq.com\",\"port\":25,\"username\":\"test@qq.com\",\"password\":\"tesxxxx\"}', '邮箱账号', 0);
INSERT INTO `send_account` VALUES (2, '50', '{\"app_id\":\"xx\",\"app_secret\":\"xxx\",\"token\":\"weixin\"}', '微信公众号配置', 0);
INSERT INTO `send_account` VALUES (3, '80', '{\"access_token\":\"access_token\",\"secret\":\"secret\"}', '钉钉自定义机器人', 0);

-- ----------------------------
-- Table structure for sms_record
-- ----------------------------
DROP TABLE IF EXISTS `sms_record`;
CREATE TABLE `sms_record`  (
  `id` bigint(0) NOT NULL,
  `message_template_id` bigint(0) NOT NULL DEFAULT 0 COMMENT '消息模板ID',
  `phone` bigint(0) NOT NULL DEFAULT 0 COMMENT '手机号',
  `supplier_id` tinyint(0) NOT NULL DEFAULT 0 COMMENT '发送短信渠道商的ID',
  `supplier_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送短信渠道商的名称',
  `msg_content` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '短信发送的内容',
  `series_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '下发批次的ID',
  `charging_num` tinyint(0) NOT NULL DEFAULT 0 COMMENT '计费条数',
  `report_content` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '回执内容',
  `status` tinyint(0) NOT NULL DEFAULT 0 COMMENT '短信状态： 10.发送 20.成功 30.失败',
  `send_date` int(0) NOT NULL DEFAULT 0 COMMENT '发送日期：20211112',
  `create_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `request_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '唯一请求 ID',
  `biz_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务id',
  `send_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '短信渠道 tencent腾讯云  aliyun阿里云 yunpian云片',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_send_date`(`send_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '短信记录信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sms_record
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
