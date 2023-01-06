package service

import (
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	data := new(UserInfoReply)

	err := models.GetUserInfo(userClaim.Identity).Find(data).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}

// UserPasswordChange 修改密码
func UserPasswordChange(c *gin.Context) {
	in := new(UserPasswordChangeRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)

	// 判断旧密码是否正确
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).
		Where("identity = ? AND password = ?", userClaim.Identity, helper.Md5(in.OldPassword)).Count(&cnt).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}
	if cnt == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "旧密码不正确",
		})
		return
	}

	// 修改密码
	err = models.DB.Model(new(models.UserBasic)).Where("identity = ?", userClaim.Identity).
		Updates(map[string]interface{}{"password": helper.Md5(in.NewPassword)}).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "网络异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
