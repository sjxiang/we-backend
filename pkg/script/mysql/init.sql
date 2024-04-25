
CREATE DATABASE IF NOT EXISTS `we_backend`;

USE `we_backend`;

-- 查看建表语句
SHOW CREATE TABLE `user`;



CREATE TABLE IF NOT EXISTS `user` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` varchar(64) DEFAULT '' COMMENT '用户id',
    `password` varchar(64) DEFAULT '' COMMENT '密码',
    `nickname` varchar(64) DEFAULT '' COMMENT '昵称',
    `email` varchar(64) DEFAULT '' COMMENT '邮箱',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`) USING BTREE,
    UNIQUE KEY `uk_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='cms账号信息';


INSERT INTO 
    `account` (`user_id`, `nickname`, `password`, `email`)
VALUES 
    ("ea7ee4ac-4b0b-4f78-846e-26e7ea70411d", "sjxiang1997", "123456789qwe", "1535484943@qq.com");

SELECT EXISTS(
    SELECT true FROM account WHERE user_id = "ea7ee4ac-4b0b-4f78-846e-26e7ea70411d"
    );



CREATE TABLE IF NOT EXISTS `t_content_details` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `content_id` varchar(255) DEFAULT '' COMMENT '内容ID',
    `title` varchar(255) DEFAULT '' COMMENT '内容标题',
    `description` text COMMENT '内容描述',
    `author` varchar(255) DEFAULT '' COMMENT '作者',
    `video_url` varchar(255) DEFAULT '' COMMENT '视频播放URL',
    `thumbnail` varchar(255) DEFAULT '' COMMENT '封面图URL',
    `category` varchar(255) DEFAULT '' COMMENT '内容分类',
    `duration` bigint DEFAULT '0' COMMENT '内容时长',
    `resolution` varchar(255) DEFAULT '' COMMENT '分辨率 如720p、1080p',
    `file_size` bigint DEFAULT '0' COMMENT '文件大小',
    `format` varchar(255) DEFAULT '' COMMENT '文件格式 如MP4、AVI',
    `quality` int DEFAULT '0' COMMENT '视频质量 1-高清 2-标清',
    `approval_status` int DEFAULT '1' COMMENT '审核状态 1-审核中 2-审核通过 3-审核不通过',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '内容创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '内容更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='内容详情';
