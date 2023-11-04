-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balence,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;