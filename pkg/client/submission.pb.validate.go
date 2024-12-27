// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: submission.proto

package client

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

// define the regex for a UUID once up-front
var _submission_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on FindStudentFormSubmissionsRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *FindStudentFormSubmissionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindStudentFormSubmissionsRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// FindStudentFormSubmissionsRequestMultiError, or nil if none found.
func (m *FindStudentFormSubmissionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FindStudentFormSubmissionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetStudentUsername()); l < 4 || l > 30 {
		err := FindStudentFormSubmissionsRequestValidationError{
			field:  "StudentUsername",
			reason: "value length must be between 4 and 30 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetFormId()); err != nil {
		err = FindStudentFormSubmissionsRequestValidationError{
			field:  "FormId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FindStudentFormSubmissionsRequestMultiError(errors)
	}

	return nil
}

func (m *FindStudentFormSubmissionsRequest) _validateUuid(uuid string) error {
	if matched := _submission_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// FindStudentFormSubmissionsRequestMultiError is an error wrapping multiple
// validation errors returned by
// FindStudentFormSubmissionsRequest.ValidateAll() if the designated
// constraints aren't met.
type FindStudentFormSubmissionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindStudentFormSubmissionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindStudentFormSubmissionsRequestMultiError) AllErrors() []error { return m }

// FindStudentFormSubmissionsRequestValidationError is the validation error
// returned by FindStudentFormSubmissionsRequest.Validate if the designated
// constraints aren't met.
type FindStudentFormSubmissionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindStudentFormSubmissionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindStudentFormSubmissionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindStudentFormSubmissionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindStudentFormSubmissionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindStudentFormSubmissionsRequestValidationError) ErrorName() string {
	return "FindStudentFormSubmissionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FindStudentFormSubmissionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindStudentFormSubmissionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindStudentFormSubmissionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindStudentFormSubmissionsRequestValidationError{}

// Validate checks the field values on FindStudentFormSubmissionsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *FindStudentFormSubmissionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindStudentFormSubmissionsResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// FindStudentFormSubmissionsResponseMultiError, or nil if none found.
func (m *FindStudentFormSubmissionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FindStudentFormSubmissionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StudentUsername

	// no validation rules for FormId

	for idx, item := range m.GetQuestionSubmissions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FindStudentFormSubmissionsResponseValidationError{
						field:  fmt.Sprintf("QuestionSubmissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FindStudentFormSubmissionsResponseValidationError{
						field:  fmt.Sprintf("QuestionSubmissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindStudentFormSubmissionsResponseValidationError{
					field:  fmt.Sprintf("QuestionSubmissions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FindStudentFormSubmissionsResponseMultiError(errors)
	}

	return nil
}

// FindStudentFormSubmissionsResponseMultiError is an error wrapping multiple
// validation errors returned by
// FindStudentFormSubmissionsResponse.ValidateAll() if the designated
// constraints aren't met.
type FindStudentFormSubmissionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindStudentFormSubmissionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindStudentFormSubmissionsResponseMultiError) AllErrors() []error { return m }

// FindStudentFormSubmissionsResponseValidationError is the validation error
// returned by FindStudentFormSubmissionsResponse.Validate if the designated
// constraints aren't met.
type FindStudentFormSubmissionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindStudentFormSubmissionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindStudentFormSubmissionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindStudentFormSubmissionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindStudentFormSubmissionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindStudentFormSubmissionsResponseValidationError) ErrorName() string {
	return "FindStudentFormSubmissionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FindStudentFormSubmissionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindStudentFormSubmissionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindStudentFormSubmissionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindStudentFormSubmissionsResponseValidationError{}

// Validate checks the field values on QuestionSubmission with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuestionSubmission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuestionSubmission with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuestionSubmissionMultiError, or nil if none found.
func (m *QuestionSubmission) ValidateAll() error {
	return m.validate(true)
}

func (m *QuestionSubmission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuestionId

	if len(errors) > 0 {
		return QuestionSubmissionMultiError(errors)
	}

	return nil
}

// QuestionSubmissionMultiError is an error wrapping multiple validation errors
// returned by QuestionSubmission.ValidateAll() if the designated constraints
// aren't met.
type QuestionSubmissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuestionSubmissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuestionSubmissionMultiError) AllErrors() []error { return m }

// QuestionSubmissionValidationError is the validation error returned by
// QuestionSubmission.Validate if the designated constraints aren't met.
type QuestionSubmissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionSubmissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionSubmissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionSubmissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionSubmissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionSubmissionValidationError) ErrorName() string {
	return "QuestionSubmissionValidationError"
}

// Error satisfies the builtin error interface
func (e QuestionSubmissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestionSubmission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionSubmissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionSubmissionValidationError{}
