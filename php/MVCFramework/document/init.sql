create table customer
(
    id           int auto_increment,
    cname        varchar(32)                        not null,
    sex          tinyint                            not null,
    birthday     date                               not null,
    address      varchar(32)                        not null,
    email        varchar(32)                        not null,
    mobile       char(11)                           not null,
    created_time datetime default current_timestamp not null,
    updated_time datetime default current_timestamp not null on update current_timestamp,
    constraint customer_pk
        primary key (id)
);

