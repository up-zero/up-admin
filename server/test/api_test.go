package test

import (
	"encoding/json"
	"fmt"
	"gitee.com/up-zero/up-admin/helper"
	"github.com/gin-gonic/gin"
	"testing"
)

var (
	baseURL = "http://localhost:9090"
	token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiSWRlbnRpdHkiOiIiLCJOYW1lIjoiZ2V0IiwiUm9sZUlkZW50aXR5IjoiMSIsImV4cCI6MTY3MzQyOTU4OH0.LoH4-MpSJV2xMAqF2cNEvjIqnNRDnnPESqhuuaodiHw"
	header  []byte
)

func init() {
	header, _ = json.Marshal(gin.H{
		"AccessToken": token,
	})
}

// 登录
func TestLoginPassword(t *testing.T) {
	m := gin.H{
		"username": "get",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	resp, err := helper.HttpPost(baseURL+"/login/password", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 获取菜单列表
func TestMenus(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/menus", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}
