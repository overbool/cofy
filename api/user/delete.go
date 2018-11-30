package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/model"
	"github.com/overbool/cofy/pkg/errno"
)

func Delete(c *gin.Context) {
	del := func() *errno.Errno {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errno.ErrBind
		}

		if err := model.DeleteUser(uint(id)); err != nil {
			return errno.ErrDatabase
		}

		return nil
	}

	if err := del(); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "delete user successful",
	})
}
