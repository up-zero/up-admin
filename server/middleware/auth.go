package middleware

import (
	"context"
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// LoginAuthCheck 登录信息认证
func LoginAuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaim, err := helper.AnalyzeToken(c.GetHeader("AccessToken"))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "登录过期，请重新登录",
			})
		} else {
			if userClaim.RoleIdentity == "" {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "非法请求",
				})
			}

			// 判断是否是超管
			isAdmin, err := models.RDB.Get(context.Background(), define.RedisRoleAdminPrefix+userClaim.RoleIdentity).Result()
			adminRoleKey := "ADMIN-" + userClaim.RoleIdentity
			if err != nil {
				roleBasic := new(models.RoleBasic)
				err = models.DB.Model(new(models.RoleBasic)).Select("is_admin").
					Where("identity = ?", userClaim.RoleIdentity).Find(roleBasic).Error
				if err != nil {
					helper.Error("[DB ERROR] : %v", err)
					c.Abort()
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "网络异常，请重试",
					})
				}
				// 同步 Redis，默认保存一周
				models.RDB.Set(context.Background(), adminRoleKey, roleBasic.IsAdmin, time.Second*3600*24*7)
				if roleBasic.IsAdmin == 1 {
					isAdmin = "1"
				} else {
					isAdmin = "0"
				}
			} else {
				// Redis 查到数据后，再续一周
				models.RDB.Expire(context.Background(), adminRoleKey, time.Second*3600*24*7)
			}
			if isAdmin == "1" {
				userClaim.IsAdmin = true
			} else {
				userClaim.IsAdmin = false
			}
			c.Set("UserClaim", userClaim)
			c.Next()
		}
	}
}

// TODO: 功能权限验证
