create table p4l_user
(
    "id" serial primary key,
    "username" varchar(256) not null unique,
    "password_hash" varchar(512) not null
);
