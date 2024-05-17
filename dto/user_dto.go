package dto

import "errors"

// User Failed Messages
const (
	MESSAGE_FAILED_CREATE_USER = "failed to create user"
	MESSAGE_FAILED_UPDATE_USER = "failed to update user"
	MESSAGE_FAILED_DELETE_USER = "failed to delete user"
	MESSAGE_FAILED_GET_USER    = "failed to get user"
)

// User Success Messages
const (
	MESSAGE_SUCCESS_CREATE_USER = "success create user"
	MESSAGE_SUCCESS_UPDATE_USER = "success update user"
	MESSAGE_SUCCESS_DELETE_USER = "success delete user"
	MESSAGE_SUCCESS_GET_USER    = "success get user"
)

// User Custom Errors
var (
	ErrCreateUser = errors.New(MESSAGE_FAILED_CREATE_USER)
	ErrUpdateUser = errors.New(MESSAGE_FAILED_UPDATE_USER)
	ErrDeleteUser = errors.New(MESSAGE_FAILED_DELETE_USER)
	ErrGetUser    = errors.New(MESSAGE_FAILED_GET_USER)
)
