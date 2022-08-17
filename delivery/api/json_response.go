package api

import (
	"github.com/gin-gonic/gin"
)

type AppHttpResponse interface {
	Send()
}

type jsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

// Send : send http code and body response
func (j *jsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

// JsonResponseSuccessBuilder : success body response builder
func JsonResponseSuccessBuilder(c *gin.Context, data interface{}) AppHttpResponse {
	httpStatusCode, resp := NewSuccessMessage(data)
	return &jsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

// JsonResponseFailedBuilder : failed body response builder
func JsonResponseFailedBuilder(c *gin.Context, err error) AppHttpResponse {
	httpStatusCode, resp := NewErrorMessage(err)
	return &jsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}
