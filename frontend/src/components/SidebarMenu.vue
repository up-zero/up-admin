<template>
  <el-menu
      :default-active="router.currentRoute.value.path"
      :collapse="collapseState"
      class="my-menu"
      :router="true"
  >
    <template v-for="menu in store.userRouters">
      <!-- 无 子菜单 -->
      <el-menu-item v-if="menu.children == null || menu.children.length === 0" :index="menu.path">
        <template v-if="menu.meta">
          <el-icon><component :is="menu.meta.icon"/></el-icon>
        </template>
        <span>{{ menu.name }}</span>
      </el-menu-item>
      <!-- 有 子菜单 -->
      <el-sub-menu v-else :index="menu.path">
        <template #title>
          <template v-if="menu.meta&&menu.meta.icon">
            <el-icon><component :is="menu.meta.icon"/></el-icon>
          </template>
          <span>{{ menu.name }}</span>
        </template>
        <el-menu-item v-for="subMenu in menu.children" :index="subMenu.path">
          <template v-if="subMenu.meta&&subMenu.meta.icon">
            <el-icon><component :is="subMenu.meta.icon"/></el-icon>
          </template>
          <span>{{ subMenu.name }}</span>
        </el-menu-item>
      </el-sub-menu>
    </template>
  </el-menu>
</template>

<script lang="ts" setup>
import { getMenus } from '@/api/menu'
import {ref, toRefs, watch} from "vue";

import { useStore } from '@/stores'
import router from "@/router";
const store = useStore()
let props = defineProps(['collapseState'])
let emits = defineEmits(['emit-select-menu'])

const handleSelect = (_: string, indexPath: string) => {
  let menu = store.menus
  let index = 0
  let arr = [];
  let toPath = "";
  for (let i = 0; i < menu.length; i++) {
    if (menu[i].identity == indexPath[index]) {
      arr.push(menu[i].name)
      index ++
      if (menu[i].path) {
        toPath = menu[i].path
      }
      if (index == indexPath.length) {
        break
      }
      let subMenu = JSON.parse(JSON.stringify(menu[i].sub_menus))
      for (let j = 0; j < subMenu.length; j ++) {
        if (subMenu[j].identity == indexPath[index]) {
          arr.push(subMenu[j].name)
          if (subMenu[j].path) {
            toPath = subMenu[j].path
          }
          break
        }
      }
      break
    }
  }
  emits('emit-select-menu', arr)
  router.push(toPath)
}
</script>

<style scoped>
.my-menu {
  height: 100vh;
}
</style>