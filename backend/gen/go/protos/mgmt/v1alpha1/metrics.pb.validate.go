// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mgmt/v1alpha1/metrics.proto

package mgmtv1alpha1

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

// Validate checks the field values on Date with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Date) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Date with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DateMultiError, or nil if none found.
func (m *Date) ValidateAll() error {
	return m.validate(true)
}

func (m *Date) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Year

	// no validation rules for Month

	// no validation rules for Day

	if len(errors) > 0 {
		return DateMultiError(errors)
	}

	return nil
}

// DateMultiError is an error wrapping multiple validation errors returned by
// Date.ValidateAll() if the designated constraints aren't met.
type DateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DateMultiError) AllErrors() []error { return m }

// DateValidationError is the validation error returned by Date.Validate if the
// designated constraints aren't met.
type DateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DateValidationError) ErrorName() string { return "DateValidationError" }

// Error satisfies the builtin error interface
func (e DateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DateValidationError{}

// Validate checks the field values on GetDailyMetricCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDailyMetricCountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDailyMetricCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDailyMetricCountRequestMultiError, or nil if none found.
func (m *GetDailyMetricCountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDailyMetricCountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDailyMetricCountRequestValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDailyMetricCountRequestValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDailyMetricCountRequestValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEnd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDailyMetricCountRequestValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDailyMetricCountRequestValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDailyMetricCountRequestValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Metric

	switch v := m.Identifier.(type) {
	case *GetDailyMetricCountRequest_AccountId:
		if v == nil {
			err := GetDailyMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for AccountId
	case *GetDailyMetricCountRequest_JobId:
		if v == nil {
			err := GetDailyMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for JobId
	case *GetDailyMetricCountRequest_RunId:
		if v == nil {
			err := GetDailyMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for RunId
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return GetDailyMetricCountRequestMultiError(errors)
	}

	return nil
}

// GetDailyMetricCountRequestMultiError is an error wrapping multiple
// validation errors returned by GetDailyMetricCountRequest.ValidateAll() if
// the designated constraints aren't met.
type GetDailyMetricCountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDailyMetricCountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDailyMetricCountRequestMultiError) AllErrors() []error { return m }

// GetDailyMetricCountRequestValidationError is the validation error returned
// by GetDailyMetricCountRequest.Validate if the designated constraints aren't met.
type GetDailyMetricCountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDailyMetricCountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDailyMetricCountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDailyMetricCountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDailyMetricCountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDailyMetricCountRequestValidationError) ErrorName() string {
	return "GetDailyMetricCountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDailyMetricCountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDailyMetricCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDailyMetricCountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDailyMetricCountRequestValidationError{}

// Validate checks the field values on GetDailyMetricCountResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDailyMetricCountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDailyMetricCountResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDailyMetricCountResponseMultiError, or nil if none found.
func (m *GetDailyMetricCountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDailyMetricCountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetDailyMetricCountResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetDailyMetricCountResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetDailyMetricCountResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetDailyMetricCountResponseMultiError(errors)
	}

	return nil
}

// GetDailyMetricCountResponseMultiError is an error wrapping multiple
// validation errors returned by GetDailyMetricCountResponse.ValidateAll() if
// the designated constraints aren't met.
type GetDailyMetricCountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDailyMetricCountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDailyMetricCountResponseMultiError) AllErrors() []error { return m }

// GetDailyMetricCountResponseValidationError is the validation error returned
// by GetDailyMetricCountResponse.Validate if the designated constraints
// aren't met.
type GetDailyMetricCountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDailyMetricCountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDailyMetricCountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDailyMetricCountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDailyMetricCountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDailyMetricCountResponseValidationError) ErrorName() string {
	return "GetDailyMetricCountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDailyMetricCountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDailyMetricCountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDailyMetricCountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDailyMetricCountResponseValidationError{}

// Validate checks the field values on DayResult with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DayResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DayResult with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DayResultMultiError, or nil
// if none found.
func (m *DayResult) ValidateAll() error {
	return m.validate(true)
}

func (m *DayResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DayResultValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DayResultValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DayResultValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Count

	if len(errors) > 0 {
		return DayResultMultiError(errors)
	}

	return nil
}

// DayResultMultiError is an error wrapping multiple validation errors returned
// by DayResult.ValidateAll() if the designated constraints aren't met.
type DayResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DayResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DayResultMultiError) AllErrors() []error { return m }

// DayResultValidationError is the validation error returned by
// DayResult.Validate if the designated constraints aren't met.
type DayResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DayResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DayResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DayResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DayResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DayResultValidationError) ErrorName() string { return "DayResultValidationError" }

// Error satisfies the builtin error interface
func (e DayResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDayResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DayResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DayResultValidationError{}

// Validate checks the field values on GetMetricCountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetMetricCountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMetricCountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMetricCountRequestMultiError, or nil if none found.
func (m *GetMetricCountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMetricCountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetMetricCountRequestValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetMetricCountRequestValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMetricCountRequestValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEnd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetMetricCountRequestValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetMetricCountRequestValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMetricCountRequestValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Metric

	switch v := m.Identifier.(type) {
	case *GetMetricCountRequest_AccountId:
		if v == nil {
			err := GetMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for AccountId
	case *GetMetricCountRequest_JobId:
		if v == nil {
			err := GetMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for JobId
	case *GetMetricCountRequest_RunId:
		if v == nil {
			err := GetMetricCountRequestValidationError{
				field:  "Identifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for RunId
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return GetMetricCountRequestMultiError(errors)
	}

	return nil
}

// GetMetricCountRequestMultiError is an error wrapping multiple validation
// errors returned by GetMetricCountRequest.ValidateAll() if the designated
// constraints aren't met.
type GetMetricCountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMetricCountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMetricCountRequestMultiError) AllErrors() []error { return m }

// GetMetricCountRequestValidationError is the validation error returned by
// GetMetricCountRequest.Validate if the designated constraints aren't met.
type GetMetricCountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMetricCountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMetricCountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMetricCountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMetricCountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMetricCountRequestValidationError) ErrorName() string {
	return "GetMetricCountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMetricCountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMetricCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMetricCountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMetricCountRequestValidationError{}

// Validate checks the field values on GetMetricCountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetMetricCountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMetricCountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMetricCountResponseMultiError, or nil if none found.
func (m *GetMetricCountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMetricCountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return GetMetricCountResponseMultiError(errors)
	}

	return nil
}

// GetMetricCountResponseMultiError is an error wrapping multiple validation
// errors returned by GetMetricCountResponse.ValidateAll() if the designated
// constraints aren't met.
type GetMetricCountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMetricCountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMetricCountResponseMultiError) AllErrors() []error { return m }

// GetMetricCountResponseValidationError is the validation error returned by
// GetMetricCountResponse.Validate if the designated constraints aren't met.
type GetMetricCountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMetricCountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMetricCountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMetricCountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMetricCountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMetricCountResponseValidationError) ErrorName() string {
	return "GetMetricCountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetMetricCountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMetricCountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMetricCountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMetricCountResponseValidationError{}