package main

import "fmt"

// Crie um Array de inteiros com 5 elementos.
// Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3.
// Imprima o novo Slice.
func main() {

	array := [5]int{1, 3, 6, 5, 12}
	fmt.Println("Array atual:", array)
	slice := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 {
			slice = append(slice, array[i])
		}
	}
	fmt.Println("Multiplos de 3 dentro do array:", slice)
}
