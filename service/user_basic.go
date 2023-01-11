package service

import (
	"gin_im/helper"
	"gin_im/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	account := c.PostForm("account")
	password := c.PostForm("password")

	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 20,
			"msg": "args empty",
			"data": nil,
		})
		return
	}

	ub, err := models.GerUserBasicByAccountPassword(account, helper.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 11,
			"msg": "authentication failed",
			"data": nil,
		})
		return
	}

	token, err := helper.GenerateToken(ub.Id, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 97,
			"msg": "system error: " + err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "login success",
		"data": gin.H{
			"token": token,
		},
	})

}