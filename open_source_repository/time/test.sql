drop table if exists order;
create table order
(
    id                  bigint primary key not null auto_increment,
    space_id            varchar(16) comment '空间id',
    bill_type           bigint      not null default 1 comment '计费方式,1=包年包月',
    order_duration           bigint      not null default 1 comment '购买时长,1=1个月,2=3个月,3=1年,4=自定义',
    product_type  bigint  not null default 1 comment '订单类型,1=计算,2=存储',
    order_amount  bigint  not null default 0 comment '订单金额',
    pay_type bigint not null default 1 comment '付费类型, 1=预付费, 2=后付费',
    pay_status bigint not null default 1 comment '支付状态,1=已支付,2=未支付,3=支付中',
    pay_method bigint not null default 1 comment '支付方式,1=对公打款',
    pay_time bigint not null default 0 comment '支付日期',
    create_user_id bigint             not null default 0,
    create_time         datetime(6)        not null default current_timestamp(6),
    update_user_id      bigint             not null default 0,
    update_time         datetime(6)        not null default current_timestamp(6) on update current_timestamp(6),
) charset = utf8 comment '订单';

drop table if exists benefit;
create table benefit
(
    id                  bigint primary key not null auto_increment,
    space_id            varchar(16) comment '空间id',
    product_type  bigint  not null default 1 comment '产品类型,1=计算,2=存储',
    total_points bigint not null default 0 comment '总点数',
    total_storage bigint not null default 0 comment '存储总容量',
    used_points bigint not null default 0 comment '已使用点数',
    left_points bigint not null default 0 comment '剩余可使用点数'
    benefit_start_time  bigint not null default 0 comment '权益开始时间',
    benefit_end_time  bigint not null default 0 comment '权益结束时间',
    benefit_status  bigint not null default 1 comment '权益状态,1=待生效,2=生效中,3=已结束,4=已取消',
    create_user_id bigint             not null default 0,
    create_time         datetime(6)        not null default current_timestamp(6),
    update_user_id      bigint             not null default 0,
    update_time         datetime(6)        not null default current_timestamp(6) on update current_timestamp(6),
) charset = utf8 comment '权益';

drop table if exists order_benefit_relation;
create table order_benefit_relation
(
    id                  bigint primary key not null auto_increment,
    order_id bigint not null default 0 comment '订单id',
    benefit_id bigint not null default 0 comment '权益id',
    create_user_id bigint             not null default 0,
    create_time         datetime(6)        not null default current_timestamp(6),
    update_user_id      bigint             not null default 0,
    update_time         datetime(6)        not null default current_timestamp(6) on update current_timestamp(6),
) charset = utf8 comment '订单权益关联表';

drop table if exists benefit_spec;
create table benefit_spec
(
    id                  bigint primary key not null auto_increment,
    benefit_id bigint not null default 0 comment '权益id',
    spec_instance_id bigint not null default 0 comment '规格id',
    spec_name varchar(128) not null default '' comment '实例规格名',
    spec_points bigint not null default 0 comment '实例点数',
    create_user_id bigint             not null default 0,
    create_time         datetime(6)        not null default current_timestamp(6),
    update_user_id      bigint             not null default 0,
    update_time         datetime(6)        not null default current_timestamp(6) on update current_timestamp(6),
) charset = utf8 comment '权益规格表';

drop table if exists benefit_job;
create table benefit_job
(
    id             bigint primary key not null auto_increment,
    space_id       varchar(16) comment '空间id',
    job_type       bigint comment '任务类型, 1.训练任务,2.开发环境,3.可视化实例,4.持久化存储',
    job_id         varchar(64) comment '任务id',
    project_id     varchar(64) comment '项目id',
    start_time     datetime(6) null comment '开始时间',
    end_time       datetime(6) null comment '停止时间',
    job_status     bigint comment '状态:1.运行(占用点数),2.停止(不占用点数)',
    total_price    bigint comment '总单价',
    last_bill_time datetime(6) comment '上次计费时间',

    create_user_id bigint             not null default 0,
    create_time         datetime(6)        not null default current_timestamp(6),
    update_user_id      bigint             not null default 0,
    update_time         datetime(6)        not null default current_timestamp(6) on update current_timestamp(6),
    unique index1 (job_id)
) charset = utf8 comment '(使用权益的)任务表';

drop table if exists benefit_job_detail;
create table benefit_job_detail
(
    id                      bigint primary key auto_increment,





    bill_job_id             bigint comment '账单任务id',
    job_id                  varchar(64) comment '任务id',
    space_id                varchar(16) comment '空间id',
    resource_pkg_history_id bigint comment '资源包id',
    pkg_type                bigint comment '资源包类型, 1.计算实例,2.持久化存储',
    price                   bigint comment '单价',
    pkg_count               bigint comment '资源包数量',
    start_time              datetime(6) comment '开始时间',
    end_time                datetime(6) comment '停止时间',
    exp_field               json comment '扩展字段',
    create_user_id          bigint      not null,
    create_time             datetime(6) not null default current_timestamp(6),
    update_user_id          bigint      not null,
    update_time             datetime(6) null     default current_timestamp(6) on update current_timestamp(6)
) charset = utf8mb4 comment '账单任务详情';
