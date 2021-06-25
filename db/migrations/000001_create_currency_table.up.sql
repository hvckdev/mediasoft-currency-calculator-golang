create table if not exists public.currency
(
    id         bigserial primary key,

    currency1  text               not null,
    currency2  text               not null,

    rate       float             not null,

    created_at timestamp default now() not null,
    updated_at timestamp default now() not null
);