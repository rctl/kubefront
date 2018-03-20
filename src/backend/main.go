package main

import "github.com/rctl/kubefront/src/backend/kubefront"

func main() {
	k := kubefront.New("my-secret-key")
	k.Serve()
}
