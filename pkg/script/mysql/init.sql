
-- 创建数据库
CREATE DATABASE IF NOT EXISTS `we_backend` DEFAULT CHARACTER SET = 'utf8mb4';
    
-- 使用数据库
USE `we_backend`;

-- 查看建表语句
SHOW CREATE TABLE `users`;

-- 删除表
DROP TABLE IF EXISTS `users`;


CREATE TABLE `users` (
    `id`           bigint(20)    NOT NULL AUTO_INCREMENT            COMMENT '自增id',
    `created_at`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `updated_at`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `email`        varchar(64)   NOT NULL DEFAULT ''                COMMENT '邮件',
    `mobile`       varchar(64)   NOT NULL DEFAULT ''                COMMENT '手机号',
    `nickname`     varchar(64)   NOT NULL DEFAULT ''                COMMENT '昵称',
    `password`     varchar(64)   NOT NULL DEFAULT ''                COMMENT '密码',
    `intro`        varchar(1024) NOT NULL DEFAULT ''                COMMENT '自我介绍',
    `avatar`       varchar(1024) NOT NULL DEFAULT ''                COMMENT '头像',
    `birthday`     datetime(3)            DEFAULT NULL              COMMENT '生日',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uk_email`(`email`) USING BTREE,
    UNIQUE INDEX `uk_mobile`(`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';


INSERT INTO 
    `users` (`id`, `nickname`, `password`, `email`, `mobile`)
VALUES 
    (24, "sjxiang1997", "123456789qwe", "1535484943@qq.com", "18812347777");

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
