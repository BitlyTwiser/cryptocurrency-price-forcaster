import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import PrimeVue from 'primevue/config'
import "primevue/resources/themes/saga-blue/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"
import Dialog from 'primevue/dialog'
import ToastService from 'primevue/toastservice'
import VueApexCharts from "vue3-apexcharts";
const app = createApp(App)

app.use(router)
app.use(VueApexCharts)
app.use(VueAxios, axios)
app.use(PrimeVue, {})
app.use(ToastService)
// eslint-disable-next-line
app.component('Dialog', Dialog)
app.mount('#app')