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
    <div class="flex-grow" />
    <el-dropdown :hide-on-click="false">
      <span class="my-dropdown-link">
        姓名<el-icon class="el-icon--right"><arrow-down /></el-icon>
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
let emits = defineEmits(['emit-collapse-state'])

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
</style>