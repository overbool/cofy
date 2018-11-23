package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/overbool/cofy/api/user"
	"github.com/overbool/cofy/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)

	// pprof router
	pprof.Register(g)

	// api for authentication
	g.POST("/v1/login", user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Register)
		u.DELETE("/:id", user.Delete)
		u.GET("/:username", user.Get)
	}

	return g
}
