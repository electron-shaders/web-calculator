import { createApp } from 'vue'
import { components, plugins } from './plugins/element';
import 'element-plus/theme-chalk/src/index.scss'
import 'element-plus/theme-chalk/display.css'
import store from './store/index'
import router from './router/index'
import App from './App.vue'
import VueClipboard from 'vue-clipboard2'
import './assets/icons/iconfont.css'

const app = createApp(App)
components.forEach(component => {
    app.component(component.name, component)
})
plugins.forEach(plugin => {
    app.use(plugin)
})
app.use(router)
app.use(store)
app.use(VueClipboard)
app.mount('#app')
