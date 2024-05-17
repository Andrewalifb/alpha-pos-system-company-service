package dto

import "errors"

// Store Failed Messages
const (
	MESSAGE_FAILED_CREATE_STORE = "failed to create store"
	MESSAGE_FAILED_UPDATE_STORE = "failed to update store"
	MESSAGE_FAILED_DELETE_STORE = "failed to delete store"
	MESSAGE_FAILED_GET_STORE    = "failed to get store"
)

// Store Success Messages
const (
	MESSAGE_SUCCESS_CREATE_STORE = "success create store"
	MESSAGE_SUCCESS_UPDATE_STORE = "success update store"
	MESSAGE_SUCCESS_DELETE_STORE = "success delete store"
	MESSAGE_SUCCESS_GET_STORE    = "success get store"
)

// Store Custom Errors
var (
	ErrCreateStore = errors.New(MESSAGE_FAILED_CREATE_STORE)
	ErrUpdateStore = errors.New(MESSAGE_FAILED_UPDATE_STORE)
	ErrDeleteStore = errors.New(MESSAGE_FAILED_DELETE_STORE)
	ErrGetStore    = errors.New(MESSAGE_FAILED_GET_STORE)
)
