<template>
  <el-menu
      default-active="2"
      @select="handleSelect"
  >
    <template v-for="menu in menus">
      <!-- 无 子菜单 -->
      <el-menu-item v-if="menu.sub_menus == null" :index="menu.identity">
        <template v-if="menu.web_icon != null">
          <el-icon><component :is="menu.web_icon"/></el-icon>
        </template>
        <span>{{ menu.name }}</span>
      </el-menu-item>
      <!-- 有 子菜单 -->
      <el-sub-menu v-else :index="menu.identity">
        <template #title>
          <template v-if="menu.web_icon != null">
            <el-icon><component :is="menu.web_icon"/></el-icon>
          </template>
          <span>{{ menu.name }}</span>
        </template>
        <el-menu-item v-for="subMenu in menu.sub_menus" :index="subMenu.identity">
          <template v-if="subMenu.web_icon != null">
            <el-icon><component :is="subMenu.web_icon"/></el-icon>
          </template>
          <span>{{ subMenu.name }}</span>
        </el-menu-item>
      </el-sub-menu>
    </template>
  </el-menu>
</template>

<script lang="ts" setup>
import { getMenus } from '@/api/menu'
import {ref} from "vue";

const handleSelect = (key: string) => {
  console.log(key)
}

let menus = ref({})

// initMenu 初始化菜单
function initMenu() {
  let menu = localStorage.getItem("menu")
  if (menu == null) {
    getMenus().then((res: any) => {
      menus.value = res.data
      console.log(res.data)
    })
  }
}

initMenu()
</script>

<style scoped>
</style>