
CREATE TABLE `user`(
    `id` bigint (20) NOT NULL AUTO_INCREMENT, #不是自增ID是为了防止其他人拿到ID知道有多少用户
    `user_id` bigint (20) NOT NULL ,
    `username` varchar (64) COLLATE utf8mb4_general_ci not null ,
    `password` varchar (64) COLLATE utf8mb4_general_ci not null ,
    `email` varchar (64) COLLATE utf8mb4_general_ci,
    `gender` tinyint (4) NOT NULL DEFAULT 0,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_username` (`username`) using BTREE ,
    UNIQUE KEY `index_user_id` (`user_id`) using BTREE
)ENGINE=InnoDB DEFAULT CHARSET = utf8mb4 COLLATE =utf8mb4_general_ci;


DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10) unsigned NOT NULL,
    `community_name` varchar(128) collate utf8mb4_general_ci not null ,
    `introduction` varchar(256) collate utf8mb4_general_ci not null ,
    `create_time` timestamp not null default current_timestamp,
    `update_time` timestamp not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
)ENGINE=InnoDB DEFAULT CHARSET = utf8mb4 COLLATE =utf8mb4_general_ci;

INSERT INTO `community` values ('1','1','Go','Golang','2016-11-01 08:10:10','2016-11-01 08:10:10');
INSERT INTO `community` values ('2','2','leetcode','算法','2020-01-01 08:10:10','2020-01-01 08:10:10');
INSERT INTO `community` values ('3','3','CS:GO','RUSH B','2018-08-01 08:10:10','2018-08-01 08:10:10');
INSERT INTO `community` values ('4','4','DOTA2','welcome to dota','2016-08-01 08:10:10','2016-08-01 08:10:10');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`(
    `id` bigint(20) not null AUTO_INCREMENT,
    `post_id` bigint(20) not null comment '帖子id',
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题',
    `content` varchar(8129) collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null comment '作者用户id',
    `community_id` bigint(20) not null comment '所属社区',
    `status` tinyint(4) not null default '1' comment '帖子状态',
    `create_time` timestamp not null default current_timestamp comment '创建时间',
    `update_time` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
)ENGINE=InnoDB DEFAULT CHARSET = utf8mb4 COLLATE =utf8mb4_general_ci;