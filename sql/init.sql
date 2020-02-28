CREATE TABLE if NOT EXISTS `heroes` (
    `id`            bigint(20)   PRIMARY KEY     NOT NULL ,
    `name`          varchar (255)    NOT NULL DEFAULT '无名氏',
    `create_date`   timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_date`   timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
) engine = InnoDB DEFAULT charset = utf8mb4;