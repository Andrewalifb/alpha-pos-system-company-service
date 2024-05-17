package dto

import "errors"

// Role Failed Messages
const (
	MESSAGE_FAILED_CREATE_ROLE = "failed to create role"
	MESSAGE_FAILED_UPDATE_ROLE = "failed to update role"
	MESSAGE_FAILED_DELETE_ROLE = "failed to delete role"
	MESSAGE_FAILED_GET_ROLE    = "failed to get role"
)

// Role Success Messages
const (
	MESSAGE_SUCCESS_CREATE_ROLE = "success create role"
	MESSAGE_SUCCESS_UPDATE_ROLE = "success update role"
	MESSAGE_SUCCESS_DELETE_ROLE = "success delete role"
	MESSAGE_SUCCESS_GET_ROLE    = "success get role"
)

// Role Custom Errors
var (
	ErrCreateRole = errors.New(MESSAGE_FAILED_CREATE_ROLE)
	ErrUpdateRole = errors.New(MESSAGE_FAILED_UPDATE_ROLE)
	ErrDeleteRole = errors.New(MESSAGE_FAILED_DELETE_ROLE)
	ErrGetRole    = errors.New(MESSAGE_FAILED_GET_ROLE)
)
