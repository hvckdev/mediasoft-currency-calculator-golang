create table if not exists public.currency
(
    id         bigserial primary key,

    currency1  text               not null,
    currency2  text               not null,

    rate       float4             not null,

    created_at date default now() not null,
    updated_at date default now() not null
);