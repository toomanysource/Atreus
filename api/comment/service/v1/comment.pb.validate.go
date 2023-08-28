// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: comment/service/v1/comment.proto

package v1

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

// Validate checks the field values on CommentListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommentListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommentListRequestMultiError, or nil if none found.
func (m *CommentListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for VideoId

	if len(errors) > 0 {
		return CommentListRequestMultiError(errors)
	}

	return nil
}

// CommentListRequestMultiError is an error wrapping multiple validation errors
// returned by CommentListRequest.ValidateAll() if the designated constraints
// aren't met.
type CommentListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentListRequestMultiError) AllErrors() []error { return m }

// CommentListRequestValidationError is the validation error returned by
// CommentListRequest.Validate if the designated constraints aren't met.
type CommentListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentListRequestValidationError) ErrorName() string {
	return "CommentListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CommentListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentListRequestValidationError{}

// Validate checks the field values on CommentListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CommentListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommentListReplyMultiError, or nil if none found.
func (m *CommentListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for StatusMsg

	for idx, item := range m.GetCommentList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommentListReplyValidationError{
						field:  fmt.Sprintf("CommentList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommentListReplyValidationError{
						field:  fmt.Sprintf("CommentList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommentListReplyValidationError{
					field:  fmt.Sprintf("CommentList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CommentListReplyMultiError(errors)
	}

	return nil
}

// CommentListReplyMultiError is an error wrapping multiple validation errors
// returned by CommentListReply.ValidateAll() if the designated constraints
// aren't met.
type CommentListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentListReplyMultiError) AllErrors() []error { return m }

// CommentListReplyValidationError is the validation error returned by
// CommentListReply.Validate if the designated constraints aren't met.
type CommentListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentListReplyValidationError) ErrorName() string { return "CommentListReplyValidationError" }

// Error satisfies the builtin error interface
func (e CommentListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentListReplyValidationError{}

// Validate checks the field values on CommentActionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommentActionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentActionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommentActionRequestMultiError, or nil if none found.
func (m *CommentActionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentActionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetToken()) < 1 {
		err := CommentActionRequestValidationError{
			field:  "Token",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for VideoId

	// no validation rules for ActionType

	// no validation rules for CommentText

	// no validation rules for CommentId

	if len(errors) > 0 {
		return CommentActionRequestMultiError(errors)
	}

	return nil
}

// CommentActionRequestMultiError is an error wrapping multiple validation
// errors returned by CommentActionRequest.ValidateAll() if the designated
// constraints aren't met.
type CommentActionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentActionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentActionRequestMultiError) AllErrors() []error { return m }

// CommentActionRequestValidationError is the validation error returned by
// CommentActionRequest.Validate if the designated constraints aren't met.
type CommentActionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentActionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentActionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentActionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentActionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentActionRequestValidationError) ErrorName() string {
	return "CommentActionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CommentActionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentActionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentActionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentActionRequestValidationError{}

// Validate checks the field values on CommentActionReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommentActionReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentActionReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommentActionReplyMultiError, or nil if none found.
func (m *CommentActionReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentActionReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for StatusMsg

	if all {
		switch v := interface{}(m.GetComment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommentActionReplyValidationError{
					field:  "Comment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommentActionReplyValidationError{
					field:  "Comment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetComment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommentActionReplyValidationError{
				field:  "Comment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CommentActionReplyMultiError(errors)
	}

	return nil
}

// CommentActionReplyMultiError is an error wrapping multiple validation errors
// returned by CommentActionReply.ValidateAll() if the designated constraints
// aren't met.
type CommentActionReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentActionReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentActionReplyMultiError) AllErrors() []error { return m }

// CommentActionReplyValidationError is the validation error returned by
// CommentActionReply.Validate if the designated constraints aren't met.
type CommentActionReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentActionReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentActionReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentActionReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentActionReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentActionReplyValidationError) ErrorName() string {
	return "CommentActionReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CommentActionReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentActionReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentActionReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentActionReplyValidationError{}

// Validate checks the field values on Comment with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Comment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Comment with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CommentMultiError, or nil if none found.
func (m *Comment) ValidateAll() error {
	return m.validate(true)
}

func (m *Comment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommentValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommentValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommentValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Content

	// no validation rules for CreateDate

	if len(errors) > 0 {
		return CommentMultiError(errors)
	}

	return nil
}

// CommentMultiError is an error wrapping multiple validation errors returned
// by Comment.ValidateAll() if the designated constraints aren't met.
type CommentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentMultiError) AllErrors() []error { return m }

// CommentValidationError is the validation error returned by Comment.Validate
// if the designated constraints aren't met.
type CommentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentValidationError) ErrorName() string { return "CommentValidationError" }

// Error satisfies the builtin error interface
func (e CommentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for FollowCount

	// no validation rules for FollowerCount

	// no validation rules for IsFollow

	// no validation rules for Avatar

	// no validation rules for BackgroundImage

	// no validation rules for Signature

	// no validation rules for TotalFavorited

	// no validation rules for WorkCount

	// no validation rules for FavoriteCount

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}
