import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/src/index.scss'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import store from './store/index'
import router from './router/index'
import App from './App.vue'
import './assets/icons/iconfont.css'

const app = createApp(App)
app.use(ElementPlus,{locale: zhCn})
app.use(router)
app.use(store)
app.mount('#app')
