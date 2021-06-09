// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-runner-api/ocp-runner-api.proto

package ocp_runner_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Runner with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Runner) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Guid

	// no validation rules for Os

	// no validation rules for Arch

	return nil
}

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

// Validate checks the field values on ListFilters with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListFilters) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListFiltersValidationError is the validation error returned by
// ListFilters.Validate if the designated constraints aren't met.
type ListFiltersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFiltersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFiltersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFiltersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFiltersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFiltersValidationError) ErrorName() string { return "ListFiltersValidationError" }

// Error satisfies the builtin error interface
func (e ListFiltersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFilters.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFiltersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFiltersValidationError{}

// Validate checks the field values on RunnersList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RunnersList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRunners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RunnersListValidationError{
					field:  fmt.Sprintf("Runners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RunnersListValidationError is the validation error returned by
// RunnersList.Validate if the designated constraints aren't met.
type RunnersListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RunnersListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RunnersListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RunnersListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RunnersListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RunnersListValidationError) ErrorName() string { return "RunnersListValidationError" }

// Error satisfies the builtin error interface
func (e RunnersListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRunnersList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RunnersListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RunnersListValidationError{}

// Validate checks the field values on OperationResult with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OperationResult) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Guid

	// no validation rules for Error

	return nil
}

// OperationResultValidationError is the validation error returned by
// OperationResult.Validate if the designated constraints aren't met.
type OperationResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OperationResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OperationResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OperationResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OperationResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OperationResultValidationError) ErrorName() string { return "OperationResultValidationError" }

// Error satisfies the builtin error interface
func (e OperationResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOperationResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OperationResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OperationResultValidationError{}