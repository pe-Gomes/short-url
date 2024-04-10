// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: short_url.sql

package db

import (
	"context"
)

const createShortURL = `-- name: CreateShortURL :one
INSERT INTO short_links (
  url,
  slug,
  user_id
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, url, slug, created_at, updated_at
`

type CreateShortURLParams struct {
	Url    string `json:"url"`
	Slug   string `json:"slug"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) CreateShortURL(ctx context.Context, arg CreateShortURLParams) (ShortLink, error) {
	row := q.db.QueryRow(ctx, createShortURL, arg.Url, arg.Slug, arg.UserID)
	var i ShortLink
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Url,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteShortURL = `-- name: DeleteShortURL :exec
DELETE FROM short_links
WHERE id = $1
`

func (q *Queries) DeleteShortURL(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteShortURL, id)
	return err
}

const getShortURL = `-- name: GetShortURL :one
SELECT id, user_id, url, slug, created_at, updated_at FROM short_links
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetShortURL(ctx context.Context, id int64) (ShortLink, error) {
	row := q.db.QueryRow(ctx, getShortURL, id)
	var i ShortLink
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Url,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getShortURLBySlug = `-- name: GetShortURLBySlug :one
SELECT id, user_id, url, slug, created_at, updated_at FROM short_links
WHERE slug = $1 LIMIT 1
`

func (q *Queries) GetShortURLBySlug(ctx context.Context, slug string) (ShortLink, error) {
	row := q.db.QueryRow(ctx, getShortURLBySlug, slug)
	var i ShortLink
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Url,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listShortURLs = `-- name: ListShortURLs :many
SELECT id, user_id, url, slug, created_at, updated_at FROM short_links
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListShortURLsParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShortURLs(ctx context.Context, arg ListShortURLsParams) ([]ShortLink, error) {
	rows, err := q.db.Query(ctx, listShortURLs, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShortLink{}
	for rows.Next() {
		var i ShortLink
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Url,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateShortURL = `-- name: UpdateShortURL :one
UPDATE short_links
SET
  url = $1,
  slug = $2
WHERE id = $3
RETURNING id, user_id, url, slug, created_at, updated_at
`

type UpdateShortURLParams struct {
	Url  string `json:"url"`
	Slug string `json:"slug"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateShortURL(ctx context.Context, arg UpdateShortURLParams) (ShortLink, error) {
	row := q.db.QueryRow(ctx, updateShortURL, arg.Url, arg.Slug, arg.ID)
	var i ShortLink
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Url,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
