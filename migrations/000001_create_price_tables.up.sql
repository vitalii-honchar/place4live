create table city
(
    "id" serial primary key,
    "name" varchar(256) not null unique
);

create table city_price
(
    "id"          serial primary key,
    "city_id"     int not null,
    "category"    varchar(256) not null,
    "name"        varchar(256) not null,
    "avg"         float not null,
    "min"         float not null,
    "max"         float not null,
    "updated_at"  timestamp not null,
    constraint fk_city foreign key(city_id) references city(id),
    constraint name unique(city_id, category, name)
);