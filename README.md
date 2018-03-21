# Kubefront - Simple cluster management

**Project vision**

Kubefront is a dashboard for Kubernetes clusters that simplifies cluster management. The development of kubefront is taking place as a part of a course in advanced web programming at Linköping University. Kubefront allows an administratior to manage and monitor the cluster on a very high level. It also allow other users to run apps (a collection of Kubernetes pods, deployments, services, ingresses) in the cluster without any real knowledge of Kubernetes itself. Kubefront can be safely exposed publicly and be used by people with no knowledge of Kubernetes.

**Functional requirements**

*User can*

- [ ] Run Kubefront in Kubernetes and access the Web UI
- [x] Basic auth system and ability to sign in as administrator
- [ ] Create manage and delete other user accounts
- [ ] See basic statistics from the cluster
- [ ] Manage the cluster with basic operations (managing minions)
- [ ] Build application definitions (templating applications)
- [ ] Create basic applications consisting of only a deployment
- [ ] Manage and delete basic applications
- [ ] Customer mode (users that can only create instances of applications with limitations)
- [ ] Create advanced applications consisting of both a deployment, service and an ingress
- [ ] Manage and delete advanced applications
- [ ] Controlling what applications are available to be created by specific users (apps "marketplace")

**Technological specification**

- Vue is used on the client side. All client-side programming is made in JavaScript and packed in Webpack as part of standard Vue procedures. View components are split when relevant.
- Go is used as the backend programming language acting as an authenticator and mediator between the client-side and Kubernetes API. Because authentication is built into the backend the backend can safely be exposed outside the cluster.
- The backend runs in a Docker container. It can connect to a cluster from within the cluster or externally through the Kubernetes API. Because of this it can be run locally and be connected to a remote cluster. 
- Where relevant, the backend contains UNIT tests to ensure quality. However, because the underlying Kubernetes API is hard to mock it will not have any UNIT tests for that interactions. In that case, e2e tests will be used instead, if required.
- Custom built bash scripts to simplify building, running, deploying and testing kubefront will be made.
- Kubefront workers (go routines launced from the Kubefront API) will be used to make async tasks for the user.

The image below explains how the Kubefront architecture looks like.

![Kubefront architecture](kubefront-architecture.png "Kubefront architecture")
