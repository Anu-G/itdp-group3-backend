package api

import (
	"errors"
	"itdp-group3-backend/utils"
	"net/http"
)

type Status struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

// NewSuccessMessage : return success body response
func NewSuccessMessage(data interface{}) (httpStatusCode int, apiResponse Response) {
	status := Status{
		ResponseCode:    "00",
		ResponseMessage: "success",
	}
	httpStatusCode = http.StatusOK
	apiResponse = Response{
		status, data,
	}
	return
}

// NewErrorMessage : return error body response
func NewErrorMessage(err error) (httpStatusCode int, apiResponse Response) {
	var userError *utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponseCode:    "X01",
			ResponseMessage: err.Error(),
		}
		httpStatusCode = http.StatusBadRequest
	}
	apiResponse = Response{
		status, nil,
	}

	return
}
