<template>
<main>
  <el-row style="margin-bottom: 15px">
    <el-col :span="2">
      <el-button size="large" type="primary" style="width: 90%" @click="showMenuDialog('create-top', null)">新增顶层菜单</el-button>
    </el-col>
  </el-row>
  <el-table :data="menus" row-key="identity" border :tree-props="{ children: 'sub_menus' }" style="width: 100%">
    <el-table-column prop="name" label="菜单名称"/>
    <el-table-column prop="web_icon" label="WebIcon" />
    <el-table-column prop="sort" label="菜单排序" />
    <el-table-column prop="level" label="菜单类型">
      <template #default="scope">
        <span v-if="scope.row.level === 0">目录</span>
        <span v-else-if="scope.row.level === 1">菜单</span>
        <span v-else>按钮</span>
      </template>
    </el-table-column>
    <el-table-column prop="path" label="路径" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button v-if="scope.row.level !== 2" type="primary" size="small" @click="showMenuDialog('create-sub', scope.row)"
        >新增子级</el-button
        >
        <el-button size="small" @click="showMenuDialog('edit', scope.row)"
        >编辑</el-button
        >
        <el-button
            size="small"
            type="danger"
            @click="handleMenuDelete(scope.row)"
        >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
  <el-dialog
      v-model="menuDialogVisible"
      :title="menuDialogType === 'edit' ? '编辑菜单': '新增菜单'"
      width="30%"
  >
    <el-form
        :model="menuManage"
        label-width="90px"
        ref="menuManageRules"
    >
      <el-form-item
          prop="name"
          label="菜单名称"
          :rules="menuManageRules.name"
      >
        <el-input placeholder="请输入菜单名称" v-model="menuManage.name"/>
      </el-form-item>
      <el-form-item
          prop="sort"
          label="菜单排序"
          :rules="menuManageRules.sort"
      >
        <el-input-number placeholder="请输入菜单排序" v-model="menuManage.sort"/>
      </el-form-item>
      <el-form-item
          prop="path"
          label="菜单路径"
      >
        <el-input placeholder="请输入菜单路径" v-model="menuManage.path"/>
      </el-form-item>
      <el-form-item
          prop="web_icon"
          label="菜单图标"
      >
        <el-input placeholder="请输入菜单图标" v-model="menuManage.web_icon"/>
      </el-form-item>
      <el-form-item
          prop="level"
          label="菜单等级"
      >
        <el-select v-model="menuManage.level" placeholder="请选择菜单等级" >
          <el-option label="目录" :value="0" />
          <el-option label="菜单" :value="1" />
          <el-option label="按钮" :value="2" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span>
        <el-button @click="menuDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmSaveMenu(menuManageRules)">确认</el-button>
      </span>
    </template>
  </el-dialog>
</main>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {getMenusList, devMenuDelete, devMenuAdd, devMenuUpdate} from "@/api/menu"
import {ElMessage, ElMessageBox} from "element-plus";
import type {FormInstance} from "element-plus";

let menus = ref()
let menuDialogVisible = ref(false)
let menuDialogType = ref('')
let menuManage = ref({
  identity: '',
  parent_identity: '',
  name: '',
  sort: 0,
  path: '',
  web_icon: '',
  level: 0,
})
const menuManageRules = reactive({
  name: [
    {
      required: true,
      message: '菜单名称不能为空',
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

const menuList = () => {
  getMenusList().then((res:any) => {
    menus.value = res.data
    console.log(res.data)
  })
}
menuList()

const showMenuDialog = (dialogType: string, row: any) => {
  menuDialogType.value = dialogType
  if (dialogType === 'create-top') {
    menuManage.value.parent_identity = ""
  } else if (dialogType === 'create-sub') {
    menuManage.value.parent_identity = row.identity
  } else {
    menuManage.value.identity = row.identity
    menuManage.value.parent_identity = row.parent_identity
    menuManage.value.name = row.name
    menuManage.value.sort = row.sort
    menuManage.value.path = row.path
    menuManage.value.web_icon = row.web_icon
    menuManage.value.level = row.level
  }
  menuDialogVisible.value = true
}

const confirmSaveMenu = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      if (menuDialogType.value === 'edit') {
        devMenuUpdate(menuManage.value).then((res: any) => {
          ElMessage.success("编辑成功")
          menuList()
          menuDialogVisible.value = false
        })
      } else {
        devMenuAdd(menuManage.value).then((res: any) => {
          ElMessage.success("新增成功")
          menuList()
          menuDialogVisible.value = false
        })
      }
    }
  })
}

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