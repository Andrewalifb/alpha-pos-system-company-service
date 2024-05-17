package dto

import "errors"

// Company Failed Messages
const (
	MESSAGE_FAILED_CREATE_COMPANY = "failed to create company"
	MESSAGE_FAILED_UPDATE_COMPANY = "failed to update company"
	MESSAGE_FAILED_DELETE_COMPANY = "failed to delete company"
	MESSAGE_FAILED_GET_COMPANY    = "failed to get company"
)

// Company Success Messages
const (
	MESSAGE_SUCCESS_CREATE_COMPANY = "success create company"
	MESSAGE_SUCCESS_UPDATE_COMPANY = "success update company"
	MESSAGE_SUCCESS_DELETE_COMPANY = "success delete company"
	MESSAGE_SUCCESS_GET_COMPANY    = "success get company"
)

// Company Custom Errors
var (
	ErrCreateCompany = errors.New(MESSAGE_FAILED_CREATE_COMPANY)
	ErrUpdateCompany = errors.New(MESSAGE_FAILED_UPDATE_COMPANY)
	ErrDeleteCompany = errors.New(MESSAGE_FAILED_DELETE_COMPANY)
	ErrGetCompany    = errors.New(MESSAGE_FAILED_GET_COMPANY)
)
