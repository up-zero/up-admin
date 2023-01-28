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
			adminRoleKey := define.RedisRoleAdminPrefix + userClaim.RoleIdentity
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

// FuncAuthCheck 功能权限验证
func FuncAuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不是超管时，判断用户是否具有相关权限
		userClaim := c.MustGet("UserClaim").(*define.UserClaim)
		// 可操作函数的 key
		funcKey := define.RedisFuncPrefix + userClaim.RoleIdentity
		if !userClaim.IsAdmin {
			// 判断key是否在Redis中存在
			keyExist, _ := models.RDB.Exists(context.Background(), funcKey).Result()
			if keyExist > 0 {
				// key存在，再续一周
				models.RDB.Expire(context.Background(), funcKey, time.Second*3600*24*7)

				fieldExist, _ := models.RDB.HExists(context.Background(), funcKey, c.Request.RequestURI).Result()
				if !fieldExist { // 权限不存在，非法访问
					c.Abort()
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "非法请求",
					})
				}
			} else {
				// key不存在，从DB中查询数据，并保存
				data, err := models.GetAuthFuncUri(userClaim.RoleIdentity)
				if err != nil {
					helper.Error("[DB ERROR] : %v", err)
					c.Abort()
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "网络异常",
					})
				}
				if len(data) == 0 {
					data["up-admin-empty"] = "get"
				}
				err = models.RDB.HSet(context.Background(), funcKey, data).Err()
				if err != nil {
					helper.Error("[DB ERROR] : %v", err)
					c.Abort()
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "网络异常",
					})
				}
				models.RDB.Expire(context.Background(), funcKey, time.Second*3600*24*7)
			}
		}
		c.Next()
	}
}
