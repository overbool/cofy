package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/model"
	"github.com/overbool/cofy/pkg/errno"
	"github.com/overbool/cofy/pkg/token"
)

func Login(c *gin.Context) {
	login := func() (string, *errno.Errno) {
		var r model.User
		if err := c.BindJSON(&r); err != nil {
			return "", errno.ErrBind
		}

		u, err := model.GetUser(r.Username)
		if err != nil {
			return "", errno.ErrUserNotFound
		}

		if err := u.Compare(r.Password); err != nil {
			return "", errno.ErrPasswordIncorrect
		}

		// Sign the json web token.
		t, err := token.Sign(c, token.Context{ID: u.ID, Username: u.Username}, "")
		if err != nil {
			return "", errno.ErrToken
		}

		return t, nil
	}

	if t, err := login(); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    err.Code,
			Message: err.Message,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code: 0,
			Data: t,
		})
	}

}
