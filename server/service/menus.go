package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Menus 获取菜单列表
func Menus(c *gin.Context) {
	userClaim := c.MustGet("UserClaim").(*define.UserClaim)
	menuField := userClaim.RoleIdentity
	if userClaim.IsAdmin {
		menuField = "ADMIN"
	}

	data := make([]*MenuReply, 0)
	roleMenus := make([]*RoleMenu, 0)
	res, err := models.RDB.HGet(context.Background(), define.RedisMenuPrefix, menuField).Result()
	if err != nil {
		fmt.Println("RUN", err.Error())
		tx, err := models.GetRoleMenus(userClaim.RoleIdentity, userClaim.IsAdmin)
		if err != nil {
			helper.Error("[DB ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据异常",
			})
			return
		}
		err = tx.Find(&roleMenus).Error
		if err != nil {
			helper.Error("[DB ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据异常",
			})
			return
		}
		data = roleMenuToMenuReply(roleMenus)
		b, _ := json.Marshal(data)
		models.RDB.HSet(context.Background(), define.RedisMenuPrefix, map[string]string{
			menuField: string(b),
		})
	} else {
		err = json.Unmarshal([]byte(res), &data)
		if err != nil {
			helper.Error("[Unmarshal ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据异常",
			})
			return
		}
		// Redis 查到数据后，再续两周
		models.RDB.Expire(context.Background(), define.RedisMenuPrefix, time.Second*3600*24*14)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}

// SetMenuList 设置-获取菜单列表
func SetMenuList(c *gin.Context) {
	Menus(c)
}

func roleMenuToMenuReply(roleMenus []*RoleMenu) []*MenuReply {
	parentId := make(map[int]int)
	reply := make([]*MenuReply, 0)
	// 一层循环，得到父级菜单
	for _, v := range roleMenus {
		if v.ParentId == 0 {
			reply = append(reply, &MenuReply{
				Identity: v.Identity,
				Name:     v.Name,
				WebIcon:  v.WebIcon,
				Sort:     v.Sort,
				Path:     v.Path,
			})
			parentId[v.Id] = len(reply) - 1
		}
	}
	// 再次循环，得到子级菜单
	for _, v := range roleMenus {
		if i, ok := parentId[v.ParentId]; ok {
			reply[i].SubMenus = append(reply[i].SubMenus, struct {
				Identity string `json:"identity"`
				Name     string `json:"name"`
				WebIcon  string `json:"web_icon"`
				Sort     int    `json:"sort"`
				Path     string `json:"path"`
			}{Identity: v.Identity, Name: v.Name, WebIcon: v.WebIcon, Sort: v.Sort, Path: v.Path})
		}
	}
	return reply
}
