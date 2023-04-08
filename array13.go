package main

import "fmt"

// Crie um Array de inteiros com 7 elementos.
// Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array.
// Imprima o Array resultante.
func main() {
	var primeiro int
	var ultimo int
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Array atual:", array)
	fmt.Print("Escreva um numero que será adicionado ao primeiro e último elemento: ")
	fmt.Scan(&primeiro)
	fmt.Scan(&ultimo)
	array[0] = array[0] + primeiro
	array[6] = array[6] + ultimo
	fmt.Println("Novo array:", array)
}
