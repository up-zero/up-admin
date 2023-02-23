// 获取菜单列表
import request from "../utils/request";

export function getMenus() {
    return request({
        url: "/menus",
        method: "get",
    })
}

export function getMenusList() {
    return request({
        url: "/set/menu/list",
        method: "get",
    })
}