import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'

const routes = [
    {
        path: '/demo',
        name: 'demo',
        component: () => import('../views/demo.vue')
    },
    {
        path: '/',
        name: 'home',
        component: () => import('../views/home.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
