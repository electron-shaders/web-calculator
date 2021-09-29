import { createApp } from 'vue'
import 'element-plus/theme-chalk/src/index.scss'
import 'element-plus/theme-chalk/display.css'
import store from './store/index'
import App from './App.vue'
import VueClipboard from 'vue-clipboard2'
import './assets/icons/iconfont.css'

const app = createApp(App)
app.use(store)
app.use(VueClipboard)
app.mount('#app')
