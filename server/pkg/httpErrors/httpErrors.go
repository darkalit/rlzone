package httpErrors

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type RestErr interface {
	Status() int
	Error() string
	Cause() interface{}
}

type RestError struct {
	ErrStatus int         `json:"Status"`
	ErrError  string      `json:"Error"`
	ErrCause  interface{} `json:"Cause"`
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - error: %s - cause: %v", e.ErrStatus, e.ErrError, e.ErrCause)
}

func (e RestError) Cause() interface{} {
	return e.ErrCause
}

func NewRestError(status int, error string, cause interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  error,
		ErrCause:  cause,
	}
}

func NewRestErrorMessage(status int, error string) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  error,
		ErrCause:  struct{}{},
	}
}

func parseValidatorError(err error) RestErr {
	switch {
	default:
		return NewRestError(http.StatusBadRequest, "Bad request", err)
	}
}

func ParseError(err error) RestErr {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewRestError(http.StatusNotFound, "Not Found", err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, "Request Timeout", err)
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, "Bad Request", err)
	case strings.Contains(strings.ToLower(err.Error()), "parse"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(err.Error(), "UUID"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}

		return NewRestError(http.StatusInternalServerError, "Internal server error", err)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseError(err).Status(), ParseError(err)
}
