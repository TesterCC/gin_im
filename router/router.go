package router

import (
	"gin_im/middlewares"
	"gin_im/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// user login
	r.POST("/login", service.Login)

	auth := r.Group("/u", middlewares.AuthCheck())

	// user detail
	auth.GET("/user/detail", service.UserDetail)

	return r
}
