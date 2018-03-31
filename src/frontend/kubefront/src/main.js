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
import upstream from "./services/upstream"

Vue.config.productionTip = false
//Default protos
Vue.prototype.$http = api
Vue.prototype.$bus = bus
Vue.prototype.$auth = auth
Vue.prototype.$nodes = nodes
Vue.prototype.$upstream = upstream

Vue.use(require('vue-moment'));

new Vue({
  router,
  render: h => h(App)
}).$mount('main')

