package main

import (
	"fmt"
	auth "test/auth"
)

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

		}
	}()
	valiteZero(divisor)
	fmt.Println(dividendo / divisor)

}
func valiteZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")

	}

}

func main() {
	divide(100, 10)
	divide(520, 8)
	divide(700, 0)
	divide(400, 4)

	auth.Login()
}
