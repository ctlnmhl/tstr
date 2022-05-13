// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/common/v1/common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Test with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Test) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Test with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TestMultiError, or nil if none found.
func (m *Test) ValidateAll() error {
	return m.validate(true)
}

func (m *Test) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Labels

	// no validation rules for CronSchedule

	if all {
		switch v := interface{}(m.GetNextRunAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "NextRunAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "NextRunAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNextRunAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "NextRunAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRunConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "RunConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "RunConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRunConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "RunConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRegisteredAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "RegisteredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "RegisteredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegisteredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "RegisteredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetArchivedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "ArchivedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestValidationError{
					field:  "ArchivedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetArchivedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestValidationError{
				field:  "ArchivedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TestMultiError(errors)
	}

	return nil
}

// TestMultiError is an error wrapping multiple validation errors returned by
// Test.ValidateAll() if the designated constraints aren't met.
type TestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestMultiError) AllErrors() []error { return m }

// TestValidationError is the validation error returned by Test.Validate if the
// designated constraints aren't met.
type TestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestValidationError) ErrorName() string { return "TestValidationError" }

// Error satisfies the builtin error interface
func (e TestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestValidationError{}

// Validate checks the field values on TestSuite with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TestSuite) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestSuite with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TestSuiteMultiError, or nil
// if none found.
func (m *TestSuite) ValidateAll() error {
	return m.validate(true)
}

func (m *TestSuite) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Labels

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestSuiteValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestSuiteValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetArchivedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "ArchivedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TestSuiteValidationError{
					field:  "ArchivedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetArchivedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestSuiteValidationError{
				field:  "ArchivedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TestSuiteMultiError(errors)
	}

	return nil
}

// TestSuiteMultiError is an error wrapping multiple validation errors returned
// by TestSuite.ValidateAll() if the designated constraints aren't met.
type TestSuiteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestSuiteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestSuiteMultiError) AllErrors() []error { return m }

// TestSuiteValidationError is the validation error returned by
// TestSuite.Validate if the designated constraints aren't met.
type TestSuiteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestSuiteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestSuiteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestSuiteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestSuiteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestSuiteValidationError) ErrorName() string { return "TestSuiteValidationError" }

// Error satisfies the builtin error interface
func (e TestSuiteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestSuite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestSuiteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestSuiteValidationError{}

// Validate checks the field values on Run with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Run) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Run with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RunMultiError, or nil if none found.
func (m *Run) ValidateAll() error {
	return m.validate(true)
}

func (m *Run) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TestId

	if all {
		switch v := interface{}(m.GetTestRunConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "TestRunConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "TestRunConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTestRunConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunValidationError{
				field:  "TestRunConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RunnerId

	// no validation rules for Result

	for idx, item := range m.GetLogs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RunValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RunValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RunValidationError{
					field:  fmt.Sprintf("Logs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetScheduledAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "ScheduledAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "ScheduledAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetScheduledAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunValidationError{
				field:  "ScheduledAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStartedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "StartedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFinishedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "FinishedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunValidationError{
					field:  "FinishedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFinishedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunValidationError{
				field:  "FinishedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RunMultiError(errors)
	}

	return nil
}

// RunMultiError is an error wrapping multiple validation errors returned by
// Run.ValidateAll() if the designated constraints aren't met.
type RunMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RunMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RunMultiError) AllErrors() []error { return m }

// RunValidationError is the validation error returned by Run.Validate if the
// designated constraints aren't met.
type RunValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RunValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RunValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RunValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RunValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RunValidationError) ErrorName() string { return "RunValidationError" }

// Error satisfies the builtin error interface
func (e RunValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRun.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RunValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RunValidationError{}

// Validate checks the field values on Runner with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Runner) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Runner with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RunnerMultiError, or nil if none found.
func (m *Runner) ValidateAll() error {
	return m.validate(true)
}

func (m *Runner) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for AcceptTestLabelSelectors

	// no validation rules for RejectTestLabelSelectors

	if all {
		switch v := interface{}(m.GetRegisteredAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunnerValidationError{
					field:  "RegisteredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunnerValidationError{
					field:  "RegisteredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegisteredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunnerValidationError{
				field:  "RegisteredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastHeartbeatAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunnerValidationError{
					field:  "LastHeartbeatAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunnerValidationError{
					field:  "LastHeartbeatAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastHeartbeatAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunnerValidationError{
				field:  "LastHeartbeatAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RunnerMultiError(errors)
	}

	return nil
}

// RunnerMultiError is an error wrapping multiple validation errors returned by
// Runner.ValidateAll() if the designated constraints aren't met.
type RunnerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RunnerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RunnerMultiError) AllErrors() []error { return m }

// RunnerValidationError is the validation error returned by Runner.Validate if
// the designated constraints aren't met.
type RunnerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RunnerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RunnerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RunnerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RunnerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RunnerValidationError) ErrorName() string { return "RunnerValidationError" }

// Error satisfies the builtin error interface
func (e RunnerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRunner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RunnerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RunnerValidationError{}

// Validate checks the field values on AccessToken with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccessToken) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccessToken with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccessTokenMultiError, or
// nil if none found.
func (m *AccessToken) ValidateAll() error {
	return m.validate(true)
}

func (m *AccessToken) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Token

	if all {
		switch v := interface{}(m.GetIssuedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "IssuedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "IssuedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIssuedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessTokenValidationError{
				field:  "IssuedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessTokenValidationError{
				field:  "ExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRevokedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "RevokedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessTokenValidationError{
					field:  "RevokedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRevokedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessTokenValidationError{
				field:  "RevokedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccessTokenMultiError(errors)
	}

	return nil
}

// AccessTokenMultiError is an error wrapping multiple validation errors
// returned by AccessToken.ValidateAll() if the designated constraints aren't met.
type AccessTokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccessTokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccessTokenMultiError) AllErrors() []error { return m }

// AccessTokenValidationError is the validation error returned by
// AccessToken.Validate if the designated constraints aren't met.
type AccessTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessTokenValidationError) ErrorName() string { return "AccessTokenValidationError" }

// Error satisfies the builtin error interface
func (e AccessTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessTokenValidationError{}

// Validate checks the field values on Test_RunConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Test_RunConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Test_RunConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Test_RunConfigMultiError,
// or nil if none found.
func (m *Test_RunConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Test_RunConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ContainerImage

	// no validation rules for Command

	// no validation rules for Env

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Test_RunConfigValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Test_RunConfigValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Test_RunConfigValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Test_RunConfigMultiError(errors)
	}

	return nil
}

// Test_RunConfigMultiError is an error wrapping multiple validation errors
// returned by Test_RunConfig.ValidateAll() if the designated constraints
// aren't met.
type Test_RunConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Test_RunConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Test_RunConfigMultiError) AllErrors() []error { return m }

// Test_RunConfigValidationError is the validation error returned by
// Test_RunConfig.Validate if the designated constraints aren't met.
type Test_RunConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Test_RunConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Test_RunConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Test_RunConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Test_RunConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Test_RunConfigValidationError) ErrorName() string { return "Test_RunConfigValidationError" }

// Error satisfies the builtin error interface
func (e Test_RunConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTest_RunConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Test_RunConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Test_RunConfigValidationError{}

// Validate checks the field values on Run_Log with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Run_Log) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Run_Log with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in Run_LogMultiError, or nil if none found.
func (m *Run_Log) ValidateAll() error {
	return m.validate(true)
}

func (m *Run_Log) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Time

	// no validation rules for OutputType

	// no validation rules for Data

	if len(errors) > 0 {
		return Run_LogMultiError(errors)
	}

	return nil
}

// Run_LogMultiError is an error wrapping multiple validation errors returned
// by Run_Log.ValidateAll() if the designated constraints aren't met.
type Run_LogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Run_LogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Run_LogMultiError) AllErrors() []error { return m }

// Run_LogValidationError is the validation error returned by Run_Log.Validate
// if the designated constraints aren't met.
type Run_LogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Run_LogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Run_LogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Run_LogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Run_LogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Run_LogValidationError) ErrorName() string { return "Run_LogValidationError" }

// Error satisfies the builtin error interface
func (e Run_LogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRun_Log.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Run_LogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Run_LogValidationError{}
