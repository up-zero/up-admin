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
