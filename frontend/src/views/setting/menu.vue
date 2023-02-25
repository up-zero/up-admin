<template>
<main>
  <el-table :data="menus" border style="width: 100%">
    <el-table-column prop="name" label="菜单名称"/>
    <el-table-column prop="web_icon" label="WebIcon" />
    <el-table-column prop="sort" label="菜单排序" />
    <el-table-column prop="path" label="路径" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
        >Edit</el-button
        >
        <el-button
            size="small"
            type="danger"
            @click="handleMenuDelete(scope.row)"
        >Delete</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</main>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {getMenusList, devMenuDelete} from "@/api/menu"
import {ElMessage, ElMessageBox} from "element-plus";

let menus = ref()

const menuList = () => {
  getMenusList().then((res:any) => {
    menus.value = res.data
    console.log(res.data)
  })
}
menuList()

const handleMenuDelete = (row: any) => {
  ElMessageBox.confirm("确认删除么？", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    autofocus: false,
  }).then(() => {
    devMenuDelete({identity: row.identity}).then((res: any) => {
      ElMessage.success("删除成功")
      menuList()
    }).catch(() => {})
  }).catch(() => {})
}
</script>

<style scoped>

</style>