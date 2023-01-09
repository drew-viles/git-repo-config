package main

import (
	"fmt"
	"gitrepoconfig/pkg/backends"
	"gitrepoconfig/pkg/config"
)

func main() {
	configData, err := config.GetRepoConfig(".gitrepo.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Configuring repo: %s\n", configData.Name)

	for _, backend := range configData.Backends {
		backendApi := backends.New(configData, backend)
		if err := backendApi.Auth(); err != nil {
			fmt.Println(err)
		}
	}
}
