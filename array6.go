package main

import "fmt"

// Crie um Array de inteiros com 10 elementos. Em seguida,
// solicite ao usuário que informe um valor e verifique
// se esse valor está presente no Array. Imprima o resultado da busca.
func main() {

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var numero int
	fmt.Print("Escreva um número para procurar na lista ")
	fmt.Scan(&numero)
	confirmacao := false
	for i := 0; i < len(array); i++ {
		if numero == array[i] {
			confirmacao = true
		}
	}
	if confirmacao {
		fmt.Println("O número", numero, "está na lista")
	} else {
		fmt.Println("O número", numero, "não está na lista")
	}
}
