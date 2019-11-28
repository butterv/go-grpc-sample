// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/v1/error/error.proto

package error

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ErrorCode with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ErrorCode) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ErrorCode

	return nil
}

// ErrorCodeValidationError is the validation error returned by
// ErrorCode.Validate if the designated constraints aren't met.
type ErrorCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorCodeValidationError) ErrorName() string { return "ErrorCodeValidationError" }

// Error satisfies the builtin error interface
func (e ErrorCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorCodeValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Error) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ErrorValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

// Validate checks the field values on Error_ErrorDetail with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Error_ErrorDetail) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ErrorCode

	// no validation rules for Locale

	// no validation rules for Message

	return nil
}

// Error_ErrorDetailValidationError is the validation error returned by
// Error_ErrorDetail.Validate if the designated constraints aren't met.
type Error_ErrorDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Error_ErrorDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Error_ErrorDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Error_ErrorDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Error_ErrorDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Error_ErrorDetailValidationError) ErrorName() string {
	return "Error_ErrorDetailValidationError"
}

// Error satisfies the builtin error interface
func (e Error_ErrorDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError_ErrorDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Error_ErrorDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Error_ErrorDetailValidationError{}
