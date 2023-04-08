package main

import "fmt"

// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.
func main() {

	array := [3][2]int{}

	for i := 0; i < len(array); i++ {
		for coluna := 0; coluna < len(array[i]); coluna++ {
			var x int
			fmt.Printf("Digite o elemento da linha %d e coluna %d ", i, coluna)
			fmt.Scan(&x)
			array[i][coluna] = x
		}
	}
	fmt.Println(array)
}
