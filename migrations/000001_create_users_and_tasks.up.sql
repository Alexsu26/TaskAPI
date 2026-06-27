create table users (
    id  bigserial   primary key,
    email   text    not null    unique,
    name    text not null,
    password_hash   text    not null,
    created_at  timestamptz not null    default NOW(),
    updated_at  timestamptz not null    default NOW()
);

create table tasks(
    id bigserial    primary key,
    user_id bigint not null    references users(id),
    title   text    not null,
    description text  not null  default '',
    status  text    not null    default 'todo',
    created_at  timestamptz not null    default NOW(),
    updated_at  timestamptz not null    default NOW()
);
