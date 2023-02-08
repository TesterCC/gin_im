package main

import (
	"fmt"
	"gin_im/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	eng := router.Router()

	fmt.Println("[*] current mode: ", gin.Mode())
	if gin.Mode() == "debug" {
		eng.Use(cors.Default()) // 在 debug 模式下, 允许跨域访问
	}

	eng.Run(":8080")
}