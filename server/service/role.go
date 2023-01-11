package service

import (
	"context"
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetRoleList 角色列表
func SetRoleList(c *gin.Context) {
	in := &SetRoleListRequest{NewQueryRequest()}
	err := c.ShouldBindQuery(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		cnt  int64
		list = make([]*SetRoleListReply, 0)
	)
	err = models.GetRoleList(in.KeyWord).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		helper.Info("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	for _, v := range list {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
		v.UpdatedAt = helper.RFC3339ToNormalTime(v.UpdatedAt)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

// SetRoleUpdateAdmin 修改角色的管理员身份
func SetRoleUpdateAdmin(c *gin.Context) {
	in := new(SetRoleUpdateAdminRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	adminRoleKey := "ADMIN-" + userClaim.RoleIdentity
	err = models.RDB.Del(context.Background(), adminRoleKey).Err()
	if err != nil {
		helper.Info("[RDB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "缓存异常",
		})
		return
	}
	err = models.DB.Model(new(models.RoleBasic)).Where("identity = ?", in.Identity).
		Updates(map[string]interface{}{"is_admin": in.IsAdmin}).Error
	if err != nil {
		helper.Info("[RDB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 延时双删
	go func() {
		time.Sleep(time.Millisecond * 200)
		err = models.RDB.Del(context.Background(), adminRoleKey).Err()
		if err != nil {
			helper.Info("[RDB ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "缓存异常",
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
