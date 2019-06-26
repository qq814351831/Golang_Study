create database driving_exam charset=utf8;
use driving_exam;
create table score(
 id int primary key auto_increment,
 name varchar(10) not null,
 score int
);