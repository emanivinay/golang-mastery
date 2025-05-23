package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	return fmt.Sprintf("Hi %s, Welcome!", name), nil
}
