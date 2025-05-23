package main

import (
	"fmt"
	"os"

	"github.com/emanivinay/golang-mastery/greetings"
)

const (
	ERR_NO_CLI_ARGS          = 1
	ERR_EMPTY_CLI_ARG_STRING = 2
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("must provide atleast one non-empty cli argument"))
		os.Exit(ERR_NO_CLI_ARGS)
	}
	greeting, err := greetings.Hello(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ERR_EMPTY_CLI_ARG_STRING)
	}
	fmt.Println(greeting)
}
