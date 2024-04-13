drop table if exists `course`;
create table `course` (
                          `id` bigint not null comment 'id',
                          `name` varchar(50) comment '名称',
                          `level` varchar(10) comment '等级',
                          `price` decimal(6,2) comment '金额（元）',
                          `desc` varchar(200) comment '描述',
                          primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='课程';

insert into course values (1, '支付系统1', '高级', 0.01, '这是一个企业级支付平台项目');
insert into course values (2, '支付系统2', '高级', 0.01, '这是一个企业级支付平台项目');
insert into course values (3, '支付系统3', '高级', 0.01, '这是一个企业级支付平台项目');
