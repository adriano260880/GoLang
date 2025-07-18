package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 10

	fmt.Println("Main-----------")
	fmt.Println(x, *y)
	fmt.Println(&x, y)
	Imprime(&x, y)
}

func Imprime(x *int, y *int) {
	fmt.Println("Imprime-----------")
	fmt.Println(x, y)
	fmt.Println(*x, *y)
}
