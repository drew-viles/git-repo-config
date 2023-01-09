package backends

import (
	"fmt"
	"gitrepoconfig/pkg/backends/github"
	"gitrepoconfig/pkg/config"
)

type BackendApi interface {
	New(config config.GitRepoConfig, backend config.Backend)
	Auth() error
	WritePermissions() error
}

func New(configData config.GitRepoConfig, backend config.Backend) BackendApi {
	if backend.Host == "github" {
		fmt.Println("github repo found, setting up now")
		api := github.Api{}
		api.New(configData, backend)
		return &api
	}
	return nil
}
