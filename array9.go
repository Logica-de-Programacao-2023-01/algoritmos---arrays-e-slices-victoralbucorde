package main

import "fmt"

// Crie um Array de floats com 6 elementos.
// Solicite ao usuário que informe um número que será adicionado em todas as posições do Array.
// Imprima o Array resultante.
func main() {
	var numero float64
	array := []float64{1.1, 2.9, 4.5, 9.2, 9.8, 7.6}
	fmt.Println("Array atual:", array)
	fmt.Print("Escreva um número para ser adicionado em todas as posições do array: ")
	fmt.Scan(&numero)
	for i := 0; i < len(array); i++ {
		array[i] = array[i] + numero
	}
	fmt.Println("Novo array:", array)
}
