package config

import (
	"encoding/json"
	"fmt"
)

type Backend struct {
	Name string
	Host string
	Url  string
	Auth BackendAuth
}
type BackendAuth struct {
	Token    TemplatedString
	TestAuth bool `json:"test_auth"`
}

func (b *Backend) UnmarshalJSON(data []byte) error {
	type tempBackend struct {
		Name string
		Host string
		Url  *string
		Auth BackendAuth
	}
	var tb tempBackend
	if err := json.Unmarshal(data, &tb); err != nil {
		return err
	}

	b.Name = tb.Name
	b.Host = tb.Host
	b.Auth = tb.Auth
	if tb.Url == nil {
		b.Url = fmt.Sprintf("https://%s/%s", b.Host, b.Name)
	} else {
		b.Url = *tb.Url
	}
	return nil
}
