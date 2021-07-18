CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table wallets
(
    id serial
        constraint wallets_pk
            primary key,
    user_id int not null,
    wallet uuid not null,
    amount int not null default 0
);

create index wallets_user_id_index
	on wallets (user_id);

create index wallets_wallet_index
	on wallets (wallet);

