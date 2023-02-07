import request from "../utils/request";

// 登录
export function loginPassword(data) {
    return request({
        url: "/login/password",
        method: "post",
        data,
    })
}