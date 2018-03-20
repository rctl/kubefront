package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"

	"github.com/rctl/kubefront/src/backend/kubefront"
)

func main() {
	var kubeconfig *string
	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	secret := flag.String("secret", "", "(required) secret key to use for user token validation")
	flag.Parse()
	if *secret == "" {
		flag.PrintDefaults()
		return
	}
	data, err := ioutil.ReadFile(*kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
	client, err := k8s.NewClient(&config)
	if err != nil {
		panic(err.Error())
	}
	k := kubefront.New("*secret", client)
	k.Serve()
}
