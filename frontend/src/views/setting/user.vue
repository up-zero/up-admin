<template>
  <main>
    <el-row style="margin-bottom: 15px">
      <el-col :span="2">
        <el-button size="large" type="primary" style="width: 90%" @click="showUserDialog('create')">创建管理员</el-button>
      </el-col>
      <!-- 角色列表搜索 -->
      <el-col :span="6">
        <el-input size="large" v-model="keyword" placeholder="请输入管理员名称"/>
      </el-col>
      <el-col :span="2">
        <el-button size="large" type="primary" :icon="Search" @click="userList">搜索</el-button>
      </el-col>
    </el-row>
    <el-table :data="users" border style="width: 100%">
      <el-table-column prop="username" label="管理员名称"/>
      <el-table-column prop="role_name" label="角色名称"/>
      <el-table-column prop="phone" label="手机号"/>
      <el-table-column prop="created_at" label="创建时间"/>
      <el-table-column prop="updated_at" label="更新时间"/>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" @click="showUserDialog('edit', scope.row)"
          >编辑
          </el-button
          >
          <el-button
              size="small"
              type="danger"
              @click="handleUserDelete(scope.row)"
          >删除
          </el-button
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
    <el-dialog
        v-model="userDialogVisible"
        :title="userDialogType === 'create' ? '创建角色': '编辑角色'"
        width="40%"
    >
      <el-form
          ref="userManageRules"
          :model="userManage"
          label-width="100px"
      >
        <el-form-item
            prop="username"
            label="管理员名称"
            :rules="userManageRules.username"
        >
          <el-input placeholder="请输入管理员名称" v-model="userManage.username"/>
        </el-form-item>
        <el-form-item
            prop="password"
            label="管理员密码"
        >
          <el-input placeholder="请输入管理员密码" v-model="userManage.password" type="password"/>
        </el-form-item>
        <el-form-item
            prop="phone"
            label="管理员手机号"
        >
          <el-input placeholder="请输入管理员手机号" v-model="userManage.phone"/>
        </el-form-item>

        <el-form-item
            prop="role_identity"
            label="管理员角色"
            :rules="userManageRules.role_identity"
        >
          <el-select v-if="preRoles.length > 0"
                     v-model="userManage.role_identity"
                     filterable
                     remote
                     reserve-keyword
                     placeholder="请输入角色名称"
                     :remote-method="roleList"
          >
            <el-option v-for="item in preRoles" :key="item.identity" :label="item.name"
                       :value="item.identity"/>
          </el-select>
          <el-select v-else
                     v-model="userManage.role_identity"
                     filterable
                     remote
                     reserve-keyword
                     placeholder="请输入角色名称"
                     :remote-method="roleList"
          >
            <el-option
                v-for="item in roles"
                :key="item.identity"
                :label="item.name"
                :value="item.identity"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
      <span>
        <el-button @click="userDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmSaveUser(userManageRules)">确认</el-button>
      </span>
      </template>
    </el-dialog>
  </main>
</template>

<script lang="ts" setup>
import {setUserAdd, setUserDelete, setUserList, setUserUpdate} from '@/api/user'
import {getRoleList} from '@/api/role'
import {reactive, ref} from "vue";
import type {FormInstance} from "element-plus";
import {ElMessage, ElMessageBox} from "element-plus";

let roles = ref()
let users = ref()
let userCount = ref()
let keyword = ref()
let currentPage = ref(1)
let size = ref(20)

let userDialogVisible = ref(false)
let userDialogType = ref("")
let userManage = ref({
  identity: "",
  username: "",
  password: "",
  phone: "",
  role_identity: ""
})
let preRoles = ref([])

const userManageRules = reactive({
  username: [
    {
      required: true,
      message: '管理员名称不能为空',
      trigger: 'blur',
    },
  ],
  role_identity: [
    {
      required: true,
      message: '角色不能为空',
      trigger: 'blur',
    }
  ]
})

const userList = () => {
  setUserList({size: size.value, page: currentPage.value, keyword: keyword.value}).then((res: any) => {
    users.value = res.data.list
    userCount.value = res.data.count
  })
}
userList()

const roleList = (roleKeyword: string) => {
  preRoles.value = []
  getRoleList({size: 100, keyword: roleKeyword}).then((res: any) => {
    roles.value = res.data.list
  })
}

// dialogType {create: 创建, update: 编辑}
const showUserDialog = (dialogType: string, row: any) => {
  userDialogType.value = dialogType
  if (dialogType === 'create') {
    userManage.value.identity = ""
  } else {
    userManage.value.identity = row.identity
    userManage.value.username = row.username
    userManage.value.phone = row.phone
    userManage.value.role_identity = row.role_identity
    preRoles.value = [{
      identity: row.role_identity,
      name: row.role_name,
    }]
    console.log(preRoles.value)
  }
  userDialogVisible.value = true
}

const confirmSaveUser = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      if (!userManage.value.identity) { // 创建
        if (!userManage.value.password) {
          ElMessage.error("密码不能为空")
          return
        }
        setUserAdd(userManage.value).then((res: any) => {
          ElMessage.success("创建成功")
          userDialogVisible.value = false
          userList()
        })
      } else { // 修改
        setUserUpdate(userManage.value).then((res: any) => {
          ElMessage.success("修改成功")
          userDialogVisible.value = false
          userList()
        })
      }
    }
  })
}

const handleUserDelete = (row: any) => {
  ElMessageBox.confirm("确认删除当前用户么？", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    autofocus: false,
  }).then(() => {
    setUserDelete({identity: row.identity}).then((res: any) => {
      ElMessage({
        message: "删除成功",
        type: 'success',
      })
      userList()
    }).catch(() => {
    })
  }).catch(() => {
  })
}
</script>

<style scoped>

</style>