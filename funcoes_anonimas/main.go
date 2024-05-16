package main

import "fmt"

func main() {

	total := func() int {
		return soma(1212, 4245, 12121, 232323, 121212) * 2
	}()

	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
