import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/login/Index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/Index.vue'),
    },
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/Index.vue'),
      children: [
        {
          path: '/home',
          name: 'new-home',
          component: () => import('@/views/home/Index.vue')
        }
      ]
    },
  ]
})

export default router
