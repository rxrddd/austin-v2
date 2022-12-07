/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : 127.0.0.1:3306
 Source Schema         : kratos-base-project-administrator

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 07/12/2022 17:13:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_administrator
-- ----------------------------
DROP TABLE IF EXISTS `sys_administrator`;
CREATE TABLE `sys_administrator`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码盐',
  `mobile` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号码',
  `nickname` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` char(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像地址',
  `status` tinyint(0) NOT NULL COMMENT '用户状态 1正常 2冻结',
  `role` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '当前角色',
  `last_login_time` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '上次登陆时间',
  `last_login_ip` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '上次登陆ip',
  `created_at` timestamp(0) NOT NULL COMMENT '创建时间',
  `updated_at` timestamp(0) NOT NULL COMMENT '更新时间',
  `deleted_at` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '删除时间 ',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `mobile_unique_idx`(`mobile`) USING BTREE COMMENT '手机号唯一索引',
  UNIQUE INDEX `username_unique_idx`(`username`) USING BTREE COMMENT '用户名唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_administrator
-- ----------------------------
INSERT INTO `sys_administrator` VALUES (18, 'admin', 'b9819e53ed8ea2b2b422ff1d2f1317ca', 'e701bf4e804804773099b4b20130d418', '18158445331', '卡牌', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '超级管理员', '2022-12-07 17:05:08', '127.0.0.1:62201', '2022-08-17 16:15:17', '2022-11-22 17:41:38', '');
INSERT INTO `sys_administrator` VALUES (27, 'test', '8ab138656a71b1e001aa12cc7298f901', 'cb1461e3e59ec7a2237bf5f5fa105ab5', '18158445332', '测试', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '测试管理员', '2022-12-07 17:05:50', '127.0.0.1:59530', '2022-12-07 11:28:48', '2022-12-07 11:28:48', '');
INSERT INTO `sys_administrator` VALUES (28, 'guest', 'a5ac55c657800544fa68377d4bb64505', 'dabba3032e7d0c51c48e7ca794e589b8', '18158445333', '游客', 'https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg', 1, '游客', '2022-12-07 17:06:15', '127.0.0.1:59572', '2022-12-07 11:29:26', '2022-12-07 11:29:26', '');

SET FOREIGN_KEY_CHECKS = 1;
