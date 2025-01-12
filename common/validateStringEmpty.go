package common

import (
	"errors"
	"strings"
)

func Validate(serviceName string) error {
	serviceName = strings.TrimSpace(serviceName)

	if serviceName == "" {
		return errors.New(`value is empty`)
	}
	return nil
}

func Validates(s []string) error {
	var valueNone int
	for _, v := range s {
		if err := Validate(v); err != nil {
			valueNone++
		}
	}

	if valueNone == len(s) {
		return errors.New(`value is empty`)
	}
	return nil
}
