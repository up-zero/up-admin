import {defineStore} from 'pinia'
// @ts-ignore
import {getMenus} from '@/api/menu'

export const useStore = defineStore('counter', {
    state: () => {
        return {
            register: false,
            menus: [],
            userRouters: [],
            tags: []
        }
    },
    getters: {},
    actions: {
        getMenus() {
            return new Promise((resolve, reject) => {
                getMenus().then((res: any) => {
                    this.menus = res.data
                    resolve(res.data)
                })
            })
        }
    }
})
