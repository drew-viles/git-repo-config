package config

import (
	"encoding/json"
	"fmt"
)

type Visibility uint8

const (
	Private Visibility = iota + 1
	Public
	Test
)

var visibilityName = map[string]Visibility{
	"private": Private,
	"public":  Public,
	"testing": Test,
}

var visibilityVal = map[Visibility]string{
	Private: "private",
	Public:  "public",
	Test:    "testing",
}

func (v *Visibility) UnmarshalJSON(data []byte) error {
	var dat string

	if err := json.Unmarshal(data, &dat); err != nil {
		return err
	}

	val, ok := visibilityName[dat]
	if !ok {
		return fmt.Errorf("visibility contains invalid value")
	}

	*v = val

	return nil
}

func (v Visibility) String() string {
	return visibilityVal[v]
}
