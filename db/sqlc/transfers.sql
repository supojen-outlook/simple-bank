-- name: CreateTransfer :one
insert into transfers(
    amount,
    from_account_id,
    to_account_id
) values (
    $1, $2, $3
) returning *;

-- name: GetTransfer :one
select * from transfers
where id = $1 limit 1;

-- name: ListTransfersFrom :many
select * from transfers
where id > $1 and from_account_id = $2
order by id
limit $3;

-- name: ListTransfersTo :many
select * from transfers
where id > $1 and to_account_id = $2
order by id
limit $3;

-- name: DeleteTransfer :exec
delete from transfers where id = $1;
