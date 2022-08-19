package middleware

import (
	"fmt"
	"itdp-group3-backend/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
	OpenToken(ctx *gin.Context) string
}

type authTokenMiddleware struct {
	token auth.Token
}

func NewTokenValidator(t auth.Token) AuthTokenMiddleware {
	newValidator := new(authTokenMiddleware)
	newValidator.token = t
	return newValidator
}

func (at *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}
		fmt.Println(&h)
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			ctx.Abort()
			return
		}

		tokenStr := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": "unauthorized",
			})
			ctx.Abort()
			return
		}

		token, err := at.token.VerifyAccessToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			ctx.Abort()
			return
		}
		userId, err := at.token.FetchAccessToken(token)
		if userId == "" || err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": "unauthorized",
			})
		}

		if token != nil {
			ctx.Set("user-id", userId)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": "unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}

func (at *authTokenMiddleware) OpenToken(ctx *gin.Context) string {
	h := authHeader{}
	fmt.Println(&h)
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		ctx.Abort()
		return ""
	}

	tokenStr := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err": "unauthorized",
		})
		ctx.Abort()
		return ""
	}

	token, err := at.token.VerifyAccessToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		ctx.Abort()
		return ""
	}
	userName, err := at.token.OpenAccessToken(token)
	if userName == "" || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err": "unauthorized",
		})
	}
	return userName
}
