package config

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type SymVer struct {
	Major int64
	Minor int64
	Patch int64
}

func (v SymVer) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *SymVer) UnmarshalJSON(data []byte) error {
	var dat string

	if err := json.Unmarshal(data, &dat); err != nil {
		return err
	}

	if err := v.FromString(dat); err != nil {
		return err
	}

	return nil
}

func (v *SymVer) FromString(input string) error {
	items := strings.Split(input, ".")
	if len(items) != 3 {
		return fmt.Errorf("version length is wrong %d", len(items))
	}

	var err error
	if v.Major, err = strconv.ParseInt(items[0], 10, 0); err != nil {
		return err
	}
	if v.Minor, err = strconv.ParseInt(items[1], 10, 0); err != nil {
		return err
	}
	if v.Patch, err = strconv.ParseInt(items[2], 10, 0); err != nil {
		return err
	}

	return nil
}
