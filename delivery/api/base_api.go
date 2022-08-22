package api

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// ParseBodyRequest : get JSON body request
func (b *BaseApi) ParseBodyRequest(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

// SuccessResponse : send success JSON response
func (b *BaseApi) SuccessResponse(c *gin.Context, data interface{}) {
	JsonResponseSuccessBuilder(c, data).Send()
}

// FailedResponse :  send failed JSON response
func (b *BaseApi) FailedResponse(c *gin.Context, err error) {
	JsonResponseFailedBuilder(c, err).Send()
}

func (b *BaseApi) SuccessDownload(c *gin.Context, filePath string) {
	FileResponseSuccessBuilder(c, filePath).Send()
}

