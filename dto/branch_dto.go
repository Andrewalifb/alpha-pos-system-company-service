package dto

import "errors"

// Branch Failed Messages
const (
	MESSAGE_FAILED_CREATE_BRANCH = "failed to create branch"
	MESSAGE_FAILED_UPDATE_BRANCH = "failed to update branch"
	MESSAGE_FAILED_DELETE_BRANCH = "failed to delete branch"
	MESSAGE_FAILED_GET_BRANCH    = "failed to get branch"
)

// Branch Success Messages
const (
	MESSAGE_SUCCESS_CREATE_BRANCH = "success create branch"
	MESSAGE_SUCCESS_UPDATE_BRANCH = "success update branch"
	MESSAGE_SUCCESS_DELETE_BRANCH = "success delete branch"
	MESSAGE_SUCCESS_GET_BRANCH    = "success get branch"
)

// Branch Custom Errors
var (
	ErrCreateBranch = errors.New(MESSAGE_FAILED_CREATE_BRANCH)
	ErrUpdateBranch = errors.New(MESSAGE_FAILED_UPDATE_BRANCH)
	ErrDeleteBranch = errors.New(MESSAGE_FAILED_DELETE_BRANCH)
	ErrGetBranch    = errors.New(MESSAGE_FAILED_GET_BRANCH)
)
