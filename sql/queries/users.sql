-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE name = $1
LIMIT 1;

-- name: DeleteALLUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT *
FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE name = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = $1, updated_at = $2
WHERE name = $3;

