-- +migrate Up

create table accounts (
    id                    serial              primary key,
    email                 varchar(1024)       not null unique,
    password              varchar(64)           not null,
    bgg_account           varchar(32)         unique,
    created_at            timestamp           not null default(current_timestamp),
    updated_at            timestamp           not null default(current_timestamp), 

    constraint chk_email check (email <> ''),
    constraint chk_password check (password <> '')
);

create table games (
    id                      int                 primary key,
    name                    varchar(1024)       not null,
    description             varchar(1024)       not null,
    image                   varchar(1024)       not null,
    thumbnail               varchar(1024)       not null,
    year                    int                 not null,
    created_at              timestamp           not null default(current_timestamp),
    updated_at              timestamp           not null default(current_timestamp)
);

create table collections (
    id  serial  primary key,
    account_id  integer constraint fk_accounts_collections references accounts on delete cascade,
    game_id integer constraint fk_games_collections references games on delete cascade,
    created_at            timestamp           not null default(current_timestamp),
    updated_at            timestamp           not null default(current_timestamp)
);