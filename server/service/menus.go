package service

import (
	"context"
	"encoding/json"
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
	menuKey := define.RedisMenuPrefix + userClaim.RoleIdentity
	if userClaim.IsAdmin {
		menuKey = define.RedisMenuPrefix + "ADMIN"
	}

	data := make([]*MenuReply, 0)
	roleMenus := make([]*RoleMenu, 0)
	res, err := models.RDB.Get(context.Background(), menuKey).Result()
	if err != nil {
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
		// 同步 Redis，默认保存两周
		models.RDB.Set(context.Background(), menuKey, string(b), time.Second*3600*24*14)
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
		models.RDB.Expire(context.Background(), menuKey, time.Second*3600*24*14)
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
			}{Identity: v.Identity, Name: v.Name, WebIcon: v.WebIcon, Sort: v.Sort})
		}
	}
	return reply
}
