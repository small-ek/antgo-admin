import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
import {createPinia} from 'pinia'

const pinia = createPinia()
const app = createApp(App)

app.use(pinia).use(router).mount('#app')
