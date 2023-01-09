package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type TemplatedString string

var re = regexp.MustCompile("(\\${([^.${}]+)\\.([^.${}]+)})")

func (t *TemplatedString) UnmarshalJSON(data []byte) error {
	var dat string
	if err := json.Unmarshal(data, &dat); err != nil {
		return err
	}

	res := re.FindAllStringSubmatch(dat, -1)

	for _, m := range res {
		if m[2] == "env" {
			val, ok := os.LookupEnv(m[3])
			if !ok {
				return fmt.Errorf("config pointing at missing env var '%s'", m[3])
			}
			dat = strings.Replace(dat, m[1], val, 1)
		}
	}

	*t = TemplatedString(dat)
	return nil
}

func (t TemplatedString) String() string {
	return string(t)
}
