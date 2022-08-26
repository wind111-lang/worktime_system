create database if not exists employee character set utf8 collate utf8_general_ci;
use employee;

create admin if not exists 'admin'@'localhost' identified by 'root';
grant all privileges on employee.* to 'admin'@'localhost';

create table if not exists employee(
    id bigint not null auto_increment,
    name varchar(255) not null,
    age int not null,
    workingStatus varchar(255) not null,
    worktime time not null,
    salary double not null,
    primary key(id)
)

insert into employee(
    name,
    age,
    workingStatus,
    worktime,
    salary
) values(
    'オーナー',
    30,
    '正社員',
    '0:00:00',
    1500
)
