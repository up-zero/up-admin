package service

import "gitee.com/up-zero/up-admin/define"

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		KeyWord: "",
	}
}
