create table minion
(
    `id`         BIGINT auto_increment comment 'ID',
    `name`       VARCHAR(50)                        not null comment '节点名字',
    `inet`       VARCHAR(15)                        not null comment 'IPv4',
    `inet6`      VARCHAR(39) null comment 'IPv6',
    `status`     TINYINT(1) not null comment '状态',
    `mac`        VARCHAR(17) null comment '网卡 MAC 地址',
    `goos`       VARCHAR(10) null comment '操作系统',
    `arch`       VARCHAR(10) null comment '系统架构',
    `semver`     VARCHAR(20) null comment '版本',
    `cpu`        int      default 0                 not null comment 'CPU 核心数',
    `pid`        int      default 0                 not null comment 'agent 启动进程 PID',
    `username`   VARCHAR(50) null comment 'agent 进程启动用户',
    `hostname`   VARCHAR(50) null comment '主机名',
    `workdir`    TEXT null comment '工作目录',
    `executable` int null comment '执行路径',
    `pinged_at`  DATETIME null comment '最近一次 ping 时间',
    `joined_at`  DATETIME null comment '最近一次连接时间',
    `created_at` DATETIME default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` DATETIME default CURRENT_TIMESTAMP not null comment '更新时间',
    constraint minion_pk
        primary key (`id`),
    constraint minion_pk2
        unique (`inet`)
) comment 'agent 节点信息表';

