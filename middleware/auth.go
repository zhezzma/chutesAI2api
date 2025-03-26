package middleware

import (
	"chutesai2api/common/config"
	"chutesai2api/model"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strings"
)

func isValidSecret(secret string) bool {
	if config.ApiSecret == "" {
		return true
	} else {
		return !lo.Contains(config.ApiSecrets, secret)
	}
}
func authHelperForOpenai(c *gin.Context) {
	secret := c.Request.Header.Get("Authorization")
	secret = strings.Replace(secret, "Bearer ", "", 1)

	b := isValidSecret(secret)

	if !b {
		c.JSON(http.StatusUnauthorized, model.OpenAIErrorResponse{
			OpenAIError: model.OpenAIError{
				Message: "API-KEY校验失败",
				Type:    "invalid_request_error",
				Code:    "invalid_authorization",
			},
		})
		c.Abort()
		return
	}

	if config.ApiSecret == "" {
		c.Request.Header.Set("Authorization", "")
	}

	c.Next()
	return
}

func OpenAIAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHelperForOpenai(c)
	}
}
