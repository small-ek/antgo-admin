import {createApp} from 'vue'
import router from '@/routers/index.js'
import '@/styles/index.less'
import App from '@/App.vue'
import piniaPersist from 'pinia-plugin-persist'
import {createPinia} from 'pinia'
import i18n from '@/utils/i18n.js'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(fas)

const pinia = createPinia()
pinia.use(piniaPersist)
const app = createApp(App)

app.component('font-awesome-icon', FontAwesomeIcon)
app.use(pinia).use(router).use(i18n).mount('#app')
