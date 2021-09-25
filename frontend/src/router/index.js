import { createRouter, createWebHashHistory } from 'vue-router'
 
const routes = [
  {
    path: '/normal',
    name: 'normal',
    component: () => import('../components/normal.vue')
  },
  {
    path: '/keyboard',
    name: 'keyboard',
    component: () => import('../components/keyboard.vue')
  }
]
 
const router = createRouter({
  history: createWebHashHistory(),
  routes
})
 
export default router