package auth

import (
	"log"
	"os"
)

func Login() {
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}
