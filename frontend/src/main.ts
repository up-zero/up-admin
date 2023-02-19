import {createApp} from 'vue'
import {createPinia} from 'pinia'
import {useStore} from '@/stores'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'

import './assets/main.css'
// import * as module from "module";

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(createPinia())
app.use(router)
app.use(ElementPlus)
let store = useStore()
app.mount('#app')

const modules = import.meta.glob('./views/**/*.vue');

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
        // @ts-ignore
        let routes=[];
        store.menus.forEach((item: any) => {
            let myRoute:any={};
            if (item.path) {
                myRoute={
                    path: item.path,
                    name: item.name,
                    meta:{
                        icon:item.web_icon
                    },
                    component:modules['./views' + item.path + '.vue']
                }
            }else{
                myRoute={
                    path:'/layout',
                    name:item.name,
                    meta:{
                        icon:item.web_icon
                    },
                    component: modules['./components/layout.vue']
                }
            }
            if (item.sub_menus) {
                myRoute.children=[]
                item.sub_menus.forEach((subItem: any) => {
                    if (subItem.path) {
                        myRoute.children.push({
                            path: subItem.path,
                            name: subItem.name,
                           meta:{
                               icon:subItem.web_icon
                           },
                            // component: ()=>import('./views' + subItem.path + '.vue')
                            component:modules['./views' + subItem.path + '.vue']
                        })
                    }
                });
            }
            routes.push(myRoute)
        })
        // @ts-ignore
        routes.forEach(item=>{
            router.addRoute('home',item)
        })
        // @ts-ignore
        store.userRouters = routes

        store.register = true
        next(to.fullPath)
    } else {
        next()
    }
})
