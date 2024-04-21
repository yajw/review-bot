CREATE TABLE `user_review` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `scene_key` varchar(64) NOT NULL DEFAULT '',
    `scene_id` bigint NOT NULL DEFAULT '0',
    `uid` bigint NOT NULL DEFAULT '0',
    `review_content` text,
    `submit_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `extra_attrs` text,
    PRIMARY KEY (`id`),
    KEY `idx_uid_scene` (`uid`, `scene_key`, `scene_id`),
    KEY `idx_uid_submit_time` (`uid`, `submit_time`),
    KEY `idx_create_time` (`create_time`),
    KEY `idx_modify_time` (`modify_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
create table if not EXISTS review_template (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `scene_key` varchar(64) not null default '',
    `tempalte` text,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    key `idx_scene_key` (scene_key)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
create table if NOT EXISTS user_order (
    `id` bigint unsigned not null AUTO_INCREMENT,
    `uid` bigint not null default 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    key `idx_uid` (uid)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;