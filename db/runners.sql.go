// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: runners.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const getRunner = `-- name: GetRunner :one
SELECT id, name, accept_test_label_selectors, reject_test_label_selectors, registered_at, last_heartbeat_at
FROM runners
WHERE id = $1
`

func (q *Queries) GetRunner(ctx context.Context, db DBTX, id uuid.UUID) (Runner, error) {
	row := db.QueryRow(ctx, getRunner, id)
	var i Runner
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AcceptTestLabelSelectors,
		&i.RejectTestLabelSelectors,
		&i.RegisteredAt,
		&i.LastHeartbeatAt,
	)
	return i, err
}

const listRunners = `-- name: ListRunners :many
SELECT id, name, accept_test_label_selectors, reject_test_label_selectors, registered_at, last_heartbeat_at
FROM runners
ORDER by last_heartbeat_at DESC, registered_at
`

func (q *Queries) ListRunners(ctx context.Context, db DBTX) ([]Runner, error) {
	rows, err := db.Query(ctx, listRunners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Runner
	for rows.Next() {
		var i Runner
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AcceptTestLabelSelectors,
			&i.RejectTestLabelSelectors,
			&i.RegisteredAt,
			&i.LastHeartbeatAt,
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

const queryRunners = `-- name: QueryRunners :many
SELECT id, name, accept_test_label_selectors, reject_test_label_selectors, registered_at, last_heartbeat_at
FROM runners
WHERE
  ($1::uuid[] IS NULL OR runners.id = ANY ($1::uuid[])) AND
  ($2::timestamptz IS NULL) OR last_heartbeat_at > $2
ORDER by name ASC
`

type QueryRunnersParams struct {
	Ids                []uuid.UUID
	LastHeartbeatSince sql.NullTime
}

func (q *Queries) QueryRunners(ctx context.Context, db DBTX, arg QueryRunnersParams) ([]Runner, error) {
	rows, err := db.Query(ctx, queryRunners, arg.Ids, arg.LastHeartbeatSince)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Runner
	for rows.Next() {
		var i Runner
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AcceptTestLabelSelectors,
			&i.RejectTestLabelSelectors,
			&i.RegisteredAt,
			&i.LastHeartbeatAt,
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

const registerRunner = `-- name: RegisterRunner :one
INSERT INTO runners (name, accept_test_label_selectors, reject_test_label_selectors, last_heartbeat_at)
VALUES (
  $1,
  $2,
  $3,
  CURRENT_TIMESTAMP
)
RETURNING id, name, accept_test_label_selectors, reject_test_label_selectors, registered_at, last_heartbeat_at
`

type RegisterRunnerParams struct {
	Name                     string
	AcceptTestLabelSelectors pgtype.JSONB
	RejectTestLabelSelectors pgtype.JSONB
}

func (q *Queries) RegisterRunner(ctx context.Context, db DBTX, arg RegisterRunnerParams) (Runner, error) {
	row := db.QueryRow(ctx, registerRunner, arg.Name, arg.AcceptTestLabelSelectors, arg.RejectTestLabelSelectors)
	var i Runner
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AcceptTestLabelSelectors,
		&i.RejectTestLabelSelectors,
		&i.RegisteredAt,
		&i.LastHeartbeatAt,
	)
	return i, err
}

const updateRunnerHeartbeat = `-- name: UpdateRunnerHeartbeat :exec
UPDATE runners
SET last_heartbeat_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) UpdateRunnerHeartbeat(ctx context.Context, db DBTX, id uuid.UUID) error {
	_, err := db.Exec(ctx, updateRunnerHeartbeat, id)
	return err
}
