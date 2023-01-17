package middlewares

import (
	"gin_im/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 11,
				"msg": "authentication failed",
				"data": nil,
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()   // 以便进入下一层的service
	}

}
