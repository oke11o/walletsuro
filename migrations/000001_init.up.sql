CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

create table wallets
(
    uuid     uuid         not null,
    user_id  int          not null,
    amount   int          not null default 0,
    currency varchar(255) not null
);

create
unique index wallet_uuid_uindex
	on wallets (uuid);

alter table wallets
    add constraint wallets_pk
        primary key (uuid);

create
index wallets_user_id_index
	on wallets (user_id);

