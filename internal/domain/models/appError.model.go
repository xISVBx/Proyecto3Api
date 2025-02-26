package models

import (
	"errors"
	"net/http"
)

type AppError struct {
	Err    error
	Status int
}

// Error implements error.
func (a *AppError) Error() string {
	panic("unimplemented")
}

func NewServerError(err error) *AppError {
	return &AppError{
		Err:    err,
		Status: http.StatusInternalServerError,
	}
}

func CreateError(err string) *AppError {
	return &AppError{
		Err:    errors.New(err),
		Status: http.StatusBadRequest,
	}
}

func NotFound(err string) *AppError {
	return &AppError{
		Err:    errors.New(err),
		Status: http.StatusNotFound,
	}
}
