import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
import piniaPersist from 'pinia-plugin-persist'
import {createPinia} from 'pinia'

const pinia = createPinia()
pinia.use(piniaPersist)
const app = createApp(App)

app.use(pinia).use(router).mount('#app')
