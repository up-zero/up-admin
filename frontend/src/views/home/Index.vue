<template>
  <main>
    <el-container>
      <el-aside class="my-aside">
        <SidebarMenu :collapseState="collapseState"/>
      </el-aside>
      <el-container>
        <el-header class="my-header">
          <TopNav @emit-collapse-state="emitCollapseState"/>
        </el-header>
        <el-main>main</el-main>
      </el-container>
    </el-container>
  </main>
</template>

<script setup lang="ts">
import SidebarMenu from "@/components/SidebarMenu.vue";
import TopNav from "@/components/TopNav.vue";
import {ref} from "vue";

let collapseState = ref(false)

function emitCollapseState(collapse:boolean) {
  collapseState.value = collapse
}

</script>

<style scoped>
.my-header {
  padding: 0;
}
.my-aside {
  width: auto;
}
</style>