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
	token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiSWRlbnRpdHkiOiIxIiwiTmFtZSI6ImdldCIsIlJvbGVJZGVudGl0eSI6IjEiLCJJc0FkbWluIjpmYWxzZSwiZXhwIjoxNjc4MTA3NTYzfQ.665FEri_rqN_k3cZFnB5MhbKzktfFRF5GpajWDOJA1Y"
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

// 获取用户信息
func TestUserInfo(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/user/info", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 修改密码
func TestUserPasswordChange(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"old_password": "654321",
		"new_password": "123456",
	})
	resp, err := helper.HttpPut(baseURL+"/user/password/change", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 角色列表
func TestSetRoleList(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/set/role/list?page=1", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 角色详情
func TestSetRoleDetail(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/set/role/detail?identity=1", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 修改角色的管理员身份
func TestSetRoleUpdateAdmin(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"identity": "1",
		"is_admin": 1,
	})
	resp, err := helper.HttpPut(baseURL+"/set/role/update/admin", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 新增角色
func TestSetRoleCreate(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"name":            "新增角色测试",
		"sort":            0,
		"is_admin":        0,
		"menu_identities": []string{"1", "2"},
		"func_identities": []string{"1"},
	})
	resp, err := helper.HttpPost(baseURL+"/set/role/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 删除角色
func TestSetRoleDelete(t *testing.T) {
	resp, err := helper.HttpDelete(baseURL+"/set/role/delete?identity=1", []byte{}, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 修改角色
func TestSetRoleUpdate(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"identity":        "d1d56591-55db-484e-96a6-d94a5a833cd9",
		"name":            "修改角色测试",
		"sort":            0,
		"is_admin":        0,
		"menu_identities": []string{"1", "2"},
		"func_identities": []string{"1"},
	})
	resp, err := helper.HttpPut(baseURL+"/set/role/update", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 设置-获取菜单列表
// go test -v -run TestSetMenuList
func TestSetMenuList(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/set/menu/list", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 功能列表
func TestSetFuncList(t *testing.T) {
	resp, err := helper.HttpGet(baseURL+"/set/func/list", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 新增菜单
func TestDevMenuAdd(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"parent_identity": "",
		"name":            "新增顶层菜单测试",
		"sort":            0,
		"web_icon":        "Search",
		"path":            "/path",
	})
	resp, err := helper.HttpPost(baseURL+"/dev/menu/add", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 修改菜单
func TestDevMenuUpdate(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"identity":        "11435198-e662-40e7-987e-ab350dd4966e",
		"parent_identity": "1",
		"name":            "修改菜单测试",
		"sort":            0,
	})
	resp, err := helper.HttpPut(baseURL+"/dev/menu/update", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 删除菜单
func TestDevMenuDelete(t *testing.T) {
	resp, err := helper.HttpDelete(baseURL+"/dev/menu/delete?identity=11435198-e662-40e7-987e-ab350dd4966e", []byte{}, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 新增功能
func TestDevFuncAdd(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"menu_identity": "2",
		"name":          "新增功能测试",
		"uri":           "/test",
		"sort":          0,
	})
	resp, err := helper.HttpPost(baseURL+"/dev/func/add", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 修改功能
func TestDevFuncUpdate(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"identity":      "1cde2a40-0b82-42e7-bc30-acf469f47c86",
		"menu_identity": "1",
		"name":          "修改功能测试",
		"uri":           "/test2",
		"sort":          10,
	})
	resp, err := helper.HttpPut(baseURL+"/dev/func/update", data, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}

// 删除功能
func TestDevFuncDelete(t *testing.T) {
	resp, err := helper.HttpDelete(baseURL+"/dev/func/delete?identity=1cde2a40-0b82-42e7-bc30-acf469f47c86", []byte{}, header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
}
