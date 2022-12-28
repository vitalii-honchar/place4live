create table city
(
    "id" serial primary key,
    "name" varchar(256) not null
);

create table category
(
    "id" serial primary key,
    "name" varchar(256) not null,
    "city_id" int not null,
    "updated_at" timestamp not null,
    constraint fk_city foreign key(city_id) references city(id)
);

create table category_price
(
    "id"          serial primary key,
    "category_id" int not null,
    "name" varchar(256) not null,
    "avg" float not null,
    "min" float not null,
    "max" float not null,
    "updated_at" timestamp not null,
    constraint fk_category foreign key(category_id) references category(id)
);