import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
import piniaPersist from 'pinia-plugin-persist'
import {createPinia} from 'pinia'
import i18n from '@/utils/i18n.js'

const pinia = createPinia()
pinia.use(piniaPersist)
const app = createApp(App)

app.use(pinia).use(router).use(i18n).mount('#app')
