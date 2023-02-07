<template>
  <main class="main">
    <div class="box">
      <div class="banner">
        Up Admin 后台管理系统
      </div>
      <el-form
          ref="ruleFormRef"
          :model="userForm"
          class="form"
          size="large"
      >
        <el-form-item
            prop="username"
            :rules="userFormRules.username"
        >
          <el-input size="large" placeholder="请输入用户名" v-model="userForm.username" />
        </el-form-item>
        <el-form-item
            prop="password"
            :rules="userFormRules.password"
        >
          <el-input size="large" placeholder="请输入密码" type="password" v-model="userForm.password" />
        </el-form-item>
        <el-form-item>
          <el-button class="default-btn" type="primary" @click="login(ruleFormRef)" size="large" style="width: 100%; ">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </main>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'
import { loginPassword } from '@/api/user'
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";

const router = useRouter()
const ruleFormRef = ref<FormInstance>()
const userFormRules = reactive({
  username: [
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur',
    },
  ],
  password: [
    {
      required: true,
      message: '密码不能为空',
      trigger: 'blur',
    },
  ]
})
let userForm = ref({})

function judgeLoginState() {
  let token = localStorage.getItem("token")
  if (token != null) {
    router.push({ name: "home"})
  }
}
judgeLoginState()

const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      loginPassword(userForm.value).then((res: any) => {
        ElMessage({
          message: res.msg,
          type: 'success',
          duration: 2 * 1000
        })
        localStorage.setItem("token", res.data.token)
        localStorage.setItem("refresh_token", res.data.refresh_token)
        setTimeout(() => {
          router.push({ name: 'home' })
        }, 1500)
      })
    } else {
      console.log('error submit!')
      return false
    }
  })
}

</script>

<style scoped>
.banner {
  width: 100%;
  height: 60px;
  line-height: 60px;
  padding-left: 20px;
  font-size: 15px;
}
.main {
  width: 100vw;
  height: 100vh;
  background: url("@/assets/img/bg.png") no-repeat center;
  background-size: cover;
  overflow: hidden;
}
.box {
  margin: 200px auto;
  width: 380px;
  border: 1px solid #cdd0d6;
  border-radius: 5px;
  background-color: #fff;
}
.form {
  padding: 20px;
}
</style>