import {createApp} from 'vue'
import piniaPersist from 'pinia-plugin-persist'
import {createPinia} from 'pinia'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
import i18n from '@/utils/i18n.js'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import {initRouter} from "@/utils/permission.js";

/* add icons to the library */
library.add(fas)

const pinia = createPinia()
pinia.use(piniaPersist)
const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)

const boot = async () => {
    app.use(pinia)
    await initRouter()
    app.use(i18n).use(router).mount('#app')
}

boot()

