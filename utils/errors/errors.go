package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	Err        string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.ErrStatus
}

func (e *apiError) Message() string {
	return e.ErrMessage
}

func (e *apiError) Error() string {
	return e.Err
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
	}
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		ErrStatus:  statusCode,
		ErrMessage: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}

	return &result, nil
}
