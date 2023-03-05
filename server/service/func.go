package service

import (
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Functions(c *gin.Context) {
	SetFuncList(c)
}

func SetFuncList(c *gin.Context) {
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	data := make([]*SetFuncListReply, 0)
	err := models.GetRoleFunctions(userClaim.RoleIdentity, userClaim.IsAdmin).Find(&data).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}
