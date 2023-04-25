// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: entrie.sql

package db

import (
	"context"
)

const createEntrie = `-- name: CreateEntrie :one
INSERT INTO entries (
    account_id, 
    amount

) VALUES ($1, $2)
RETURNING id, account_id, amount, created_at
`

type CreateEntrieParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntrie(ctx context.Context, arg CreateEntrieParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntrie, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntrie = `-- name: DeleteEntrie :exec
DELETE FROM entries WHERE id = $1
`

func (q *Queries) DeleteEntrie(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntrie, id)
	return err
}

const getEntrie = `-- name: GetEntrie :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1
`

func (q *Queries) GetEntrie(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntrie, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM entries
WHERE account_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3
`

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntrie = `-- name: UpdateEntrie :one
UPDATE entries SET amount = $1
WHERE id = $2
RETURNING id, account_id, amount, created_at
`

type UpdateEntrieParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateEntrie(ctx context.Context, arg UpdateEntrieParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntrie, arg.Amount, arg.ID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}