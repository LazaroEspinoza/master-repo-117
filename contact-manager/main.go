package main

import (
	"fmt"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("No se puede dividir por cero")

	}
	return dividendo / divisor, nil

}

func main() {

	file, err := os.Create("Hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("Hola, Lazaro Espinoza"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Numeros:", num)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
	fmt.Println("Resultado:", result)

}
