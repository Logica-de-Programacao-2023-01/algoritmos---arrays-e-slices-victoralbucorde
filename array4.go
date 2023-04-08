package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice
// e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
func main() {

	slice := []int{}
	var tamanhoSlice int
	fmt.Print("Escreva o tamanho desejado do slice ")
	fmt.Scan(&tamanhoSlice)
	fmt.Println("Escreva os termos dentros do slice")
	for i := 0; i < tamanhoSlice; i++ {
		var x int
		fmt.Scan(&x)
		slice = append(slice, x)
	}
	fmt.Println(slice)
}
