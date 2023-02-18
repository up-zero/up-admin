import {createApp} from 'vue'
import {createPinia} from 'pinia'
import {useCounterStore} from '@/stores'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'

import './assets/main.css'

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(createPinia())
app.use(router)
app.use(ElementPlus)
let store = useCounterStore()
app.mount('#app')

// const modules = import.meta.glob('./views/**/*.vue');

// 注册路由
router.beforeEach(async (to, form, next) => {
    let menus = localStorage.getItem("menu")
    if (!store.register && to.path != '/login') {
        let res;
        if (!menus) {
            res = await store.getMenus()
            localStorage.setItem('menu', JSON.stringify(res))
        } else {
            store.menus = JSON.parse(menus)
        }
        // 处理 menu
        store.menus.forEach((item: any) => {
            if (item.path) {
                router.addRoute('home', {
                    path: item.path,
                    name: item.name,
                    component: ()=>import('./views/' + item.path + '.vue')
                })
            }
            if (item.sub_menus) {
                item.sub_menus.forEach((subItem: any) => {
                    if (subItem.path) {
                        router.addRoute('home', {
                            path: subItem.path,
                            name: subItem.name,
                            component: ()=>import('./views/' + subItem.path + '.vue')
                            // component: modules['./views/' + subItem.path + '.vue']
                        })
                    }
                });
            }
        })
        store.userRouters = router.getRoutes()
console.log(store.userRouters)
        store.register = true
        next(to.fullPath)
    } else {
        next()
    }
})
