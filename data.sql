-- role
create table if not exists public.role
(
    id   integer not null
    primary key,
    name varchar(50)
    );

alter table public.role
    owner to postgres;

-- token
create table if not exists public.token
(
    id           integer not null
    primary key,
    user_id      varchar(32),
    access_token varchar(50),
    is_valid     smallint
    );

alter table public.token
    owner to postgres;

-- user
create table if not exists public."user"
(
    id        varchar(50) not null
    primary key,
    role_id   integer,
    name      varchar(50),
    email     varchar(50)
    unique,
    password  varchar(50),
    create_at date,
    update_at date,
    is_delete smallint,
    is_valid  smallint
    );

alter table public."user"
    owner to postgres;
