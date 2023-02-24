<template>
<main>
  <el-table :data="roles" border style="width: 100%">
    <el-table-column prop="name" label="角色名称" />
    <el-table-column prop="sort" label="角色排序"/>
    <el-table-column label="是否是超管">
      <template #default="scope">
        <el-switch v-model="scope.row.is_admin" :active-value="1" :inactive-value="0" :before-change="handleChangeAdminState.bind(this, scope.row)"/>
      </template>
    </el-table-column>
    <el-table-column prop="created_at" label="创建时间"/>
    <el-table-column prop="updated_at" label="更新时间"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
        >Edit</el-button
        >
        <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
        >Delete</el-button
        >
      </template>
    </el-table-column>
  </el-table>
  <div style="line-height: 66px">
    <el-pagination
        background
        v-model:current-page="currentPage"
        layout="total, prev, pager, next"
        :total="roleCount"
        :default-page-size="size"
        @current-change="roleList"
    />
  </div>
</main>
</template>

<script lang="ts" setup>
import {getRoleList, setRoleUpdateAdmin} from '@/api/role'
import {ref} from "vue";
import {ElMessageBox} from "element-plus";

let roles = ref()
let roleCount = ref()
let currentPage = ref(1)
let size = ref(20)

const roleList = () => {
  getRoleList({size: size.value, page: currentPage.value}).then((res: any) => {
    roles.value = res.data.list
    roleCount.value = res.data.count
  })
}
roleList()

const handleChangeAdminState = (row: any) => {
  let admin_id = row.admin_id
  return new Promise((resolve, reject) => {
    ElMessageBox.confirm("确认修改当前角色身份么？", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
      autofocus: false
    }).then(() => {
      setRoleUpdateAdmin({identity: row.identity, is_admin: row.is_admin === 1 ? 0 : 1}).then((res: any) => {
        resolve(true)
      }).catch(() => {
        reject(false)
      })
    }).catch(() => {
      reject(false)
    })
  })
}
</script>

<style scoped>

</style>