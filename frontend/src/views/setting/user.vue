<template>
  <main>
    <el-row style="margin-bottom: 15px">
      <el-col :span="2">
        <el-button size="large" type="primary" style="width: 90%" @click="showRoleDialog('create')">创建管理员</el-button>
      </el-col>
      <!-- 角色列表搜索 -->
      <el-col :span="6">
        <el-input size="large" v-model="keyword" placeholder="请输入管理员名称" />
      </el-col>
      <el-col :span="2">
        <el-button size="large" type="primary" :icon="Search" @click="userList">搜索</el-button>
      </el-col>
    </el-row>
    <el-table :data="users" border style="width: 100%">
      <el-table-column prop="username" label="管理员名称" />
      <el-table-column prop="role_name" label="角色名称" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="created_at" label="创建时间" />
      <el-table-column prop="updated_at" label="更新时间" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" @click="showRoleDialog('edit', scope.row)"
          >编辑</el-button
          >
          <el-button
              size="small"
              type="danger"
              @click="handleRoleDelete(scope.row)"
          >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <div style="line-height: 66px">
      <el-pagination
          background
          v-model:current-page="currentPage"
          layout="total, prev, pager, next"
          :total="userCount"
          :default-page-size="size"
          @current-change="userList"
      />
    </div>
  </main>
</template>

<script lang="ts" setup>
import {setUserList} from '@/api/user'
import {ref} from "vue";

let users = ref()
let userCount = ref()
let keyword = ref()
let currentPage = ref(1)
let size = ref(20)

const userList = () => {
  setUserList({size: size.value, page: currentPage.value, keyword: keyword.value}).then((res: any) => {
    users.value = res.data.list
    userCount.value = res.data.count
  })
}
userList()
</script>

<style scoped>

</style>