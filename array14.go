package main

import "fmt"

// Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão
// ser trocados de posição. Imprima o Slice resultante.
func main() {
	var posicao1 int
	var posicao2 int
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Slice atual:", slice)
	fmt.Print("Insira dois elementos que devem ser trocados de posição(entre 0-6): ")
	fmt.Scan(&posicao1)
	fmt.Scan(&posicao2)
	salvarposicao := slice[posicao1]
	slice[posicao1] = slice[posicao2]
	slice[posicao2] = salvarposicao
	fmt.Println("Novo slice:", slice)
}
