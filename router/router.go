package router

import (
	"gin_im/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// user login
	r.POST("/login", service.Login)
	return r
}
