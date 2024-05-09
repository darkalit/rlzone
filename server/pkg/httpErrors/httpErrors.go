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

func parseValidatorError(err error) RestErr {
	switch {
	default:
		return RestError{
			ErrStatus: http.StatusBadRequest,
			ErrError:  "Bad request",
			ErrCause:  err,
		}
	}
}

func ParseError(err error) RestErr {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return RestError{
			ErrStatus: http.StatusNotFound,
			ErrError:  "Not Found",
			ErrCause:  err,
		}
	case errors.Is(err, context.DeadlineExceeded):
		return RestError{
			ErrStatus: http.StatusRequestTimeout,
			ErrError:  "Request timeout",
			ErrCause:  err,
		}
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return RestError{
			ErrStatus: http.StatusBadRequest,
			ErrError:  "Bad request",
			ErrCause:  err,
		}
	case strings.Contains(strings.ToLower(err.Error()), "parse"):
		return RestError{
			ErrStatus: http.StatusBadRequest,
			ErrError:  err.Error(),
			ErrCause:  err,
		}
	case strings.Contains(err.Error(), "UUID"):
		return RestError{
			ErrStatus: http.StatusBadRequest,
			ErrError:  err.Error(),
			ErrCause:  err,
		}
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}

		return RestError{
			ErrStatus: http.StatusInternalServerError,
			ErrError:  "Internal server error",
			ErrCause:  err,
		}
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseError(err).Status(), ParseError(err)
}
