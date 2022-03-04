package hello

import (
	"fmt"
	"greetings"
)

func greet() {
	message := greetings.Hello("Harry")
	fmt.Println(message)
}
