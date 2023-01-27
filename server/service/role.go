package service

import (
	"context"
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	adminRoleKey := define.RedisRoleAdminPrefix + userClaim.RoleIdentity
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

// SetRoleCreate 角色创建
func SetRoleCreate(c *gin.Context) {
	in := new(SetRoleCreateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		menuIds = make([]uint, 0)
		funcIds = make([]uint, 0)
	)
	// 获取菜单ID
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity IN ?", in.MenuIdentities).Scan(&menuIds).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 获取功能ID
	err = models.DB.Model(new(models.FunctionBasic)).Select("id").
		Where("identity IN ?", in.FuncIdentities).Scan(&funcIds).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 组装数据
	// 角色
	rb := &models.RoleBasic{
		Identity: helper.UUID(),
		Name:     in.Name,
		IsAdmin:  in.IsAdmin,
		Sort:     in.Sort,
	}
	// 授权的菜单
	rms := make([]*models.RoleMenu, len(menuIds))
	for i, _ := range rms {
		rms[i] = &models.RoleMenu{
			MenuId: menuIds[i],
		}
	}
	// 授权的功能
	rfs := make([]*models.RoleFunction, len(funcIds))
	for i, _ := range rfs {
		rfs[i] = &models.RoleFunction{
			FunctionId: funcIds[i],
		}
	}
	// 新增数据
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 角色
		err = tx.Create(rb).Error
		if err != nil {
			return err
		}
		// 授权的菜单
		for _, v := range rms {
			v.RoleId = rb.ID
		}
		if len(rms) > 0 {
			err = tx.Create(rms).Error
			if err != nil {
				return err
			}
		}
		// 授权的功能
		for _, v := range rfs {
			v.RoleId = rb.ID
		}
		if len(rfs) > 0 {
			err = tx.Create(rfs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
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
		"msg":  "创建成功",
	})
}

// SetRoleDelete 角色删除
func SetRoleDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 判断当前角色是否已关联用户
	var cnt int64
	err := models.DB.Model(new(models.UserBasic)).Where("role_identity = ?", identity).Count(&cnt).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前角色存在关联的用户，不可删除",
		})
		return
	}
	// 删除角色
	err = models.DB.Where("identity = ?", identity).Delete(new(models.RoleBasic)).Error
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
		"msg":  "删除成功",
	})
}

// SetRoleUpdate 角色修改
func SetRoleUpdate(c *gin.Context) {
	in := new(SetRoleUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		menuIds = make([]uint, 0)
		funcIds = make([]uint, 0)
		rb      = new(models.RoleBasic)
	)
	// 获取菜单ID
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity IN ?", in.MenuIdentities).Scan(&menuIds).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 获取功能ID
	err = models.DB.Model(new(models.FunctionBasic)).Select("id").
		Where("identity IN ?", in.FuncIdentities).Scan(&funcIds).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 角色ID
	err = models.DB.Model(new(models.RoleBasic)).Select("id").
		Where("identity = ?", in.Identity).Find(rb).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 组装数据
	// 授权的菜单
	rms := make([]*models.RoleMenu, len(menuIds))
	for i, _ := range rms {
		rms[i] = &models.RoleMenu{
			RoleId: rb.ID,
			MenuId: menuIds[i],
		}
	}
	// 授权的功能
	rfs := make([]*models.RoleFunction, len(funcIds))
	for i, _ := range rfs {
		rfs[i] = &models.RoleFunction{
			RoleId:     rb.ID,
			FunctionId: funcIds[i],
		}
	}
	// Redis Key 删除
	err = redisRoleDelete(in.Identity)
	if err != nil {
		helper.Info("[RDB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "缓存异常",
		})
		return
	}
	// 修改数据
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 更新角色
		err = tx.Model(new(models.RoleBasic)).Where("id = ?", rb.ID).Updates(map[string]interface{}{
			"name":     in.Name,
			"is_admin": in.IsAdmin,
			"sort":     in.Sort,
		}).Error
		if err != nil {
			return err
		}
		// 删除老数据
		// 授权的菜单
		err = tx.Where("role_id = ?", rb.ID).Delete(new(models.RoleMenu)).Error
		if err != nil {
			return err
		}
		// 授权的功能
		err = tx.Where("role_id = ?", rb.ID).Delete(new(models.RoleFunction)).Error
		if err != nil {
			return err
		}
		// 增加新数据
		// 授权的菜单
		if len(rms) > 0 {
			err = tx.Create(rms).Error
			if err != nil {
				return err
			}
		}
		// 授权的功能
		if len(rfs) > 0 {
			err = tx.Create(rfs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 延迟双删
	go func() {
		time.Sleep(time.Millisecond * 200)
		err = redisRoleDelete(in.Identity)
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

// redisRoleDelete 删除角色相关的缓存数据
func redisRoleDelete(roleIdentity string) error {
	err := models.RDB.Del(context.Background(), define.RedisRoleAdminPrefix+roleIdentity).Err()
	if err != nil {
		return err
	}
	err = models.RDB.Del(context.Background(), define.RedisMenuPrefix+roleIdentity).Err()
	if err != nil {
		return err
	}
	err = models.RDB.Del(context.Background(), define.RedisFuncPrefix+roleIdentity).Err()
	if err != nil {
		return err
	}
	return nil
}
