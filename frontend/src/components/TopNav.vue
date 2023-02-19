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
  </main>
</template>

<script lang="ts" setup>
import {ref, watch} from "vue";
import Logout from "@/views/login/Logout.vue";
import {useRoute} from "vue-router";
import {useStore} from "@/stores";
import router from "@/router";

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