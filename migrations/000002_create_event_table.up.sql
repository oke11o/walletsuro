create table events
(
    id              serial
        constraint events_pk
            primary key,
    user_id         int          not null,
    amount          int          not null default 0,
    currency        varchar(255) not null,
    wallet_uuid     uuid         not null,
    type            varchar(255) not null default '',
    date            timestamp    not null default current_timestamp,
    additional_data varchar(255) null
);