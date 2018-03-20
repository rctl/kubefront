package core

import "k8s.io/client-go/kubernetes"

//Context is used for kubefront backend states
type Context struct {
	Config *Config
	Client *kubernetes.Clientset
}
