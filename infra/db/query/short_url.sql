-- name: CreateShortURL :one
INSERT INTO short_links (
  url,
  slug,
  user_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetShortURL :one
SELECT * FROM short_links
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetShortURLBySlug :one
SELECT * FROM short_links
WHERE slug = $1 LIMIT 1;

-- name: ListShortURLs :many
SELECT * FROM short_links
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateShortURL :one
UPDATE short_links
SET
  url = $1,
  slug = $2
WHERE id = $3
RETURNING *;

-- name: DeleteShortURL :exec
DELETE FROM short_links
WHERE id = $1;