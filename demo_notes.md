# Kubefront

Easy to user event driven dashboard for Kubernetes.


*Status of project*
- Authentication and access control.
- WebSocktets for emitting live cluster events.
- Frontend *event-bus* allowing components to listen to and react to events.
- Custom built frontend "services" to manage states and keeping track of collections.
- Support for Pods, Deployments, Services and Nodes.

*Weaknesses*
- Access control should have been setup to better reflect Kubernetes RBAC.
- Collection in frontend "services" can be refreshed unnecessarily if user navigates often in the UI.
- Components are not build 100% unifom, which creates "spaghetti code".
- Worker logic (running Go routines) can sometimes jump out of sync in Vue.

*Technologies*
- Vue is used for frontend. Different UI components are split into Vue components making for easy re-use.
- WebSockets for event driven system, no states needs to be fully refreshed unless user changes its UI "scope".
- JWT for tokens, route-based authentication.

## Architecture

![Kubefront architecture](kubefront-architecture.png "Kubefront architecture")

## Frontend Spec

![](https://docs.google.com/drawings/d/e/2PACX-1vShYqvxSYGSLs-0xPv1Hk1-FricvrW4lEqgDvCWZk9Ql94zEeywk-D5uhHvYRNZ2oAojnpYsa2V4TOW/pub?w=768&h=553)

To give you an example:

![](https://docs.google.com/drawings/d/e/2PACX-1vTm6ti5uPC3M7aSty3L8M_3yqCAa3L6p1UN4MQc-6GiHIHH2xvOkOyFWmfRZGPw-3w4nNHCRhlexbc8/pub?w=838&h=272)

## System demo

Now I will show you the system in action

## Code demo

We will take a look at theese files:

`src/frontend/main.js`

`src/frontend/bus.js`

`src/frontend/services/pods.js`

`src/frontend/services/upstream.js`

`src/frontend/services/workers.js`

`src/frontend/App.vue`

`src/frontend/views/Dashboard.vue`

`src/frontend/components/NodeList.vue`

`src/frontend/components/NodeAllocatableMemory.vue`

`src/frontend/views/Deployments.vue`

`src/frontend/components/DeploymentList.vue`