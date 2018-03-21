import Vue from 'vue'
import Router from 'vue-router'
import Welcome from './views/Welcome.vue'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/auth',
      name: 'signin',
      component: Welcome
    },
    {
      path: '/',
      name: 'home',
      component: Home
    }
  ]
})
