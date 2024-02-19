-- +migrate Up

create table accounts (
    id                      integer             primary key,
    email                   string              not null unique,
    password                string              not null,
    bgg_account             string              unique.
    created_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
    updated_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),

    constraint chk_email check (email <> ''),
    constraint chk_password check (password <> '')
);

create table games (
    id              integer     primary key,
    name            string      not null,
    description     string      not null,
    image           string      not null,
    thumbnail       string      not null,
    year            int not     null,
    created_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
    updated_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now'))
);

create table collections (
    id  integer primary key,
    account_id  integer constraint fk_accounts_collections references accounts on delete cascade,
    game_id integer constraint fk_games_collections references accounts on delete cascade,
    created_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
    updated_at              datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now'))
);