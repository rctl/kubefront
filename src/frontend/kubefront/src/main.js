import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import 'materialize-css/sass/materialize.scss'
import './assets/sass/_material-icons.scss'
import M from "materialize-css"

Vue.config.productionTip = false

window.axios = axios.create({
  baseURL: "http://localhost:8081/"
})

Vue.prototype.$http = window.axios

new Vue({
  router,
  render: h => h(App)
}).$mount('main')
