-- name: CreateEntry :one
insert into entries(
    amount,
    account_id
) values (
    $1, $2
) returning *;

-- name: GetEntry :one
select * from entries
where id = $1 limit 1;

-- name: ListEntries :many
select * from entries
where id > $1 and account_id = $2
order by created_at desc
limit $3;

-- name: DeleteEntry :exec
delete from entries where id = $1;