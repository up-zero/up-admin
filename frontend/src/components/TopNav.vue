<template>
  <el-menu
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
      :ref="myMenu"
  >
    <el-icon class="my-state" @click="changeCollapseState">
      <Fold v-if="!collapse" />
      <Expand v-else/>
    </el-icon>
    <span v-for="menuName in selectMenuArr" class="my-menu-name">/&nbsp;{{menuName}}</span>

    <div class="flex-grow" />
    <el-dropdown :hide-on-click="false">
      <span class="my-dropdown-link">
        {{ username }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <Logout />
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </el-menu>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import Logout from "@/views/login/Logout.vue";

let myMenu = ref();
let collapse = ref(false) // 是否折叠
let props = defineProps(['selectMenuArr'])
let emits = defineEmits(['emit-collapse-state'])
let username = localStorage.getItem("username")

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  console.log(myMenu)
}

const changeCollapseState = () => {
  collapse.value = !collapse.value
  emits('emit-collapse-state', collapse.value)
}
</script>

<style scoped>
.flex-grow {
  flex-grow: 1;
}
.my-state {
  padding: 20px;
}
.my-state:hover {
  cursor: pointer;
}
.my-dropdown-link {
  line-height: 57.6px;
  margin-right: 20px;
}
.my-dropdown-link:hover {
  cursor: pointer;
}
.my-menu-name {
  color: #97a8be;
  font-size: 14px;
  padding-right: 10px;
  line-height: 55px;
}
</style>