package main

import "fmt"

// Crie um Array de inteiros com 10 elementos.
// Crie um novo Slice que contenha apenas os elementos pares do Array original.
// Imprima o novo Slice.
func main() {

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int

	fmt.Println("Lista atual: ", array)
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			slice = append(slice, array[i])
		}
	}
	fmt.Println("Numeros pares da antiga lista: ", slice)
}
