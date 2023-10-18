package greetings

import "fmt"

func hello(name string) string {
	message := fmt.Sprintf("Hola, %v! Bienvenido!", name)
	return message
}

