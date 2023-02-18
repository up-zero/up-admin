import {defineStore} from 'pinia'
import {getMenus} from '@/api/menu'

export const useCounterStore = defineStore('counter', {
    state: () => {
        return {
            register: false,
            menus: [],
            userRouters: []
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
