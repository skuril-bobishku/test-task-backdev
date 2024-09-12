CREATE TABLE sessions
(
s_id serial primary key,
refresh_crypt varchar(255) not null unique,
exp_time timestamp not null
);

CREATE TABLE users
(
gu_id serial primary key,
username varchar(255) not null,
email varchar(255) not null unique,
ipadd varchar(255),
login varchar(255) not null unique,
pass varchar(255) not null,
session_id int,
foreign key (session_id) references sessions(s_id) on delete cascade
);