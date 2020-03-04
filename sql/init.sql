CREATE TABLE IF NOT EXISTS `heroes` (
    `id`            bigint(20)   PRIMARY KEY     NOT NULL AUTO_INCREMENT,
    `name`          varchar (255)    NOT NULL DEFAULT '无名氏',
    `create_date`   timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_date`   timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
) engine = InnoDB DEFAULT charset = utf8mb4;


CREATE TABLE IF NOT EXISTS `blog_pages` (
    `id`            bigint(20)      PRIMARY KEY     NOT NULL    AUTO_INCREMENT,
    `title`         varchar(255)    NOT NULL    COMMENT '标题',
    `update_time`   timestamp       NOT NULL    DEFAULT CURRENT_TIMESTAMP   ON UPDATE CURRENT_TIMESTAMP   COMMENT '更新时间',
    `create_time`   timestamp       NOT NULL    DEFAULT CURRENT_TIMESTAMP   COMMENT '生成时间',
    `introduce`     mediumtext      NOT NULL    COMMENT '正文',
    `img_url`       tinytext        DEFAULT NULL    COMMENT '文章图片'
) engine = InnoDB DEFAULT charset = utf8mb4;

CREATE TABLE IF NOT EXISTS `blog_tags` (
    `id`            bigint(20)      PRIMARY KEY     NOT NULL    AUTO_INCREMENT,
    `tag`           varchar(255)    NOT NULL    COMMENT '标签名'
) engine = InnoDB DEFAULT charset = utf8mb4;

CREATE TABLE IF NOT EXISTS `blog_page_to_tags` (
    `id`            bigint(20)      PRIMARY KEY     NOT NULL    AUTO_INCREMENT,
    `page_id`       bigint(20)      NOT NULL    COMMENT '文章ID',
    `tag_id`        bigint(20)      NOT NULL    COMMENT '标签ID'
) engine = InnoDB DEFAULT charset = utf8mb4;