<template>
<main>
  <el-row style="margin-bottom: 15px">
    <el-col :span="2">
      <el-button size="large" type="primary" style="width: 90%" @click="showRoleDialog('create')">创建角色</el-button>
    </el-col>
    <!-- 角色列表搜索 -->
    <el-col :span="6">
      <el-input size="large" v-model="keyword" placeholder="请输入角色名称" />
    </el-col>
    <el-col :span="2">
      <el-button size="large" type="primary" :icon="Search" @click="roleList">搜索</el-button>
    </el-col>
  </el-row>
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
        :total="roleCount"
        :default-page-size="size"
        @current-change="roleList"
    />
  </div>
  <el-dialog
      v-model="roleDialogVisible"
      :title="roleDialogType === 'create' ? '创建角色': '编辑角色'"
      width="30%"
  >
    <el-form
        ref="roleManageRules"
        :model="roleManage"
        label-width="90px"
    >
      <el-form-item
          prop="name"
          label="角色名称"
          :rules="roleManageRules.name"
      >
        <el-input placeholder="请输入角色名称" v-model="roleManage.name"/>
      </el-form-item>
      <el-form-item
          prop="sort"
          label="角色排序"
          :rules="roleManageRules.sort"
      >
        <el-input placeholder="请输入角色排序" v-model="roleManage.sort"/>
      </el-form-item>
      <el-form-item
          prop="is_admin"
          label="超级管理员"
      >
        <el-switch v-model="roleManage.is_admin" :active-value="1" :inactive-value="0"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span>
        <el-button @click="roleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmSaveRole(roleManageRules)">确认</el-button>
      </span>
    </template>
  </el-dialog>
</main>
</template>

<script lang="ts" setup>
import {getRoleList, setRoleUpdateAdmin, setRoleDelete} from '@/api/role'
import {reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import type { FormInstance} from "element-plus";
import { Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'

let roles = ref()
let roleCount = ref()
let keyword = ref()
let currentPage = ref(1)
let size = ref(20)

let roleDialogVisible = ref(false)
let roleDialogType = ref("")
let roleManage = ref({
  name: "",
  sort: 0,
  is_admin: 0,
  menu_identities: [],
  func_identities: []
})
const roleManageRules = reactive({
  name: [
    {
      required: true,
      message: '角色名称不能为空',
      trigger: 'blur',
    },
  ],
  sort: [
    {
      required: true,
      message: '排序不能为空',
      trigger: 'blur',
    }
  ]
})

const roleList = () => {
  getRoleList({size: size.value, page: currentPage.value, keyword: keyword.value}).then((res: any) => {
    roles.value = res.data.list
    roleCount.value = res.data.count
  })
}
roleList()

// dialogType {create: 创建, update: 编辑}
const showRoleDialog = (dialogType: string) => {
  roleDialogVisible.value = true
  roleDialogType.value = dialogType
}

// confirmSaveRole 角色提交
const confirmSaveRole = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log("RUN ..")
      console.log(roleManage.value)
    }
  })
}

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

const handleRoleDelete = (row: any) => {
  ElMessageBox.confirm("确认删除当前角色么？", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    autofocus: false,
  }).then(() => {
    setRoleDelete({identity: row.identity}).then((res: any) => {
      ElMessage({
        message: "删除成功",
        type: 'success',
      })
      roleList()
    }).catch(() => {})
  }).catch(() => {})
}
</script>

<style scoped>

</style>