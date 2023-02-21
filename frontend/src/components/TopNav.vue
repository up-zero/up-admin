<template>
  <main>
    <el-menu
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="false"
        :ref="myMenu"
    >
      <el-icon class="my-state" @click="changeCollapseState">
        <Fold v-if="!collapse"/>
        <Expand v-else/>
      </el-icon>
      <span v-for="(menuName,i) in route.matched" class="my-menu-name"
            v-show="i!==0">/&nbsp;{{ menuName.name === 'Index' ? '首页' : menuName.name }}</span>

      <div class="flex-grow"/>
      <el-dropdown :hide-on-click="false">
      <span class="my-dropdown-link">
        {{ username }}<el-icon class="el-icon--right"><arrow-down/></el-icon>
      </span>
        <template #dropdown>
          <el-dropdown-menu>
            <!--            <change-password />-->
            <el-dropdown-item @click="changePasswordDialogVisible = true">修改密码</el-dropdown-item>
            <Logout/>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-menu>
    <div class="my-line-box" v-if="store.tags.length > 0">
      <el-tag
          v-for="tag in store.tags"
          closable
          class="my-tag"
          :key="tag.path"
          @close="tagCloseHandle(tag)"
          @click="toPathHandle(tag)"
          :effect="tag.path === route.path ? 'dark':'plain'"
      >
        {{ tag.name }}
      </el-tag>
    </div>
    <el-dialog
        v-model="changePasswordDialogVisible"
        title="修改密码"
        width="30%"
    >
      <el-form
          ref="ruleFormRef"
          :model="userForm"
          label-width="90px"
      >
        <el-form-item
            prop="old_password"
            label="旧密码"
            :rules="userFormRules.old_password"
        >
          <el-input placeholder="请输入旧密码" v-model="userForm.old_password" type="password"/>
        </el-form-item>
        <el-form-item
            prop="new_password"
            label="新密码"
            :rules="userFormRules.new_password"
        >
          <el-input placeholder="请输入新密码" v-model="userForm.new_password" type="password"/>
        </el-form-item>
        <el-form-item
            prop="confirm_password"
            label="确认密码"
            :rules="userFormRules.confirm_password"
        >
          <el-input placeholder="请确认密码" v-model="userForm.confirm_password" type="password"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span>
        <el-button @click="changePasswordDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="changePasswordConfirm(ruleFormRef)">确认</el-button>
      </span>
      </template>
    </el-dialog>
  </main>
</template>

<script lang="ts" setup>
import {reactive, ref, watch} from "vue";
import Logout from "@/views/login/Logout.vue";
import {useRoute} from "vue-router";
import {useStore} from "@/stores";
import router from "@/router";
import type {FormInstance} from "element-plus";
import {ElMessage} from "element-plus";
import { userPasswordChange } from '@/api/user'

const route = useRoute()
const store = useStore()
let myMenu = ref();
let collapse = ref(false) // 是否折叠
let props = defineProps(['selectMenuArr'])
let emits = defineEmits(['emit-collapse-state'])
let username = localStorage.getItem("username")

const changeCollapseState = () => {
  collapse.value = !collapse.value
  emits('emit-collapse-state', collapse.value)
}

watch(() => route.path, (path) => {
  let index = store.tags.findIndex(item => item.path == path)
  if (index == -1) {
    let name = route.matched[route.matched.length - 1].name
    if (name == 'Index') {
      name = '首页'
    }
    store.tags.push({
      path: path,
      name: name
    })
  }
})

const tagCloseHandle = (tag: any) => {
  let index = store.tags.findIndex(item => item.path == tag.path)
  store.tags.splice(index, 1)
}

const toPathHandle = (tag: any) => {
  router.push(tag.path)
}

const changePasswordDialogVisible = ref(false)
const ruleFormRef = ref<FormInstance>()
const userFormRules = reactive({
  old_password: [
    {
      required: true,
      message: '旧密码不能为空',
      trigger: 'blur',
    },
  ],
  new_password: [
    {
      required: true,
      message: '新密码不能为空',
      trigger: 'blur',
    },
  ],
  confirm_password: [
    {
      required: true,
      message: '确认密码不能为空',
      trigger: 'blur',
    },
  ]
})
let userForm = ref({
  old_password: "",
  new_password: "",
  confirm_password: ""
})

const changePasswordConfirm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      if (userForm.value.new_password != userForm.value.confirm_password) {
        ElMessage({
          message: "新旧密码不一致",
          type: 'error',
        })
        return
      }
      userPasswordChange(userForm.value).then((res: any) => {
        ElMessage({
          message: res.msg,
          type: 'success',
          duration: 2 * 1000
        })
      })
      changePasswordDialogVisible.value = false
    } else {
      console.log('error submit!')
      return false
    }
  })
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

.my-line-box {
  padding: 10px;
  border-bottom: solid 1px var(--el-menu-border-color);
}

.my-tag {
  margin-right: 12px;
  cursor: pointer;
}
</style>