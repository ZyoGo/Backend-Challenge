package http

import (
	"errors"
	"net/http"

	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/ZyoGo/Backend-Challange/pkg/http/middleware/logger"
	"go.uber.org/zap"
)

// status error response
const (
	ErrBadRequest     = "BAD_REQUEST"
	ErrValidation     = "VALIDATION_ERROR"
	ErrInternalServer = "SERVER_ERROR"
	ErrNotFound       = "NOT_FOUND"
	ErrUnauthorized   = "UNAUTHORIZED"
)

type DefaultResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code     int    `json:"code"`
	Status   string `json:"status,omitempty"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		Code:    400,
		Status:  ErrBadRequest,
		Message: "Bad Request",
	}
}

func NewValidationResponse(errMsg string) DefaultResponse {
	return DefaultResponse{
		Code:    400,
		Status:  ErrValidation,
		Message: errMsg,
	}
}

func NewUnauthorizedResponse(msg string) DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusUnauthorized,
		Status:  ErrUnauthorized,
		Message: msg,
	}
}

func MapErrorToResponse(err error) ErrorResponse {
	defaultErr := ErrorResponse{Code: http.StatusInternalServerError, Status: ErrInternalServer, Message: "Internal Server Error"}
	var ierr *derrors.Error
	if !errors.As(err, &ierr) {
		logger.GetLogger().Debug("Internal Error", zap.String("error", err.Error()))
		return defaultErr
	} else {
		logger.GetLogger().Debug("Internal Error", zap.String("error", ierr.Error()))
		cases := map[derrors.ErrorCode]ErrorResponse{
			derrors.ErrorCodeBadRequest:   {Code: http.StatusBadRequest, Status: ErrBadRequest, Message: ierr.Error(), Internal: ierr},
			derrors.ErrorCodeNotFound:     {Code: http.StatusNotFound, Status: ErrNotFound, Message: ierr.Error(), Internal: ierr},
			derrors.ErrorCodeUnauthorized: {Code: http.StatusUnauthorized, Status: ErrUnauthorized, Message: "Unauthorized", Internal: ierr},
		}

		catchErr, found := cases[ierr.Code()]
		if found {
			return catchErr
		}
	}
	return defaultErr
}
