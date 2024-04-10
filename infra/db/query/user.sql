-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
  name = $1,
  password = $2
WHERE id = $3
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;