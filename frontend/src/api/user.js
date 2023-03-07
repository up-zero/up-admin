import request from "../utils/request";

// 登录
export function loginPassword(data) {
    return request({
        url: "/login/password",
        method: "post",
        data,
    })
}

// 修改密码
export function userPasswordChange(data) {
    return request({
        url: "/user/password/change",
        method: "put",
        data,
    })
}

// 获取管理员列表
export function setUserList(params) {
    return request({
        url: "/set/user/list",
        method: "get",
        params: params
    })
}

// 创建管理员
export function setUserAdd(data) {
    return request({
        url: "/set/user/add",
        method: "post",
        data
    })
}

// 修改管理员
export function setUserUpdate(data) {
    return request({
        url: "/set/user/update",
        method: "put",
        data
    })
}

// 删除管理员
export function setUserDelete(params) {
    return request({
        url: "/set/user/delete",
        method: "delete",
        params: params
    })
}
