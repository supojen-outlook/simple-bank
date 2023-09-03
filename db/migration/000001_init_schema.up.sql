create table accounts (
    id bigserial primary key,
    owner varchar(10) not null,
    balance bigint not null,
    currency varchar(10) not null,
    created_at timestamptz not null default now()
);

create table entries (
    id bigserial primary key,
    amount bigint not null,
    account_id bigint not null references accounts(id),
    created_at timestamptz not null default now()
);

create table transfers (
    id bigserial primary key,
    amount bigint not null,
    from_account_id bigint not null references accounts(id),
    to_account_id bigint not null references accounts(id),
    created_at timestamptz not null default now()
);