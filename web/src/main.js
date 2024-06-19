import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/style.css'
import App from '@/App.vue'

const app = createApp(App)

app.use(router).mount('#app')
