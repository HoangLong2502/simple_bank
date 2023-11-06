-- name: CreateTransfer :one
INSERT INTO tranfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM tranfers
WHERE id = $1;