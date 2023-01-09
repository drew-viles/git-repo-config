package github

import (
	"context"
	"fmt"
	githubapi "github.com/google/go-github/v49/github"
	"gitrepoconfig/pkg/config"
	"golang.org/x/oauth2"
	"log"
)

type Api struct {
	configData config.GitRepoConfig
	backend    config.Backend
	client     *githubapi.Client
}

func (b *Api) New(configData config.GitRepoConfig, backend config.Backend) {
	b.configData = configData
	b.backend = backend
}

func (b *Api) Auth() error {
	fmt.Printf("authenticating '%s'\n", b.configData.Name)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: b.backend.Auth.Token.String()},
	)
	tc := oauth2.NewClient(ctx, ts)
	b.client = githubapi.NewClient(tc)
	if b.backend.Auth.TestAuth {
		fmt.Println("Testing Auth")
		_, resp, err := b.client.Users.Get(ctx, "")
		if err != nil {
			return err
		}
		if !resp.TokenExpiration.IsZero() {
			log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
		}
	}
	return nil
}

func (b *Api) WritePermissions() error {
	fmt.Println("writing permissions to api")
	return nil
}
