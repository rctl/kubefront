package core

import "github.com/ericchiang/k8s"

//Context is used for kubefront backend states
type Context struct {
	Config *Config
	Client *k8s.Client
}
