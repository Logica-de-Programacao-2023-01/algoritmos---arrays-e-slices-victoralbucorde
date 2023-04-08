package main

import "fmt"

// Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
func main() {

	array := [4]float64{3.2, 12.2, 19.2, 8.2}
	soma := 1.0

	for i := 0; i < len(array); i++ {
		soma *= array[i]
	}
	fmt.Println(soma)
}
