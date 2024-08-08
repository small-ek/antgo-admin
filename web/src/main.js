import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
const app = createApp(App)

app.use(router).mount('#app')
