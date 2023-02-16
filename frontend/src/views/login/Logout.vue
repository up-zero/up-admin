<template>
  <el-dropdown-item @click="logOut">退出登录</el-dropdown-item>
</template>

<script lang="ts" setup>
import {ElMessageBox} from "element-plus";

const logOut = () => {
  ElMessageBox.confirm("确认退出么？", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    autofocus: false
  }).then(() => {
    // 清除用户信息
    localStorage.removeItem("refresh_token")
    localStorage.removeItem("token")
    localStorage.removeItem("username")
    localStorage.removeItem("menu")
    // 跳转到登录页
    window.location.href = "/login"
  }).catch(() => {})
}
</script>

<style scoped>

</style>