import request from "../utils/request";

export function loginPassword(data) {
    return request({
        url: "/login/password",
        method: "post",
        data,
    })
}