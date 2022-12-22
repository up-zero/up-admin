package test

import (
	"encoding/json"
	"fmt"
	"gitee.com/up-zero/up-admin/helper"
	"github.com/gin-gonic/gin"
	"testing"
)

var baseURL = "http://localhost:9090"

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
