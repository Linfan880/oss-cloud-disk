SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
                       `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                       `identity` varchar(36) DEFAULT NULL,
                       `hash` varchar(32) DEFAULT NULL COMMENT '文件的唯一标识',
                       `name` varchar(255) DEFAULT NULL,
                       `ext` varchar(30) DEFAULT NULL COMMENT '文件扩展名',
                       `size` int(11) DEFAULT NULL COMMENT '文件大小',
                       `path` varchar(255) DEFAULT NULL COMMENT '文件路径',
                       `is_shared` TINYINT(1) DEFAULT 0 COMMENT '是否为共享文件，0表示不共享(默认)，1表示共享',
                       `created_at` datetime DEFAULT NULL,
                       `updated_at` datetime DEFAULT NULL,
                       `deleted_at` datetime DEFAULT NULL,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for share_pool
-- ----------------------------
DROP TABLE IF EXISTS `shared_pool`;
CREATE TABLE `shared_pool` (
                           `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                           `user_identity` varchar(255) NOT NULL,
                           `user_name` varchar(32) DEFAULT NULL,
                           `resource_name` varchar(32) DEFAULT NULL,
                           `resource_identity` varchar(255) DEFAULT NULL,
                           `ext` varchar(10) DEFAULT NULL COMMENT '文件扩展名',
                           `size` int(100) DEFAULT NULL COMMENT '文件大小',
                           `path` varchar(10) DEFAULT NULL COMMENT '文件路径',
                           `click_num` int(11) DEFAULT '0' COMMENT '点击次数',
                           `created_at` datetime DEFAULT NULL,
                           `updated_at` datetime DEFAULT NULL,
                           `deleted_at` datetime DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                      `identity` varchar(255) DEFAULT NULL COMMENT '用户唯一标识',
                      `name` varchar(60) DEFAULT NULL COMMENT '用户名',
                      `password_hash` varchar(32) DEFAULT NULL COMMENT '加密后的密码',
                      `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
                      `created_at` datetime DEFAULT NULL,
                      `updated_at` datetime DEFAULT NULL,
                      `deleted_at` datetime DEFAULT NULL,
                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_repository
-- ----------------------------
DROP TABLE IF EXISTS `user_repository`;
CREATE TABLE `user_repository` (
                               `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                               `user_identity` varchar(255) DEFAULT NULL,
                               `user_name` varchar(60) DEFAULT NULL,
                               `resource_name` int(255) DEFAULT NULL,
                               `ext` varchar(30) DEFAULT NULL COMMENT '文件扩展名',
                               `size` int(11) DEFAULT NULL COMMENT '文件大小',
                               `path` varchar(255) DEFAULT NULL COMMENT '文件路径',
                               `is_shared` TINYINT(1) DEFAULT 0 COMMENT '是否为共享文件，0表示不共享(默认)，1表示共享',
                               `created_at` datetime DEFAULT NULL,
                               `updated_at` datetime DEFAULT NULL,
                               `deleted_at` datetime DEFAULT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;