create
    database if not exists axapi;

use axapi;

create table if not exists user
(
    id              bigint auto_increment comment 'id' primary key,
    account         varchar(256)                          not null comment '账号',
    username        varchar(256)                          null comment '用户名',
    password        varchar(512)                          null comment '密码',
    avatar          varchar(1024)                         null comment '用户头像',
    email           varchar(128)                          null comment '邮箱',
    phone           varchar(128)                          null comment '电话',
    profile         varchar(2048)                         null comment '用户简介',
    gender          tinyint     default 0                 not null comment '男性:0, 女性:1',
    role            varchar(16) default 'user'            not null comment '用户角色:user, admin...',
    status          tinyint     default 0                 not null comment '账号状态(0- 正常 1- 封号)',
    access_key      varchar(256)                          null comment 'access_key',
    secret_key      varchar(256)                          null comment 'secret_ey',
    balance         bigint      default 30                not null comment '钱包余额,注册送30币',
    invitation_code varchar(256)                          null comment '邀请码',
    create_time     datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time     datetime    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete       tinyint     default 0                 not null comment '是否删除',
    index idx_user_id (id),
    index idx_user_account (account),
    index idx_user_username (username)
) comment '用户表' collate = utf8mb4_unicode_ci
                   engine = InnoDB;

create table if not exists interface_info
(
    id              bigint auto_increment comment 'id' primary key,
    name            varchar(256)                           not null comment '接口名称',
    url             varchar(256)                           not null comment '接口地址',
    user_id         bigint       default 0                 null comment '发布人',
    method          varchar(256)                           not null comment '请求方法',
    request_params  text                                   null comment '接口请求参数',
    response_params text                                   null comment '接口响应参数',
    reduce_score    bigint       default 0                 null comment '扣除积分数',
    request_example text                                   null comment '请求示例',
    request_header  text                                   null comment '请求头',
    response_header text                                   null comment '响应头',
    return_Format   varchar(512) default 'JSON'            null comment '返回格式(JSON等等)',
    description     text                                   null comment '描述信息',
    status          tinyint      default 0                 not null comment '接口状态(0- 默认下线 1- 上线)',
    totalInvokes    bigint       default 0                 not null comment '接口总调用次数',
    avatarUrl       varchar(1024)                          null comment '接口头像',
    create_time     datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time     datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete       tinyint      default 0                 not null comment '是否删除',
    index idx_interface_info_id (id),
    index idx_interface_info_name (name)
) comment '接口信息表' collate = utf8mb4_unicode_ci
                       engine = InnoDB;