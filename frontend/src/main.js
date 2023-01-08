import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import Axios from 'axios'
import VueAxios from 'vue-axios'

import '../node_modules/element-plus/dist/index.css'

const app = createApp(App)
app.use(ElementPlus)
app.use(VueAxios, Axios)
app.mount('#app')
