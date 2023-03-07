/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.8
Source Server Version : 50736
Source Host           : 192.168.1.8:3306
Source Database       : up-admin

Target Server Type    : MYSQL
Target Server Version : 50736
File Encoding         : 65001

Date: 2023-03-07 21:00:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for function_basic
-- ----------------------------
DROP TABLE IF EXISTS `function_basic`;
CREATE TABLE `function_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `uri` varchar(255) DEFAULT NULL,
  `sort` int(11) DEFAULT '0',
  `menu_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_function_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of function_basic
-- ----------------------------
INSERT INTO `function_basic` VALUES ('1', null, null, null, '1', 'test', '/test', '0', '1');
INSERT INTO `function_basic` VALUES ('2', '2023-02-03 21:32:51.915', '2023-02-03 21:39:59.816', '2023-02-03 21:41:54.149', '1cde2a40-0b82-42e7-bc30-acf469f47c86', '修改功能测试', '/test2', '10', '1');

-- ----------------------------
-- Table structure for menu_basic
-- ----------------------------
DROP TABLE IF EXISTS `menu_basic`;
CREATE TABLE `menu_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `identity` varchar(36) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `sort` int(11) DEFAULT '0',
  `web_icon` varchar(100) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `level` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_menu_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu_basic
-- ----------------------------
INSERT INTO `menu_basic` VALUES ('1', null, null, null, '1', '0', '首页', '0', 'HomeFilled', '/home', '1');
INSERT INTO `menu_basic` VALUES ('2', null, '2023-03-05 20:42:41.909', null, '2', '0', '设置', '0', 'Setting', '/setting', '0');
INSERT INTO `menu_basic` VALUES ('3', null, null, null, '3', '2', '角色管理', '0', null, '/setting/role', '1');
INSERT INTO `menu_basic` VALUES ('4', null, '2023-03-05 19:03:38.691', null, '4', '2', '菜单管理', '50', '', '/setting/menu', '1');
INSERT INTO `menu_basic` VALUES ('5', '2023-02-02 21:40:52.666', '2023-02-02 21:52:07.558', '2023-02-02 21:57:21.319', '11435198-e662-40e7-987e-ab350dd4966e', '1', '修改菜单测试', '0', null, null, '0');
INSERT INTO `menu_basic` VALUES ('6', '2023-02-02 21:41:28.243', '2023-02-02 21:41:28.243', '2023-02-02 21:57:21.319', 'a22472f9-c1e6-4ed8-a6f1-07f72d5cf6cc', '0', '新增菜单测试', '0', 'Select', null, '0');
INSERT INTO `menu_basic` VALUES ('7', '2023-02-02 21:42:20.668', '2023-02-02 21:42:20.668', '2023-02-02 21:57:21.319', '74189e41-4d77-425a-8017-f3deafbef233', '2', '新增菜单测试', '0', null, null, '0');
INSERT INTO `menu_basic` VALUES ('8', '2023-03-05 15:55:34.591', '2023-03-05 19:03:18.842', '2023-03-05 19:07:54.819', '098c30d5-9154-41f7-b5b4-31d77f4d788f', '0', 'top3', '0', 'Search', '/path', '0');
INSERT INTO `menu_basic` VALUES ('9', '2023-03-05 16:17:08.379', '2023-03-05 19:02:24.270', '2023-03-05 19:07:56.981', 'b7f71057-22ec-4cb6-8a99-bc3d760dc3b5', '0', 'top2', '100', 'Search', '/top', '0');
INSERT INTO `menu_basic` VALUES ('10', '2023-03-05 16:23:35.611', '2023-03-05 18:49:03.459', '2023-03-05 19:07:58.963', '2aa33033-27d1-4faa-90be-22683c123ddc', '0', 'top4', '1000', 'Search2', '/top2', '0');
INSERT INTO `menu_basic` VALUES ('11', '2023-03-05 18:47:42.482', '2023-03-05 18:50:44.520', '2023-03-05 19:07:44.578', 'f8910048-6e36-4a8b-8e02-fa40c26615d8', '0', 'top1', '1', 'Search1', '/top3', '0');
INSERT INTO `menu_basic` VALUES ('12', '2023-03-05 19:06:45.841', '2023-03-05 19:06:45.841', '2023-03-05 19:07:52.026', '2148c1b9-f567-4414-a210-155f7ef2601f', '8', 'top3-子菜单', '0', '', '/top3/sub', '0');
INSERT INTO `menu_basic` VALUES ('13', '2023-03-05 20:09:07.980', '2023-03-05 20:42:48.908', '2023-03-05 23:14:03.960', '8597e2b3-1994-4fc9-802f-f1535714dd4d', '0', '一号菜单', '100', 'Search', '/one', '1');
INSERT INTO `menu_basic` VALUES ('14', '2023-03-05 20:15:18.005', '2023-03-05 20:15:18.005', '2023-03-05 23:14:01.807', 'cc10e495-12b2-4af2-9bb4-4d01ebf9959d', '0', '二号菜单', '20', 'Search', '/two', '1');
INSERT INTO `menu_basic` VALUES ('15', '2023-03-05 20:17:46.095', '2023-03-05 20:41:53.471', null, '895132bd-106e-4705-9515-5f2bf677cb68', '13', '一号-一级', '0', '', '/one/one', '2');
INSERT INTO `menu_basic` VALUES ('16', '2023-03-05 20:17:57.982', '2023-03-05 20:17:57.982', '2023-03-05 20:39:57.710', 'a14f6c24-e94d-4c05-8bd4-3c308024f02d', '15', '一号-二级', '0', '', '', '0');
INSERT INTO `menu_basic` VALUES ('17', '2023-03-05 20:19:11.649', '2023-03-05 20:21:14.574', '2023-03-05 20:39:54.875', 'c8839ca6-d03d-4f27-bb61-a46e9677a765', '16', '一号-三级-按钮', '0', '', '', '2');
INSERT INTO `menu_basic` VALUES ('18', '2023-03-05 23:42:26.523', '2023-03-05 23:46:29.207', '2023-03-05 23:55:33.417', '251603d1-6e17-46d5-bd38-ddb049286baf', '0', 'top', '0', 'Top', '/top', '1');
INSERT INTO `menu_basic` VALUES ('19', '2023-03-05 23:50:52.562', '2023-03-05 23:51:46.214', null, 'cee73782-2072-41b0-b739-d0c2afeb8d9a', '18', 'top2', '0', '', '/top/top2', '1');
INSERT INTO `menu_basic` VALUES ('20', '2023-03-06 00:19:05.137', '2023-03-06 00:19:05.137', '2023-03-06 21:14:35.426', 'b3e52aa5-0f61-407d-9fbb-8891cdc0ff48', '0', 'top', '0', 'Top', '/top', '0');
INSERT INTO `menu_basic` VALUES ('21', '2023-03-06 00:19:20.323', '2023-03-06 00:19:20.323', '2023-03-06 21:14:33.645', '536d25db-30c3-4acd-b3e8-f3528a8a180f', '20', 'top2', '0', '', '/top/top2', '1');
INSERT INTO `menu_basic` VALUES ('22', '2023-03-06 21:14:58.878', '2023-03-06 21:14:58.878', null, '010938c8-0b0b-4cc1-89a4-3504a3bae4cf', '2', '管理员管理', '0', '', '/setting/user', '1');

-- ----------------------------
-- Table structure for role_basic
-- ----------------------------
DROP TABLE IF EXISTS `role_basic`;
CREATE TABLE `role_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT '0',
  `sort` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_role_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_basic
-- ----------------------------
INSERT INTO `role_basic` VALUES ('1', '2023-01-14 14:36:36.297', '2023-01-28 22:17:47.555', null, '1', '超级管理员', '1', '0');
INSERT INTO `role_basic` VALUES ('2', '2023-01-14 14:36:36.297', '2023-03-06 00:18:29.180', null, 'd1d56591-55db-484e-96a6-d94a5a833cd9', '修改角色测试', '0', '0');
INSERT INTO `role_basic` VALUES ('3', '2023-03-05 22:33:41.573', '2023-03-05 23:12:50.859', null, '2b4ab912-d0c0-45b1-8d5e-dbc3dfb2623f', '修改测试', '0', '0');
INSERT INTO `role_basic` VALUES ('4', '2023-03-05 22:34:28.601', '2023-03-05 22:34:35.741', '2023-03-05 22:34:39.813', '7133084e-9a22-4291-9c14-379281e01ebf', '超管2', '0', '0');

-- ----------------------------
-- Table structure for role_function
-- ----------------------------
DROP TABLE IF EXISTS `role_function`;
CREATE TABLE `role_function` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `function_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_role_function_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_function
-- ----------------------------
INSERT INTO `role_function` VALUES ('1', null, null, null, '1', '1');
INSERT INTO `role_function` VALUES ('2', '2023-01-14 14:36:36.320', '2023-01-14 14:36:36.320', '2023-01-27 21:19:57.182', '2', '1');
INSERT INTO `role_function` VALUES ('3', '2023-01-27 21:19:57.187', '2023-01-27 21:19:57.187', '2023-01-28 22:17:29.078', '2', '1');
INSERT INTO `role_function` VALUES ('4', '2023-01-28 22:17:29.080', '2023-01-28 22:17:29.080', '2023-01-28 22:17:47.611', '2', '1');
INSERT INTO `role_function` VALUES ('5', '2023-01-28 22:17:47.613', '2023-01-28 22:17:47.613', '2023-01-28 22:18:13.898', '2', '1');
INSERT INTO `role_function` VALUES ('6', '2023-01-28 22:18:13.901', '2023-01-28 22:18:13.901', null, '2', '1');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `menu_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_role_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES ('1', null, null, null, '1', '1');
INSERT INTO `role_menu` VALUES ('2', null, null, null, '1', '2');
INSERT INTO `role_menu` VALUES ('3', null, null, null, '1', '3');
INSERT INTO `role_menu` VALUES ('4', '2023-01-14 14:36:36.319', '2023-01-14 14:36:36.319', '2023-01-27 21:19:57.180', '2', '1');
INSERT INTO `role_menu` VALUES ('5', '2023-01-14 14:36:36.319', '2023-01-14 14:36:36.319', '2023-01-27 21:19:57.180', '2', '2');
INSERT INTO `role_menu` VALUES ('6', '2023-01-27 21:19:57.186', '2023-01-27 21:19:57.186', '2023-01-28 22:17:29.076', '2', '1');
INSERT INTO `role_menu` VALUES ('7', '2023-01-27 21:19:57.186', '2023-01-27 21:19:57.186', '2023-01-28 22:17:29.076', '2', '2');
INSERT INTO `role_menu` VALUES ('8', '2023-01-28 22:17:29.079', '2023-01-28 22:17:29.079', '2023-01-28 22:17:47.610', '2', '1');
INSERT INTO `role_menu` VALUES ('9', '2023-01-28 22:17:29.079', '2023-01-28 22:17:29.079', '2023-01-28 22:17:47.610', '2', '2');
INSERT INTO `role_menu` VALUES ('10', '2023-01-28 22:17:47.612', '2023-01-28 22:17:47.612', '2023-01-28 22:18:13.897', '2', '1');
INSERT INTO `role_menu` VALUES ('11', '2023-01-28 22:17:47.612', '2023-01-28 22:17:47.612', '2023-01-28 22:18:13.897', '2', '2');
INSERT INTO `role_menu` VALUES ('12', '2023-01-28 22:18:13.899', '2023-01-28 22:18:13.899', '2023-03-06 00:14:39.854', '2', '1');
INSERT INTO `role_menu` VALUES ('13', '2023-01-28 22:18:13.899', '2023-01-28 22:18:13.899', '2023-03-06 00:14:39.854', '2', '2');
INSERT INTO `role_menu` VALUES ('14', '2023-03-05 22:33:41.575', '2023-03-05 22:33:41.575', '2023-03-05 23:12:50.861', '3', '1');
INSERT INTO `role_menu` VALUES ('15', '2023-03-05 22:33:41.575', '2023-03-05 22:33:41.575', '2023-03-05 23:12:50.861', '3', '2');
INSERT INTO `role_menu` VALUES ('16', '2023-03-05 22:33:41.575', '2023-03-05 22:33:41.575', '2023-03-05 23:12:50.861', '3', '3');
INSERT INTO `role_menu` VALUES ('17', '2023-03-05 22:33:41.575', '2023-03-05 22:33:41.575', '2023-03-05 23:12:50.861', '3', '4');
INSERT INTO `role_menu` VALUES ('18', '2023-03-05 23:12:50.862', '2023-03-05 23:12:50.862', null, '3', '1');
INSERT INTO `role_menu` VALUES ('19', '2023-03-05 23:12:50.862', '2023-03-05 23:12:50.862', null, '3', '2');
INSERT INTO `role_menu` VALUES ('20', '2023-03-05 23:12:50.862', '2023-03-05 23:12:50.862', null, '3', '3');
INSERT INTO `role_menu` VALUES ('21', '2023-03-05 23:12:50.862', '2023-03-05 23:12:50.862', null, '3', '4');
INSERT INTO `role_menu` VALUES ('22', '2023-03-05 23:12:50.862', '2023-03-05 23:12:50.862', null, '3', '14');
INSERT INTO `role_menu` VALUES ('23', '2023-03-06 00:14:39.855', '2023-03-06 00:14:39.855', '2023-03-06 00:17:26.949', '2', '1');
INSERT INTO `role_menu` VALUES ('24', '2023-03-06 00:14:39.855', '2023-03-06 00:14:39.855', '2023-03-06 00:17:26.949', '2', '2');
INSERT INTO `role_menu` VALUES ('25', '2023-03-06 00:14:39.855', '2023-03-06 00:14:39.855', '2023-03-06 00:17:26.949', '2', '3');
INSERT INTO `role_menu` VALUES ('26', '2023-03-06 00:17:26.951', '2023-03-06 00:17:26.951', '2023-03-06 00:18:29.182', '2', '1');
INSERT INTO `role_menu` VALUES ('27', '2023-03-06 00:17:26.951', '2023-03-06 00:17:26.951', '2023-03-06 00:18:29.182', '2', '2');
INSERT INTO `role_menu` VALUES ('28', '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', null, '2', '1');
INSERT INTO `role_menu` VALUES ('29', '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', null, '2', '2');
INSERT INTO `role_menu` VALUES ('30', '2023-03-06 00:18:29.184', '2023-03-06 00:18:29.184', null, '2', '3');

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `identity` varchar(36) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(36) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `wx_union_id` varchar(255) DEFAULT NULL,
  `wx_open_id` varchar(255) DEFAULT NULL,
  `role_identity` varchar(36) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES ('1', '2023-02-21 22:17:59.365', '2023-02-21 22:17:59.365', null, '1', 'get', 'e10adc3949ba59abbe56e057f20f883e', '13333333333', null, null, '1', null);
INSERT INTO `user_basic` VALUES ('2', '2023-02-21 22:17:59.365', '2023-02-21 22:17:59.365', null, '2', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '13344445555', null, null, 'd1d56591-55db-484e-96a6-d94a5a833cd9', null);
INSERT INTO `user_basic` VALUES ('5', '2023-03-06 23:05:51.427', '2023-03-06 23:05:51.427', null, '4285b6fa-0559-4274-8b19-4c78e4d04df7', 'get1', 'e10adc3949ba59abbe56e057f20f883e', '', '', '', '1', '');
