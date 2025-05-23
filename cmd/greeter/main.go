package main

import (
	"fmt"

	"github.com/emanivinay/golang-mastery/greetings"
)

func main() {
	greeting := greetings.Hello("Vinay")
	fmt.Println(greeting)
}
