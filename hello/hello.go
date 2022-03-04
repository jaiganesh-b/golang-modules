package main

import (
	"fmt"

	"codebot.com/greetings"
)

func main() {
	message := greetings.Hello("Harry")
	fmt.Println(message)
}
