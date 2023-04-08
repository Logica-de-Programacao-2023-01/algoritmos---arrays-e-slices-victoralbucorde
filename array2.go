package main

import "fmt"

// Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.
func main() {

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
}
