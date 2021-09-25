import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/src/index.scss'
import '@element-plus/icons'
import router from './router/index'
import App from './App.vue'
import './assets/icons/iconfont.css'

createApp(App).use(ElementPlus).use(router).mount('#app')
