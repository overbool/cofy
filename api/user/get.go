package user

import (
	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/model"
	"github.com/overbool/cofy/pkg/errno"
	"net/http"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    errno.ErrUserNotFound.Code,
			Message: errno.ErrUserNotFound.Message,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Message: "get user successfully",
		Data: user,
	})
}
