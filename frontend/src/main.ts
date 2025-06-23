import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import './assets/main.css'
import { createPinia } from 'pinia'
import { router } from './router'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faChartLine, faBuilding, faTag, faBullseye, faStarHalfStroke, faCalendar, faCircleCheck , faCircleDot, faChartColumn, faHeart, faUpLong, faDownLong, faPause, faArrowLeft, faArrowRight, faBackward, faForward} from '@fortawesome/free-solid-svg-icons'
library.add(faChartLine, faBuilding, faTag, faBullseye, faStarHalfStroke, faCalendar, faCircleCheck, faCircleDot, faChartColumn, faHeart, faUpLong, faDownLong, faPause, faArrowLeft, faArrowRight, faBackward, faForward)


const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(createPinia())
app.use(router)
app.mount('#app')
