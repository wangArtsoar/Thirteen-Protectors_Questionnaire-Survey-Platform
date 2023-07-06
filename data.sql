-- role
create table if not exists public.role
(
    id   integer generated always as identity
    primary key,
    name varchar(50)
    );

alter table public.role
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
    password  char(60),
    create_at date,
    update_at date,
    is_delete smallint,
    is_valid  smallint
    );

alter table public."user"
    owner to postgres;
