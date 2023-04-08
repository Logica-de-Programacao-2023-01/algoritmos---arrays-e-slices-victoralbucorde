package main

import "fmt"

// Crie um Array de floats com 10 elementos.
// Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5.
// Imprima o novo Slice.
func main() {

	array := [10]float64{1.9, 5.1, 5.0, 4.6, 9.9, 11.2, 7.8, 8.4, 6.6, 10.1}
	var slice []float64
	fmt.Println("Atual lista: ", array)
	for i := 0; i < len(array); i++ {
		if array[i] > 5 {
			slice = append(slice, array[i])
		}
	}
	fmt.Println("Nova lista com numeros maiores que 5: ", slice)
}
