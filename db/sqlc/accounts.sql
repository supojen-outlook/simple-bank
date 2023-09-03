-- name: CreateAccount :one
insert into accounts(
  owner,
  balance,
  currency
) values (
  $1, $2, $3
) returning *;

-- name: GetAccount :one
select * from accounts
where id = $1 limit 1;

-- name: GetAccountForUpdate :one
select * from accounts
where id = $1 limit 1
for no key update;

-- name: ListAccounts :many
select * from accounts
where id > $1
order by id
limit $2;

-- name: UpdtaeAccount :one
update accounts
set balance = $2
where id = $1
RETURNING * ;

-- name: AddAccountBalance :one
update accounts
set balance = balance + sqlc.arg(amount)
where id = sqlc.arg(id)
RETURNING * ;

-- name: DeleteAccount :exec
delete from accounts where id = $1;