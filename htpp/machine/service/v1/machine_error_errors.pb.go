// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_NOT_FOUND_ERROR.String() && e.Code == 500
}

func ErrorNotFoundError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_NOT_FOUND_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsCreateFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_CREATE_FAILED.String() && e.Code == 500
}

func ErrorCreateFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_CREATE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsAddressConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_ADDRESS_CONFLICT.String() && e.Code == 500
}

func ErrorAddressConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_ADDRESS_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsUuidGenerateFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_UUID_GENERATE_FAILED.String() && e.Code == 500
}

func ErrorUuidGenerateFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_UUID_GENERATE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUuidParseFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_UUID_PARSE_FAILED.String() && e.Code == 500
}

func ErrorUuidParseFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_UUID_PARSE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCronConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_CRON_CONFLICT.String() && e.Code == 500
}

func ErrorCronConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_CRON_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsCronSetupFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MachineServiceErrorReason_CRON_SETUP_FAILED.String() && e.Code == 500
}

func ErrorCronSetupFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MachineServiceErrorReason_CRON_SETUP_FAILED.String(), fmt.Sprintf(format, args...))
}
