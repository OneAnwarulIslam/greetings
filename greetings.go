package greetings

import "fmt"

// Append string in a message
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Wellcome!", name)
	return message
}
