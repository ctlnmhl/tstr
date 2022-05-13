// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Querier interface {
	ArchiveTest(ctx context.Context, id uuid.UUID) error
	ArchiveTestSuite(ctx context.Context, id uuid.UUID) error
	AssignRun(ctx context.Context, arg AssignRunParams) (AssignRunRow, error)
	AuthAccessToken(ctx context.Context, tokenHash string) (AuthAccessTokenRow, error)
	CreateTestRunConfig(ctx context.Context, arg CreateTestRunConfigParams) (TestRunConfig, error)
	DefineTestSuite(ctx context.Context, arg DefineTestSuiteParams) (TestSuite, error)
	GetAccessToken(ctx context.Context, id uuid.UUID) (GetAccessTokenRow, error)
	GetRun(ctx context.Context, id uuid.UUID) (GetRunRow, error)
	GetRunner(ctx context.Context, id uuid.UUID) (Runner, error)
	GetTest(ctx context.Context, id uuid.UUID) (GetTestRow, error)
	GetTestSuite(ctx context.Context, id uuid.UUID) (TestSuite, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	IssueAccessToken(ctx context.Context, arg IssueAccessTokenParams) (IssueAccessTokenRow, error)
	ListAccessTokens(ctx context.Context, arg ListAccessTokensParams) ([]ListAccessTokensRow, error)
	ListRunners(ctx context.Context, heartbeatSince sql.NullTime) ([]Runner, error)
	ListRuns(ctx context.Context, arg ListRunsParams) ([]ListRunsRow, error)
	ListTests(ctx context.Context, labels pgtype.JSONB) ([]ListTestsRow, error)
	ListTestsIDsMatchingLabelKeys(ctx context.Context, arg ListTestsIDsMatchingLabelKeysParams) ([]ListTestsIDsMatchingLabelKeysRow, error)
	ListTestsToSchedule(ctx context.Context) ([]ListTestsToScheduleRow, error)
	RegisterRunner(ctx context.Context, arg RegisterRunnerParams) (Runner, error)
	RegisterTest(ctx context.Context, arg RegisterTestParams) (RegisterTestRow, error)
	RevokeAccessToken(ctx context.Context, id uuid.UUID) error
	ScheduleRun(ctx context.Context, arg ScheduleRunParams) (ScheduleRunRow, error)
	UpdateRun(ctx context.Context, arg UpdateRunParams) error
	UpdateTest(ctx context.Context, arg UpdateTestParams) error
	UpdateTestSuite(ctx context.Context, arg UpdateTestSuiteParams) error
	listTestSuites(ctx context.Context, labels pgtype.JSONB) ([]TestSuite, error)
}

var _ Querier = (*Queries)(nil)
