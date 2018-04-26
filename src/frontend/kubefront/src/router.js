import Vue from 'vue'
import Router from 'vue-router'
import Welcome from './views/Welcome.vue'
import Nodes from './views/Nodes.vue'
import Pods from './views/Pods.vue'
import Dashboard from './views/Dashboard.vue'
import Deployments from './views/Deployments.vue'
import Services from './views/Services.vue'

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
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/nodes',
      name: 'nodes',
      component: Nodes
    },
    {
      path: '/pods',
      name: 'pods',
      component: Pods
    },
    {
      path: '/deployments',
      name: 'deployments',
      component: Deployments
    },
    {
      path: '/services',
      name: 'services',
      component: Services
    }
  ]
})
