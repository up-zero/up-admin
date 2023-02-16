<template>
  <el-menu
      default-active="2"
      @select="handleSelect"
      :collapse="collapseState"
      class="my-menu"
  >
    <template v-for="menu in menus">
      <!-- 无 子菜单 -->
      <el-menu-item v-if="menu.sub_menus == null" :index="menu.identity">
        <template v-if="menu.web_icon">
          <el-icon><component :is="menu.web_icon"/></el-icon>
        </template>
        <span>{{ menu.name }}</span>
      </el-menu-item>
      <!-- 有 子菜单 -->
      <el-sub-menu v-else :index="menu.identity">
        <template #title>
          <template v-if="menu.web_icon">
            <el-icon><component :is="menu.web_icon"/></el-icon>
          </template>
          <span>{{ menu.name }}</span>
        </template>
        <el-menu-item v-for="subMenu in menu.sub_menus" :index="subMenu.identity">
          <template v-if="subMenu.web_icon">
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
import {ref, watch} from "vue";

let menus = ref()
let props = defineProps(['collapseState'])
let emits = defineEmits(['emit-select-menu'])

const handleSelect = (_: string, indexPath: string) => {
  let menu = menus.value
  let index = 0
  let arr = [];
  for (let i = 0; i < menu.length; i++) {
    if (menu[i].identity == indexPath[index]) {
      arr.push(menu[i].name)
      index ++
      if (index == indexPath.length) {
        break
      }
      let subMenu = JSON.parse(JSON.stringify(menu[i].sub_menus))
      for (let j = 0; j < subMenu.length; j ++) {
        if (subMenu[j].identity == indexPath[index]) {
          arr.push(subMenu[j].name)
          break
        }
      }
      break
    }
  }
  console.log(arr)
  emits('emit-select-menu', arr)
}

// initMenu 初始化菜单
function initMenu() {
  let menu = localStorage.getItem("menu")
  if (menu == null) {
    getMenus().then((res: any) => {
      menus.value = res.data
      localStorage.setItem("menu", JSON.stringify(res.data))
      console.log(res.data)
    })
  } else {
    menus.value = JSON.parse(menu)
  }
}

initMenu()
</script>

<style scoped>
.my-menu {
  height: 100vh;
}
</style>