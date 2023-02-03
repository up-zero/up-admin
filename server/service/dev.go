package service

import (
	"gitee.com/up-zero/up-admin/helper"
	"gitee.com/up-zero/up-admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DevMenuAdd 新增菜单
func DevMenuAdd(c *gin.Context) {
	in := new(DevMenuAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	// 1. 获取父级ID
	var parentId uint
	if in.ParentIdentity != "" {
		err = models.DB.Model(new(models.MenuBasic)).Select("id").
			Where("identity = ?", in.ParentIdentity).Scan(&parentId).Error
		if err != nil {
			helper.Error("[DB ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据库异常",
			})
			return
		}
	}
	// 2. 保存数据
	err = models.DB.Create(&models.MenuBasic{
		Identity: helper.UUID(),
		ParentId: parentId,
		Name:     in.Name,
		Sort:     in.Sort,
	}).Error
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
		"msg":  "新增成功",
	})
}

// DevMenuUpdate 修改菜单
func DevMenuUpdate(c *gin.Context) {
	in := new(DevMenuUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.Identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 获取父级ID
	var parentId uint
	if in.ParentIdentity != "" {
		err = models.DB.Model(new(models.MenuBasic)).Select("id").
			Where("identity = ?", in.ParentIdentity).Scan(&parentId).Error
		if err != nil {
			helper.Error("[DB ERROR] : %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据库异常",
			})
			return
		}
	}
	// 2. 更新数据
	err = models.DB.Model(new(models.MenuBasic)).Where("identity = ?", in.Identity).Updates(map[string]interface{}{
		"parent_id": parentId,
		"name":      in.Name,
		"sort":      in.Sort,
	}).Error
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
		"msg":  "修改成功",
	})
}

// DevMenuDelete 删除菜单
func DevMenuDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	err := models.DB.Where("identity = ?", identity).Delete(new(models.MenuBasic)).Error
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

// DevFuncAdd 新增功能
func DevFuncAdd(c *gin.Context) {
	in := new(DevFuncAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.MenuIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 获取菜单ID
	var menuId uint
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity = ?", in.MenuIdentity).Scan(&menuId).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2. 保存数据
	err = models.DB.Create(&models.FunctionBasic{
		Identity: helper.UUID(),
		MenuId:   menuId,
		Name:     in.Name,
		Uri:      in.Uri,
		Sort:     in.Sort,
	}).Error
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
		"msg":  "新增成功",
	})
}

// DevFuncUpdate 修改功能
func DevFuncUpdate(c *gin.Context) {
	in := new(DevFuncUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Error("[BindJSON ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if in.Identity == "" || in.MenuIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 1. 获取菜单ID
	var menuId uint
	err = models.DB.Model(new(models.MenuBasic)).Select("id").
		Where("identity = ?", in.MenuIdentity).Scan(&menuId).Error
	if err != nil {
		helper.Error("[DB ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	// 2. 更新数据
	err = models.DB.Model(new(models.FunctionBasic)).Where("identity = ?", in.Identity).Updates(map[string]interface{}{
		"menu_id": menuId,
		"name":    in.Name,
		"uri":     in.Uri,
		"sort":    in.Sort,
	}).Error
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
		"msg":  "修改成功",
	})
}

// DevFuncDelete 删除功能
func DevFuncDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	err := models.DB.Where("identity = ?", identity).Delete(new(models.FunctionBasic)).Error
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
