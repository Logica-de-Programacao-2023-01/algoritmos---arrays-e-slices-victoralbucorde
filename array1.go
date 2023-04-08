package main

import "fmt"

// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
func main() {

	array := [3]int{1, 2, 3}
	resultado := array[1] + array[2] + array[0]

	fmt.Println("A soma dos elementos dentro do array é", resultado)
}
