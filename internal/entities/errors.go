package entities

import (
	"errors"
	"net/http"
)

type ServiceError interface {
	Code() int
	Error() string
	Message() string
}

type BadRequestError struct {
	code    int
	err     error
	message string
}

func (e BadRequestError) Code() int {
	return e.code
}

func (e BadRequestError) Error() string {
	return e.err.Error()
}

func (e BadRequestError) Message() string {
	return e.message
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		code:    http.StatusBadRequest,
		err:     errors.New(message),
		message: message,
	}
}

type InternalError struct {
	code    int
	err     error
	message string
}

func (e InternalError) Code() int {
	return e.code
}

func (e InternalError) Error() string {
	return e.err.Error()
}

func (e InternalError) Message() string {
	return e.message
}

func NewInternalError(e error) InternalError {
	return InternalError{
		code:    http.StatusInternalServerError,
		err:     e,
		message: "internal error",
	}
}
