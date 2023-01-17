package service

import (
	"gin_im/helper"
	"gin_im/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")

	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 20,
			"msg":  "args empty",
			"data": nil,
		})
		return
	}

	ub, err := models.GerUserBasicByAccountPassword(account, helper.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 11,
			"msg":  "authentication failed",
			"data": nil,
		})
		return
	}

	token, err := helper.GenerateToken(ub.Id, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 97,
			"msg":  "system error: " + err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "login success",
		"data": gin.H{
			"token": token,
		},
	})

}

func UserDetail(c *gin.Context) {
	// query user detail data
	u, _ := c.Get("user_claims")
	uc := u.(*helper.UserClaims)

	userBasic, err := models.GetUserBasicById(uc.Id)   // uc.Id, str not ObjectiveID

	// get user basic by email   // self custom also ok
	//userBasic, err := models.GetUserBasicByEmail(uc.Email)

	if err != nil {
		log.Printf("[E] DB ERROR: %v\n", err)

		c.JSON(http.StatusOK, gin.H{
			"code": 98,
			"msg": "database query exception",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": userBasic,
	})

}
