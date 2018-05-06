//Global imports required to run the app
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import 'materialize-css/sass/materialize.scss'
import './assets/sass/_material-icons.scss'
import M from "materialize-css"
import bus from "./bus"
import api from "./api"
import auth from "./services/auth"
import nodes from "./services/nodes"
import pods from "./services/pods"
import upstream from "./services/upstream"
import workers from "./services/workers"
import deployments from "./services/deployments"
import services from "./services/services"

//Vue global config
Vue.config.productionTip = false

//System components
Vue.prototype.$http = api
Vue.prototype.$bus = bus

//Services
Vue.prototype.$auth = auth
Vue.prototype.$nodes = nodes
Vue.prototype.$pods = pods
Vue.prototype.$upstream = upstream
Vue.prototype.$workers = workers
Vue.prototype.$deployments = deployments
Vue.prototype.$services = services

//Vue
Vue.use(require('vue-moment'));

new Vue({
  router,
  render: h => h(App)
}).$mount('main')

