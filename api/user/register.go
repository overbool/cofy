package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/model"
	"github.com/overbool/cofy/pkg/errno"
)

func Register(c *gin.Context) {
	register := func() *errno.Errno {
		var r model.User
		if err := c.BindJSON(&r); err != nil {
			return errno.ErrBind
		}

		u := model.User{
			Username: r.Username,
			Password: r.Password,
			Email:    r.Email,
		}

		if err := u.Encrypt(); err != nil {
			return errno.ErrEncrypt
		}

		if err := u.Create(); err != nil {
			return errno.ErrDatabase
		}

		return nil
	}

	if err := register(); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "register successfully",
	})
}
