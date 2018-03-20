# Kubefront - Simple cluster management

Kubefront is a dashboard for Kubernetes clusters that simplifies cluster management. The development of kubefront is taking place as a part of a course in advanced web programming at Linköping University. Kubefront allows an administratior to manage and monitor the cluster on a very high level. It also allow other users to run apps (a collection of Kubernetes pods, deployments, services, ingresses) in the cluster without any real knowledge of Kubernetes itself. 

**Kubefront can be safely exposed publicly and be used by people with no knowledge of Kubernetes.**

**Tech:**

- Frontend is built using Vue
- Backend is built using Go
- Websockets allow live data to be presented in the web UI

**Progress report:**

- [ ] Run Kubefront in Kubernetes and access the Web UI
- [ ] Sign in as administrator and create new accounts with specific permissions
- [ ] Manage and delete user accounts
- [ ] See basic statistics from the cluster
- [ ] Manage the cluster with basic operations (managing minions)
- [ ] Build application definitions (templating applications)
- [ ] Create basic applications consisting of only a deployment
- [ ] Manage and delete basic applications
- [ ] Customer mode (users that can only create instances of applications with limitations)
- [ ] Create advanced applications consisting of both a deployment, service and an ingress
- [ ] Manage and delete advanced applications
- [ ] Controlling what applications are available to be created by specific users (apps "marketplace")
