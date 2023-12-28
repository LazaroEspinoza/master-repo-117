package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 && t.Minute() < 59 {
		fmt.Println("Es de manana!!")

	} else if t.Hour() < 17 && t.Minute() < 59 {
		fmt.Println("Ya es de tarde!")

	} else {
		fmt.Println("Ya es de noche!")
	}

	switch t := time.Now(); {
	case t.Hour() < 12 && t.Minute() < 59:
		fmt.Println("Ya es de manana!!")
	case t.Hour() < 17 && t.Minute() < 59:
		fmt.Println("Ya es de tarde!!")
	default:
		fmt.Println("Ya es de noche!!")

	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run --> Windows")
	case "linux":
		fmt.Println("Go run --> Linux")
	case "darwin":
		fmt.Println("Go run --> MacOS")
	default:
		fmt.Println("Go run --> Otro OS")
	}
}
