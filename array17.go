package main

import "fmt"

// Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.
func main() {

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var soma int
	fmt.Println("Lista Atual: ", array)
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			soma += array[i]
		}
	}
	fmt.Println("Valor da sua soma é igual a ", soma)
}
