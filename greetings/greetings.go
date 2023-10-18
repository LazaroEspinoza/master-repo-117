package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hola, %v! Bienvenido!", name)
	return message
}
