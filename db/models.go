// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type AccessTokenScope string

const (
	AccessTokenScopeAdmin   AccessTokenScope = "admin"
	AccessTokenScopeControl AccessTokenScope = "control"
	AccessTokenScopeData    AccessTokenScope = "data"
	AccessTokenScopeRunner  AccessTokenScope = "runner"
)

func (e *AccessTokenScope) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccessTokenScope(s)
	case string:
		*e = AccessTokenScope(s)
	default:
		return fmt.Errorf("unsupported scan type for AccessTokenScope: %T", src)
	}
	return nil
}

type NullAccessTokenScope struct {
	AccessTokenScope AccessTokenScope
	Valid            bool // Valid is true if AccessTokenScope is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccessTokenScope) Scan(value interface{}) error {
	if value == nil {
		ns.AccessTokenScope, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccessTokenScope.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccessTokenScope) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.AccessTokenScope, nil
}

type RunResult string

const (
	RunResultUnknown RunResult = "unknown"
	RunResultPass    RunResult = "pass"
	RunResultFail    RunResult = "fail"
	RunResultError   RunResult = "error"
)

func (e *RunResult) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RunResult(s)
	case string:
		*e = RunResult(s)
	default:
		return fmt.Errorf("unsupported scan type for RunResult: %T", src)
	}
	return nil
}

type NullRunResult struct {
	RunResult RunResult
	Valid     bool // Valid is true if RunResult is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRunResult) Scan(value interface{}) error {
	if value == nil {
		ns.RunResult, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RunResult.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRunResult) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.RunResult, nil
}

type AccessToken struct {
	ID                 uuid.UUID
	Name               string
	TokenHash          string
	Scopes             []AccessTokenScope
	IssuedAt           sql.NullTime
	ExpiresAt          sql.NullTime
	RevokedAt          sql.NullTime
	NamespaceSelectors []string
}

type Run struct {
	ID            uuid.UUID
	TestID        uuid.UUID
	TestRunConfig pgtype.JSONB
	TestMatrixID  uuid.NullUUID
	Labels        pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	Logs          pgtype.JSONB
	ResultData    pgtype.JSONB
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
}

type Runner struct {
	ID                       uuid.UUID
	Name                     string
	AcceptTestLabelSelectors pgtype.JSONB
	RejectTestLabelSelectors pgtype.JSONB
	RegisteredAt             sql.NullTime
	LastHeartbeatAt          sql.NullTime
	NamespaceSelectors       []string
}

type SchemaMigration struct {
	Version string
}

type Test struct {
	ID           uuid.UUID
	Name         string
	RunConfig    pgtype.JSONB
	Labels       pgtype.JSONB
	Matrix       pgtype.JSONB
	CronSchedule sql.NullString
	NextRunAt    sql.NullTime
	RegisteredAt sql.NullTime
	UpdatedAt    sql.NullTime
	Namespace    string
}
