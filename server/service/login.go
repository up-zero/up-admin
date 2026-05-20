package service

import (
	"errors"
	"net/http"

	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/logger"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPassword 用户名密码登录
func LoginPassword(c *gin.Context) {
	in := new(LoginPasswordRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		logger.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	// 根据账号、密码查询用户
	ub, err := models.GetUserBasicByUsernamePassword(in.Username, helper.Md5(in.Password))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
	}

	// 生成 token
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Username, ub.RoleIdentity, define.TokenExpireAt())
	if err != nil {
		logger.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// 生成 token
	refreshToken, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Username, ub.RoleIdentity, define.RefreshTokenExpireAt())
	if err != nil {
		logger.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	data := &LoginPasswordReply{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": data,
	})
}
