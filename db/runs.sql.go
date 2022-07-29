// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: runs.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const appendLogsToRun = `-- name: AppendLogsToRun :exec
UPDATE runs
SET logs = COALESCE(logs, '[]'::jsonb) || $1
WHERE id = $2
`

type AppendLogsToRunParams struct {
	Logs pgtype.JSONB
	ID   uuid.UUID
}

func (q *Queries) AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error {
	_, err := db.Exec(ctx, appendLogsToRun, arg.Logs, arg.ID)
	return err
}

const assignRun = `-- name: AssignRun :one
UPDATE runs
SET runner_id = $1::uuid
WHERE runs.id = (
  SELECT id
  FROM runs AS selected_runs
  WHERE selected_runs.test_id = ANY($2::uuid[]) AND selected_runs.runner_id IS NULL
  ORDER BY selected_runs.scheduled_at ASC
  LIMIT 1
)
RETURNING runs.id, runs.test_id, runs.test_run_config, runs.runner_id, runs.result, runs.logs, runs.result_data, runs.scheduled_at, runs.started_at, runs.finished_at
`

type AssignRunParams struct {
	RunnerID uuid.UUID
	TestIds  []uuid.UUID
}

func (q *Queries) AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (Run, error) {
	row := db.QueryRow(ctx, assignRun, arg.RunnerID, arg.TestIds)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const getRun = `-- name: GetRun :one
SELECT id, test_id, test_run_config, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
FROM runs
WHERE runs.id = $1
`

func (q *Queries) GetRun(ctx context.Context, db DBTX, id uuid.UUID) (Run, error) {
	row := db.QueryRow(ctx, getRun, id)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const listRuns = `-- name: ListRuns :many
SELECT id, test_id, test_run_config, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
FROM runs
`

func (q *Queries) ListRuns(ctx context.Context, db DBTX) ([]Run, error) {
	rows, err := db.Query(ctx, listRuns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Run
	for rows.Next() {
		var i Run
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ResultData,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
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

const queryRuns = `-- name: QueryRuns :many
SELECT id, test_id, test_run_config, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
FROM runs
WHERE
  ($1::uuid[] IS NULL OR runs.id = ANY ($1::uuid[])) AND
  ($2::uuid[] IS NULL OR runs.test_id = ANY ($2::uuid[])) AND
  ($3::uuid[] IS NULL OR runs.test_id = ANY (
      SELECT tests.id
      FROM test_suites
      JOIN tests
      ON tests.labels @> test_suites.labels
      WHERE test_suites.id = ANY ($3::uuid[])
    )) AND
  ($4::uuid[] IS NULL OR runner_id = ANY ($4::uuid[])) AND
  ($5::text[] IS NULL OR result::text = ANY ($5::text[])) AND
  ($6::timestamptz IS NULL OR scheduled_at < $6::timestamptz) AND
  ($7::timestamptz IS NULL OR scheduled_at > $7::timestamptz) AND
  ($8::timestamptz IS NULL OR started_at < $8::timestamptz) AND
  ($9::timestamptz IS NULL OR started_at > $9::timestamptz) AND
  ($10::timestamptz IS NULL OR finished_at < $10::timestamptz) AND
  ($11::timestamptz IS NULL OR finished_at > $11::timestamptz)
`

type QueryRunsParams struct {
	Ids             []uuid.UUID
	TestIds         []uuid.UUID
	TestSuiteIds    []uuid.UUID
	RunnerIds       []uuid.UUID
	Results         []string
	ScheduledBefore sql.NullTime
	ScheduledAfter  sql.NullTime
	StartedBefore   sql.NullTime
	StartedAfter    sql.NullTime
	FinishedBefore  sql.NullTime
	FinishedAfter   sql.NullTime
}

func (q *Queries) QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]Run, error) {
	rows, err := db.Query(ctx, queryRuns,
		arg.Ids,
		arg.TestIds,
		arg.TestSuiteIds,
		arg.RunnerIds,
		arg.Results,
		arg.ScheduledBefore,
		arg.ScheduledAfter,
		arg.StartedBefore,
		arg.StartedAfter,
		arg.FinishedBefore,
		arg.FinishedAfter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Run
	for rows.Next() {
		var i Run
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ResultData,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
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

const resetOrphanedRuns = `-- name: ResetOrphanedRuns :exec
UPDATE runs
SET runner_id = NULL
WHERE
  result = 'unknown' AND
  started_at IS NULL AND
  scheduled_at < $1::timestamptz
`

func (q *Queries) ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error {
	_, err := db.Exec(ctx, resetOrphanedRuns, before)
	return err
}

const runSummaryForRunner = `-- name: RunSummaryForRunner :many
SELECT id, test_id, test_run_config, runner_id, result, scheduled_at, started_at, finished_at, result_data
FROM runs
WHERE runs.runner_id = $1::uuid
ORDER by runs.scheduled_at desc
LIMIT $2
`

type RunSummaryForRunnerParams struct {
	RunnerID uuid.UUID
	Limit    int32
}

type RunSummaryForRunnerRow struct {
	ID            uuid.UUID
	TestID        uuid.UUID
	TestRunConfig pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
	ResultData    pgtype.JSONB
}

func (q *Queries) RunSummaryForRunner(ctx context.Context, db DBTX, arg RunSummaryForRunnerParams) ([]RunSummaryForRunnerRow, error) {
	rows, err := db.Query(ctx, runSummaryForRunner, arg.RunnerID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RunSummaryForRunnerRow
	for rows.Next() {
		var i RunSummaryForRunnerRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.RunnerID,
			&i.Result,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ResultData,
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

const runSummaryForTest = `-- name: RunSummaryForTest :many
SELECT id, test_id, test_run_config, runner_id, result, scheduled_at, started_at, finished_at, result_data
FROM runs
WHERE runs.test_id = $1
ORDER by runs.scheduled_at desc
LIMIT $2
`

type RunSummaryForTestParams struct {
	TestID uuid.UUID
	Limit  int32
}

type RunSummaryForTestRow struct {
	ID            uuid.UUID
	TestID        uuid.UUID
	TestRunConfig pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
	ResultData    pgtype.JSONB
}

func (q *Queries) RunSummaryForTest(ctx context.Context, db DBTX, arg RunSummaryForTestParams) ([]RunSummaryForTestRow, error) {
	rows, err := db.Query(ctx, runSummaryForTest, arg.TestID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RunSummaryForTestRow
	for rows.Next() {
		var i RunSummaryForTestRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.RunnerID,
			&i.Result,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ResultData,
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

const scheduleRun = `-- name: ScheduleRun :one
WITH latest_run_config AS (
  SELECT run_config
  FROM tests
  WHERE tests.id = $1
)
INSERT INTO runs (test_id, test_run_config)
VALUES ($1, latest_run_config)
RETURNING id, test_id, test_run_config, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
`

func (q *Queries) ScheduleRun(ctx context.Context, db DBTX, testID uuid.UUID) (Run, error) {
	row := db.QueryRow(ctx, scheduleRun, testID)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const updateResultData = `-- name: UpdateResultData :exec
UPDATE runs
SET result_data = result_data || $1::jsonb 
WHERE id = $2
`

type UpdateResultDataParams struct {
	ResultData pgtype.JSONB
	ID         uuid.UUID
}

func (q *Queries) UpdateResultData(ctx context.Context, db DBTX, arg UpdateResultDataParams) error {
	_, err := db.Exec(ctx, updateResultData, arg.ResultData, arg.ID)
	return err
}

const updateRun = `-- name: UpdateRun :exec
UPDATE runs
SET
  result = $1,
  started_at = $2::timestamptz,
  finished_at = $3::timestamptz
WHERE id = $4
`

type UpdateRunParams struct {
	Result     NullRunResult
	StartedAt  sql.NullTime
	FinishedAt sql.NullTime
	ID         uuid.UUID
}

func (q *Queries) UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error {
	_, err := db.Exec(ctx, updateRun,
		arg.Result,
		arg.StartedAt,
		arg.FinishedAt,
		arg.ID,
	)
	return err
}
