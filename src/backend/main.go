package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rctl/kubefront/src/backend/kubefront"
)

func main() {
	secret := flag.String("secret", "", "(required) secret key to use for user token validation")
	var kubeconfig *string
	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	dbpath := flag.String("db", "./data.db", "(optional) path to sqlite3 database file")
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
	db, err := sql.Open("sqlite3", *dbpath)
	if err != nil {
		panic(err.Error())
	}
	k := kubefront.New(*secret, client, db)
	if err := k.InititalizeEmptyDatabase(); err != nil {
		fmt.Println("Failed to initialize database.")
		fmt.Println(err.Error())
	}
	password, err := k.CreateAdminUser()
	if err != nil {
		fmt.Println("Failed to create admin user.")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Admin password is: %s\n", password)
	}
	k.Serve(":8081")
}
