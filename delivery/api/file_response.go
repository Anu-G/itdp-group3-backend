package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type fileResponse struct {
	c              *gin.Context
	httpStatusCode int
	fileName       string
}

// Send : send http code and body response
func (f *fileResponse) Send() {
	f.c.File(f.fileName)
}

func FileResponseSuccessBuilder(c *gin.Context, fileName string) AppHttpResponse {
	return &fileResponse{
		c,
		http.StatusOK,
		fileName,
	}
}