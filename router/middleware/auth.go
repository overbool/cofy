package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/api/user"
	"github.com/overbool/cofy/pkg/errno"
	"github.com/overbool/cofy/pkg/token"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			c.JSON(http.StatusOK, user.Response{
				Code:    errno.ErrTokenInvalid.Code,
				Message: errno.ErrTokenInvalid.Message,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
