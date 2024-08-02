/*
 Navicat Premium Data Transfer

 Source Server         : local.8
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : 127.0.0.1:3306
 Source Schema         : jenny

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 02/08/2024 17:49:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `is_draft` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '草稿: 1=是, 2=否,',
  `draft_from_id` int(11) NOT NULL DEFAULT 0 COMMENT '草稿来源ID: 0表示需要新建',
  `draft_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 1=已同步, 2=审核中, 3=审核拒绝, 4=审核通过待同步',
  `small_title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '简略标题',
  `small_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '简略图',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标题',
  `background_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '背景图',
  `content_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内容类型: 1=txt, 2=html, 3=md,',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '内容',
  `release_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `last_modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次编辑时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=正常, 2=禁用, 3=下架',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_release_time`(`release_time` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '文章' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT 'PId',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'code',
  `is_menu` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是菜单：1=是, 2=否',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=正常, 2=禁用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pid_code`(`pid` ASC, `code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category_item_relation
-- ----------------------------
DROP TABLE IF EXISTS `category_item_relation`;
CREATE TABLE `category_item_relation`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类ID: category.id',
  `item_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '内容类型: 1=游戏，2=文章，',
  `item_id` int(11) NOT NULL DEFAULT 0 COMMENT '内容ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `category_item_unique`(`category_id` ASC, `item_type` ASC, `item_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '分类-内容-关联关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标题',
  `list_small_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '列表缩略图',
  `small_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '详情缩略图',
  `origin_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '原图',
  `size` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '原图大小 byte',
  `weight` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '原图宽 像素',
  `high` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '原图高 像素',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '原图类型',
  `is_free` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '免费：1=是，2=否，3=未知',
  `keyword` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '关键字(多个逗号分隔)',
  `source` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源：1=游戏，2=movie，3=后台上传，4=ai生成',
  `source_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源ID',
  `source_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '来源code',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=正常，2=审核中，3=审核拒绝，4=禁用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_source_id`(`source_id` ASC) USING BTREE,
  INDEX `source_code`(`source_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '图片' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for personal_access_tokens
-- ----------------------------
DROP TABLE IF EXISTS `personal_access_tokens`;
CREATE TABLE `personal_access_tokens`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `personal_access_tokens_token_unique`(`token` ASC) USING BTREE,
  INDEX `personal_access_tokens_tokenable_type_tokenable_id_index`(`tokenable_type` ASC, `tokenable_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for steam_game
-- ----------------------------
DROP TABLE IF EXISTS `steam_game`;
CREATE TABLE `steam_game`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cn_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '中文名称',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `steam_appid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'steam_appid',
  `cc` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'cc区',
  `l` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '语言',
  `header_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '缩略图',
  `capsule_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '列表缩略图',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面图',
  `background_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '背景图',
  `short_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '一句话简介',
  `about_the_game` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '关于游戏',
  `detail_description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '简介',
  `required_age` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄限制',
  `is_violent` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '包含暴力、血腥内容：1=是，2=否，3=未知',
  `is_sexy` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '包含色情内容：1=是，2=否，3=未知',
  `supported_languages` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '支持的语言',
  `support_chinese` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支持中文：1=是，2=否，3=未知',
  `is_free` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '免费：1=是，2=否，3=未知',
  `have_dic` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有dlc：1=是，2=否，3=未知',
  `dic_ids` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'dlc_ids',
  `price_currency` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '价格-货币',
  `price_initial` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '价格-原价（分）',
  `price_final` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '价格-现价（分）',
  `price_discount_percent` decimal(8, 2) NOT NULL DEFAULT 0.00 COMMENT '价格-折扣',
  `price_unit` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '价格-单位',
  `windows` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支持windows：1=是，2=否，3=未知',
  `mac` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支持mac：1=是，2=否，3=未知',
  `linux` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支持linux：1=是，2=否，3=未知',
  `network` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '联网：1=是，2=否，3=未知',
  `images` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有图片：1=是，2=否，3=未知',
  `movies` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有视频：1=是，2=否，3=未知',
  `genre` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类型(多个逗号分隔)',
  `developers` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开发商(多个逗号分隔)',
  `publishers` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '发行商(多个逗号分隔)',
  `content_descriptors_notes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '提醒文案',
  `coming_soon` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已发布：1=是，2=否，3=未知',
  `release_date` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '发布时间',
  `release_time` int(11) NOT NULL DEFAULT 0 COMMENT '发布时间戳',
  `last_modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次更新时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=正常，2=审核中，3=审核拒绝，4=禁用，5=删除',
  `used` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=已使用，2=未使用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_steam_appid`(`steam_appid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16441 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'steam游戏' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for steam_game_image
-- ----------------------------
DROP TABLE IF EXISTS `steam_game_image`;
CREATE TABLE `steam_game_image`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `steam_game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'steam游戏ID',
  `path_thumbnail` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '缩略图',
  `path_full` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '原图',
  `image_code` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片标记',
  `used` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=已使用，2=未使用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_steam_game_id`(`steam_game_id` ASC) USING BTREE,
  INDEX `idx_image_code`(`image_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 143994 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'steam游戏-图片' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for steam_game_video
-- ----------------------------
DROP TABLE IF EXISTS `steam_game_video`;
CREATE TABLE `steam_game_video`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `steam_game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'steam游戏ID',
  `origin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'steam游戏movieID',
  `thumbnail` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '缩略图',
  `webm_480` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'webm 480 视频地址',
  `webm_max` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'webm max 视频地址',
  `mp4_480` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'mp4 480 视频地址',
  `mp4_max` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'mp4 max 视频地址',
  `used` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1=已使用，2=未使用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_steam_game_id`(`steam_game_id` ASC) USING BTREE,
  INDEX `idx_origin_id`(`origin_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22711 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'steam游戏-视频' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `email_verified_at` timestamp NULL DEFAULT NULL COMMENT '邮箱验证时间',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态:1=正常,2=待激活,3=禁止登录,4=禁用,5=注销,',
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `users_email_unique`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
