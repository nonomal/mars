// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: metrics/metrics.proto

package metrics

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

// Validate checks the field values on MetricsTopPodRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetricsTopPodRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsTopPodRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetricsTopPodRequestMultiError, or nil if none found.
func (m *MetricsTopPodRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsTopPodRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetNamespace()) < 1 {
		err := MetricsTopPodRequestValidationError{
			field:  "Namespace",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPod()) < 1 {
		err := MetricsTopPodRequestValidationError{
			field:  "Pod",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetricsTopPodRequestMultiError(errors)
	}

	return nil
}

// MetricsTopPodRequestMultiError is an error wrapping multiple validation
// errors returned by MetricsTopPodRequest.ValidateAll() if the designated
// constraints aren't met.
type MetricsTopPodRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsTopPodRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsTopPodRequestMultiError) AllErrors() []error { return m }

// MetricsTopPodRequestValidationError is the validation error returned by
// MetricsTopPodRequest.Validate if the designated constraints aren't met.
type MetricsTopPodRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsTopPodRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsTopPodRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsTopPodRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsTopPodRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsTopPodRequestValidationError) ErrorName() string {
	return "MetricsTopPodRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsTopPodRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsTopPodRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsTopPodRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsTopPodRequestValidationError{}

// Validate checks the field values on MetricsTopPodResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetricsTopPodResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsTopPodResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetricsTopPodResponseMultiError, or nil if none found.
func (m *MetricsTopPodResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsTopPodResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Cpu

	// no validation rules for Memory

	// no validation rules for HumanizeCpu

	// no validation rules for HumanizeMemory

	// no validation rules for Time

	// no validation rules for Length

	if len(errors) > 0 {
		return MetricsTopPodResponseMultiError(errors)
	}

	return nil
}

// MetricsTopPodResponseMultiError is an error wrapping multiple validation
// errors returned by MetricsTopPodResponse.ValidateAll() if the designated
// constraints aren't met.
type MetricsTopPodResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsTopPodResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsTopPodResponseMultiError) AllErrors() []error { return m }

// MetricsTopPodResponseValidationError is the validation error returned by
// MetricsTopPodResponse.Validate if the designated constraints aren't met.
type MetricsTopPodResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsTopPodResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsTopPodResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsTopPodResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsTopPodResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsTopPodResponseValidationError) ErrorName() string {
	return "MetricsTopPodResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsTopPodResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsTopPodResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsTopPodResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsTopPodResponseValidationError{}

// Validate checks the field values on MetricsCpuMemoryInNamespaceRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *MetricsCpuMemoryInNamespaceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsCpuMemoryInNamespaceRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// MetricsCpuMemoryInNamespaceRequestMultiError, or nil if none found.
func (m *MetricsCpuMemoryInNamespaceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsCpuMemoryInNamespaceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNamespaceId() <= 0 {
		err := MetricsCpuMemoryInNamespaceRequestValidationError{
			field:  "NamespaceId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetricsCpuMemoryInNamespaceRequestMultiError(errors)
	}

	return nil
}

// MetricsCpuMemoryInNamespaceRequestMultiError is an error wrapping multiple
// validation errors returned by
// MetricsCpuMemoryInNamespaceRequest.ValidateAll() if the designated
// constraints aren't met.
type MetricsCpuMemoryInNamespaceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsCpuMemoryInNamespaceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsCpuMemoryInNamespaceRequestMultiError) AllErrors() []error { return m }

// MetricsCpuMemoryInNamespaceRequestValidationError is the validation error
// returned by MetricsCpuMemoryInNamespaceRequest.Validate if the designated
// constraints aren't met.
type MetricsCpuMemoryInNamespaceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsCpuMemoryInNamespaceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsCpuMemoryInNamespaceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsCpuMemoryInNamespaceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsCpuMemoryInNamespaceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsCpuMemoryInNamespaceRequestValidationError) ErrorName() string {
	return "MetricsCpuMemoryInNamespaceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsCpuMemoryInNamespaceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsCpuMemoryInNamespaceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsCpuMemoryInNamespaceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsCpuMemoryInNamespaceRequestValidationError{}

// Validate checks the field values on MetricsCpuMemoryInNamespaceResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *MetricsCpuMemoryInNamespaceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsCpuMemoryInNamespaceResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// MetricsCpuMemoryInNamespaceResponseMultiError, or nil if none found.
func (m *MetricsCpuMemoryInNamespaceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsCpuMemoryInNamespaceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Cpu

	// no validation rules for Memory

	if len(errors) > 0 {
		return MetricsCpuMemoryInNamespaceResponseMultiError(errors)
	}

	return nil
}

// MetricsCpuMemoryInNamespaceResponseMultiError is an error wrapping multiple
// validation errors returned by
// MetricsCpuMemoryInNamespaceResponse.ValidateAll() if the designated
// constraints aren't met.
type MetricsCpuMemoryInNamespaceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsCpuMemoryInNamespaceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsCpuMemoryInNamespaceResponseMultiError) AllErrors() []error { return m }

// MetricsCpuMemoryInNamespaceResponseValidationError is the validation error
// returned by MetricsCpuMemoryInNamespaceResponse.Validate if the designated
// constraints aren't met.
type MetricsCpuMemoryInNamespaceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsCpuMemoryInNamespaceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsCpuMemoryInNamespaceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsCpuMemoryInNamespaceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsCpuMemoryInNamespaceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsCpuMemoryInNamespaceResponseValidationError) ErrorName() string {
	return "MetricsCpuMemoryInNamespaceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsCpuMemoryInNamespaceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsCpuMemoryInNamespaceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsCpuMemoryInNamespaceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsCpuMemoryInNamespaceResponseValidationError{}

// Validate checks the field values on MetricsCpuMemoryInProjectRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *MetricsCpuMemoryInProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsCpuMemoryInProjectRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// MetricsCpuMemoryInProjectRequestMultiError, or nil if none found.
func (m *MetricsCpuMemoryInProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsCpuMemoryInProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetProjectId() <= 0 {
		err := MetricsCpuMemoryInProjectRequestValidationError{
			field:  "ProjectId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetricsCpuMemoryInProjectRequestMultiError(errors)
	}

	return nil
}

// MetricsCpuMemoryInProjectRequestMultiError is an error wrapping multiple
// validation errors returned by
// MetricsCpuMemoryInProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type MetricsCpuMemoryInProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsCpuMemoryInProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsCpuMemoryInProjectRequestMultiError) AllErrors() []error { return m }

// MetricsCpuMemoryInProjectRequestValidationError is the validation error
// returned by MetricsCpuMemoryInProjectRequest.Validate if the designated
// constraints aren't met.
type MetricsCpuMemoryInProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsCpuMemoryInProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsCpuMemoryInProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsCpuMemoryInProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsCpuMemoryInProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsCpuMemoryInProjectRequestValidationError) ErrorName() string {
	return "MetricsCpuMemoryInProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsCpuMemoryInProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsCpuMemoryInProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsCpuMemoryInProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsCpuMemoryInProjectRequestValidationError{}

// Validate checks the field values on MetricsCpuMemoryInProjectResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *MetricsCpuMemoryInProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricsCpuMemoryInProjectResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// MetricsCpuMemoryInProjectResponseMultiError, or nil if none found.
func (m *MetricsCpuMemoryInProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricsCpuMemoryInProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Cpu

	// no validation rules for Memory

	if len(errors) > 0 {
		return MetricsCpuMemoryInProjectResponseMultiError(errors)
	}

	return nil
}

// MetricsCpuMemoryInProjectResponseMultiError is an error wrapping multiple
// validation errors returned by
// MetricsCpuMemoryInProjectResponse.ValidateAll() if the designated
// constraints aren't met.
type MetricsCpuMemoryInProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsCpuMemoryInProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsCpuMemoryInProjectResponseMultiError) AllErrors() []error { return m }

// MetricsCpuMemoryInProjectResponseValidationError is the validation error
// returned by MetricsCpuMemoryInProjectResponse.Validate if the designated
// constraints aren't met.
type MetricsCpuMemoryInProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsCpuMemoryInProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsCpuMemoryInProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsCpuMemoryInProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsCpuMemoryInProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsCpuMemoryInProjectResponseValidationError) ErrorName() string {
	return "MetricsCpuMemoryInProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsCpuMemoryInProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsCpuMemoryInProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsCpuMemoryInProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsCpuMemoryInProjectResponseValidationError{}
