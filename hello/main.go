package main

import (
	"fmt"
	"log"

	"github.com/MASTER-REPO-117/greetings"
)

func main() {
	log.SetPrefix("Greetings:\n")
	log.SetFlags(0)

	names := []string{"Christopher", "Pier", "Camila", "Camerina", "Raymundo"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	//message, err := greetings.Hello(names())
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(messages)

}
