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

// SetUserList 管理员列表
func SetUserList(c *gin.Context) {
	in := &SetUserListRequest{NewQueryRequest()}
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
		list = make([]*SetUserListReply, 0)
	)
	err = models.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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

// SetUserAdd 管理员创建
func SetUserAdd(c *gin.Context) {
	in := new(SetUserAddRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否存在
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).Where("username = ?", in.Username).Count(&cnt).Error
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
			"msg":  "用户已存在",
		})
		return
	}

	// 2. 创建用户数据
	err = models.DB.Create(&models.UserBasic{
		Identity:     helper.UUID(),
		RoleIdentity: in.RoleIdentity,
		Username:     in.Username,
		Password:     helper.Md5(in.Password),
		Phone:        in.Phone,
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
		"msg":  "创建成功",
	})
}

// SetUserUpdate 管理员修改
func SetUserUpdate(c *gin.Context) {
	in := new(SetUserUpdateRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		helper.Info("[INPUT ERROR] : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否已存在
	var cnt int64
	err = models.DB.Model(new(models.UserBasic)).Where("identity != ? AND username = ?", in.Identity, in.Username).Count(&cnt).Error
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
			"msg":  "用户已存在",
		})
		return
	}
	// 2. 修改数据
	err = models.DB.Model(new(models.UserBasic)).Where("identity = ?", in.Identity).Updates(map[string]any{
		"role_identity": in.RoleIdentity,
		"username":      in.Username,
		"phone":         in.Phone,
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

// SetUserDelete 管理员删除
func SetUserDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}
	// 删除管理员
	err := models.DB.Where("identity = ?", identity).Delete(new(models.UserBasic)).Error
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
