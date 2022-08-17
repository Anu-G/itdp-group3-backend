package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

// Error : custom built-in Error()
func (e *AppError) Error() string {
	return fmt.Sprintf("type: %d, code:%s, err:%s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

// RequiredError : error for empty required field
func RequiredError(field string) error {
	msg := fmt.Sprintf("%s can't be empty", field)
	return &AppError{
		ErrorCode:    "X02",
		ErrorMessage: msg,
		ErrorType:    http.StatusBadRequest,
	}
}

// WrongNumberInput : error for non-number input
func WrongNumberInput(field string) error {
	msg := fmt.Sprintf("%s must be a number", field)
	return &AppError{
		ErrorCode:    "X03",
		ErrorMessage: msg,
		ErrorType:    http.StatusBadRequest,
	}
}

// DataNotFoundError : error for data not found
func DataNotFoundError() error {
	return &AppError{
		ErrorCode:    "X04",
		ErrorMessage: "No Data Found",
		ErrorType:    http.StatusBadRequest,
	}
}
