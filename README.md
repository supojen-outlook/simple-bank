## Create an postgres container

> docker pull postgres:alpine3.17
> docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:alpine3.17

```sql

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
  created_at timestamptz not null default now(),
  account_id bigint not null references accounts(id)
);

create table transfers (
  id bigserial primary key,
  amount bigint not null,
  account_id bigint references accounts(id),
  from_account_id bigint not null references accounts(id),
  to_account_id bigint not null references accounts(id)
);

```


## Database Migration 

> brew install golang-migrate
> migrate -help
> migrate create -ext sql -dir db/migration -seq init_schema
* -ext        : extension name is sql
* -dir        : output directory
* -seq        : swith sequential number in the file name
* init_shcema : the ffile name


__up file__
```sql
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
  created_at timestamptz not null default now(),
  account_id bigint not null references accounts(id)
);

create table transfers (
  id bigserial primary key,
  amount bigint not null,
  account_id bigint references accounts(id),
  from_account_id bigint not null references accounts(id),
  to_account_id bigint not null references accounts(id)
);
```

__down file__
```sql
drop table if exists entries;
drop table if exists transfers;
drop table if exists accounts;
```